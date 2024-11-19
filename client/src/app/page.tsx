"use client";

import LoginButton from "@/components/auth/LoginButton";
import LogoutButton from "@/components/auth/LogoutButton";

export default function Home() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <h2 className="text-lg font-bold mb-4">Welcome to my Chat App </h2>
      <LoginButton />
      <LogoutButton />
    </div>
  );
}
