"use client";

import LoginButton from "@/components/auth/LoginButton";
import LogoutButton from "@/components/auth/LogoutButton";

export default function Home() {
  return (
    <div className="flex flex-col space-y-6 items-center justify-center min-h-screen">
      <div className="flex flex-col items-center">
        <h2 className="text-6xl font-bold">CHAT APP</h2>
        <p className="">Create an account to start chatting</p>
      </div>
      <LoginButton />
    </div>
  );
}
