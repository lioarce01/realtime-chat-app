import { useAuth0 } from "@auth0/auth0-react";
import {
  useRegisterUserMutation,
  useLazyGetUserByIdQuery,
} from "@/redux/api/userApi";
import { useEffect } from "react";

const useSyncAuth0WithBackend = () => {
  const { user, isAuthenticated } = useAuth0();
  const [register] = useRegisterUserMutation();
  const [getUserById, { data: userData, error: userError }] =
    useLazyGetUserByIdQuery();

  useEffect(() => {
    if (isAuthenticated && user) {
      const { email, sub, name, picture } = user;

      getUserById(sub)
        .unwrap()
        .then((existingUser) => {
          console.log("User already exists:", existingUser);
        })
        .catch(() => {
          register({
            email,
            username: name || "@"[0],
            sub: sub,
            profile_pic: picture,
          })
            .unwrap()
            .then((response: any) => {
              console.log("User registered successfully:", response);
            })
            .catch((error: any) => {
              console.error("Error during registration:", error);
            });
        });
    }
  }, [isAuthenticated, user, getUserById, register]);
};

export default useSyncAuth0WithBackend;
