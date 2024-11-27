import React from "react";
import { useGetUserByIdQuery, useGetUserChatsQuery } from "@/redux/api/userApi";
import { SidebarProps } from "@/types/UserTypes";
import SidebarChats from "./SidebarChats";
import UserMenu from "./UserMenu";
import Searchbar from "./Searchbar";

const Sidebar: React.FC<SidebarProps> = ({
  userId,
  setSelectedChatId,
  selectedChatId,
  dbUser,
  onCreateChat,
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
      <div className="p-4 flex justify-between items-center">
        <h1 className="text-2xl font-bold">Chats</h1>
        <UserMenu dbUser={dbUser} />
      </div>
      <Searchbar onCreateChat={onCreateChat} />
      <SidebarChats
        chats={chats}
        setSelectedChatId={setSelectedChatId}
        selectedChatId={selectedChatId}
        userId={userId}
        onCreateChat={onCreateChat}
      />
    </div>
  );
};

export default Sidebar;
