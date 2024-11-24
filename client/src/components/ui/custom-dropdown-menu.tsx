import React, { useState, useRef, useEffect } from "react";
import { EllipsisIcon } from "lucide-react";
import Image from "next/image";
import { UserMenuProps } from "@/types/UserTypes";
import LogoutButton from "../auth/LogoutButton";

const CustomDropdownMenu: React.FC<UserMenuProps> = ({ dbUser }) => {
  const [isOpen, setIsOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);
  const menuRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node)
      ) {
        setIsOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => {
      document.removeEventListener("mousedown", handleClickOutside);
    };
  }, []);

  useEffect(() => {
    if (menuRef.current) {
      if (isOpen) {
        menuRef.current.classList.remove("opacity-0", "scale-95");
        menuRef.current.classList.add("opacity-100", "scale-100");
      } else {
        menuRef.current.classList.add("opacity-0", "scale-95");
        menuRef.current.classList.remove("opacity-100", "scale-100");
      }
    }
  }, [isOpen]);

  const toggleDropdown = () => setIsOpen(!isOpen);

  return (
    <div className="flex items-center space-x-2 relative" ref={dropdownRef}>
      <Image
        src={dbUser?.user?.profile_pic || "/default-profile-picture.jpg"}
        alt="profile picture"
        width={35}
        height={35}
        className="rounded-full"
      />
      <button
        className="hover:bg-neutral-800 p-1 transition-all duration-300 rounded-full"
        onClick={toggleDropdown}
        aria-haspopup="true"
        aria-expanded={isOpen}
      >
        <EllipsisIcon />
        <span className="sr-only">Open user menu</span>
      </button>
      <div
        ref={menuRef}
        className={`absolute right-0 top-full mt-2 w-48 rounded-md shadow-lg bg-neutral-800 ring-1 text-white ring-black ring-opacity-5 divide-y divide-neutral-700 focus:outline-none z-10 transition-all duration-200 ease-out transform origin-top-right ${
          isOpen ? "" : "invisible pointer-events-none"
        }`}
      >
        <div className="px-4 py-3">
          <p className="text-sm leading-5 font-medium">
            {dbUser?.user?.username}
          </p>
          <p className="text-sm leading-5 font-medium text-neutral-400">
            {dbUser?.user?.email}
          </p>
        </div>
        <div>
          <a
            href="/profile"
            className="block px-4 py-2 text-sm hover:bg-neutral-700 transition-all duration-300 ease-in-out"
          >
            Profile
          </a>
          <p className="block px-4 py-2 text-sm hover:bg-neutral-700 transition-all duration-300 ease-in-out rounded-b">
            <LogoutButton />
          </p>
        </div>
      </div>
    </div>
  );
};

export default CustomDropdownMenu;
