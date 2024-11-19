import React from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { Button } from "@/components/ui/button";

const LogoutButton = () => {
  const { isAuthenticated, logout, isLoading } = useAuth0();

  const handleLogout = () => {
    logout({
      logoutParams: { returnTo: window.location.origin },
    });
  };
  return (
    <div>
      {isAuthenticated && (
        <Button disabled={isLoading} onClick={handleLogout}>
          Logout
        </Button>
      )}
    </div>
  );
};

export default LogoutButton;