import React from "react";
import { useGetUserByIdQuery, useGetUserChatsQuery } from "@/redux/api/userApi";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { AvatarImage } from "@radix-ui/react-avatar";
import { SidebarProps } from "@/types/UserTypes";
import SidebarChats from "./SidebarChats";

const Sidebar: React.FC<SidebarProps> = ({
  userId,
  setSelectedChatId,
  selectedChatId,
}) => {
  const { data, error } = useGetUserByIdQuery(userId, {
    skip: !userId,
  });

  const { data: chatsData } = useGetUserChatsQuery(data?.user?.id, {
    skip: !data?.user?.id,
  });

  const chats = chatsData?.chats;

  if (error) {
    return <div className="w-1/3 p-4">Failed to load chats.</div>;
  }

  return (
    <div className="w-1/3 bg-neutral-900 border-r border-neutral-800 flex flex-col">
      <div className="p-4">
        <h1 className="text-2xl font-bold">Chats</h1>
      </div>
      <SidebarChats
        chats={chats}
        setSelectedChatId={setSelectedChatId}
        selectedChatId={selectedChatId}
        userId={userId}
      />
    </div>
  );
};

export default Sidebar;
