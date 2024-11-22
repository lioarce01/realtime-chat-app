import { useAuth0 } from "@auth0/auth0-react";
import useSyncAuth0WithBackend from "@/hooks/useSyncAuth0WithBackend";
import { ReactNode } from "react";

const AuthSyncWrapper = ({ children }: { children: ReactNode }) => {
  const { isAuthenticated, isLoading } = useAuth0();

  useSyncAuth0WithBackend();

  if (isLoading) return <p>Loading...</p>;
  if (!isAuthenticated) return <p>Please log in.</p>;

  return <>{children}</>;
};

export default AuthSyncWrapper;
