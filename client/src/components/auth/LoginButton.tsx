"use client";

import React from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { Button } from "@/components/ui/button";

const LoginButton = () => {
  const { loginWithRedirect, isAuthenticated, isLoading } = useAuth0();
  return (
    <div>
      {!isAuthenticated && (
        <Button disabled={isLoading} onClick={() => loginWithRedirect()}>
          Login
        </Button>
      )}
    </div>
  );
};

export default LoginButton;
