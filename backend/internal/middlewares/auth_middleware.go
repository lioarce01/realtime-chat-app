package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gin-gonic/gin"
)

type AuthConfig struct {
    Audience     string
    Domain       string
    ClientOrigin string
}

type CustomClaims struct {
    Scope string `json:"scope"`
}

func (c CustomClaims) Validate(ctx context.Context) error {
    return nil
}

func LoadAuthConfig() (*AuthConfig, error) {
    config := &AuthConfig{
        Audience:     os.Getenv("AUTH0_AUDIENCE"),
        Domain:       os.Getenv("AUTH0_DOMAIN"),
        ClientOrigin: os.Getenv("CLIENT_ORIGIN"),
    }

    if config.Audience == "" || config.Domain == "" {
        return nil, errors.New("missing required Auth0 environment variables")
    }

    return config, nil
}

func AuthMiddleware(config *AuthConfig) gin.HandlerFunc {
    issuerURL := "https://" + config.Domain + "/"
    
    jwksURL, err := url.Parse(issuerURL)
    if err != nil {
        log.Fatalf("Failed to parse issuer URL: %v", err)
    }
    jwksURL.Path = ".well-known/jwks.json"
    
    provider := jwks.NewCachingProvider(jwksURL, 5*time.Minute)

    jwtValidator, err := validator.New(
        provider.KeyFunc,
        validator.RS256,
        issuerURL,
        []string{config.Audience},
        validator.WithCustomClaims(
            func() validator.CustomClaims {
                return &CustomClaims{}
            },
        ),
    )

    if err != nil {
        log.Fatalf("Failed to set up the validator: %v", err)
    }

    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"})
            return
        }

        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
            return
        }

        token := parts[1]
        claims, err := jwtValidator.ValidateToken(c.Request.Context(), token)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        c.Set("claims", claims)
        c.Next()
    }
}