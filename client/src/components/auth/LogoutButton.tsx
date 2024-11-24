import React from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { Button } from "@/components/ui/button";

const LogoutButton = () => {
  const { isAuthenticated, logout, isLoading } = useAuth0();

  const handleLogout = async () => {
    await logout({
      logoutParams: { returnTo: window.location.origin },
    });
  };

  return (
    <>
      {isAuthenticated && (
        <button
          disabled={isLoading}
          onClick={handleLogout}
          className="w-full text-left"
        >
          Logout
        </button>
      )}
    </>
  );
};

export default LogoutButton;
