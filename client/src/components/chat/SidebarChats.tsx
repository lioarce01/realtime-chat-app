import React from "react";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { SidebarProps } from "@/types/UserTypes";

const SidebarChats: React.FC<SidebarProps> = ({
  chats,
  setSelectedChatId,
  selectedChatId,
  userId,
}) => {
  return (
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
        <div className="p-4 text-neutral-300 w-full flex items-center justify-center">
          Loading chats...
        </div>
      )}
    </div>
  );
};

export default SidebarChats;
