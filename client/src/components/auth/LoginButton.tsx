"use client";

import React from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { Button } from "@/components/ui/button";

const LoginButton = () => {
  const { loginWithRedirect, isAuthenticated, isLoading } = useAuth0();

  const handleLogin = async () => {
    await loginWithRedirect({
      appState: { returnTo: "/chat" },
    });
  };

  return (
    <div>
      {!isAuthenticated && (
        <Button disabled={isLoading} onClick={handleLogin}>
          Login
        </Button>
      )}
    </div>
  );
};

export default LoginButton;
