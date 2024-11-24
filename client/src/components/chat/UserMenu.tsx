import React from "react";
import { UserMenuProps } from "@/types/UserTypes";
import CustomDropdownMenu from "@/components/ui/custom-dropdown-menu";

const UserMenu: React.FC<UserMenuProps> = ({ dbUser }) => {
  return <CustomDropdownMenu dbUser={dbUser} />;
};

export default UserMenu;
