import { useAuth0 } from "@auth0/auth0-react";
import useSyncAuth0WithBackend from "@/hooks/useSyncAuth0WithBackend";
import { ReactNode } from "react";

const AuthSyncWrapper = ({ children }: { children: ReactNode }) => {
  const { isAuthenticated, isLoading } = useAuth0();

  useSyncAuth0WithBackend();

  if (isLoading)
    return (
      <div className="w-full bg-black bg-opacity-50 h-screen flex text-white justify-center items-center">
        {" "}
        Loading...
      </div>
    );
  if (!isAuthenticated) return <p>Please log in.</p>;

  return <>{children}</>;
};

export default AuthSyncWrapper;
