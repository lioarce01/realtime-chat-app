import React from "react";
import { useGetUserByIdQuery, useGetUserChatsQuery } from "@/redux/api/userApi";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { AvatarImage } from "@radix-ui/react-avatar";
import { SidebarProps } from "@/types/UserTypes";

const Sidebar: React.FC<SidebarProps> = ({
  userId,
  setSelectedChatId,
  selectedChatId,
}) => {
  const { data, isLoading, error } = useGetUserByIdQuery(userId, {
    skip: !userId,
  });

  const { data: chatsData } = useGetUserChatsQuery(data?.user?.id, {
    skip: !data?.user?.id,
  });

  const chats = chatsData?.chats;

  if (isLoading) {
    return <div className="w-1/3 p-4">Loading chats...</div>;
  }

  if (error) {
    return <div className="w-1/3 p-4">Failed to load chats.</div>;
  }

  return (
    <div className="w-1/3 bg-neutral-900 border-r border-neutral-800 flex flex-col">
      <div className="p-4">
        <h1 className="text-2xl font-bold">Chats</h1>
      </div>
      <div className="flex-1 overflow-y-auto">
        {chats ? (
          chats.map((chat: any) => (
            <div
              key={chat.id}
              className={`p-4 flex items-center border-b border-neutral-800 cursor-pointer ${
                selectedChatId === chat.id ? "bg-neutral-800" : ""
              }`}
              onClick={() => setSelectedChatId(chat.id)}
            >
              <Avatar>
                <AvatarImage
                  src={
                    chat.user1_id === userId
                      ? chat.user2?.profile_pic
                      : chat.user1?.profile_pic
                  }
                />
                <AvatarFallback>?</AvatarFallback>
              </Avatar>
              <div className="ml-3">
                <h2 className="font-bold">
                  {chat.user1_id === userId
                    ? chat.user2?.username
                    : chat.user1?.username}
                </h2>
                <p className="text-sm text-gray-500">
                  {chat.last_message?.content}
                </p>
              </div>
            </div>
          ))
        ) : (
          <div className="p-4 text-neutral-500">No chats available</div>
        )}
      </div>
    </div>
  );
};

export default Sidebar;
