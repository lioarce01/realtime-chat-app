import React from "react";
import { useGetUserByIdQuery } from "@/redux/api/userApi";
import { Avatar, AvatarFallback } from "@/components/ui/avatar";
import { AvatarImage } from "@radix-ui/react-avatar";

interface SidebarProps {
  userId: string;
  setSelectedChatId: (chatId: string | null) => void;
  selectedChatId: string | null;
}

const Sidebar: React.FC<SidebarProps> = ({
  userId,
  setSelectedChatId,
  selectedChatId,
}) => {
  const { data, isLoading, error } = useGetUserByIdQuery(userId);

  const chats = data?.user?.chats ?? [];

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
        {chats.length > 0 ? (
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
                  src={`https://api.dicebear.com/6.x/initials/svg?seed=${
                    chat.user1_id === userId ? chat.user2_id : chat.user1_id
                  }`}
                />
                <AvatarFallback>?</AvatarFallback>
              </Avatar>
              <div className="ml-3">
                <h2 className="font-bold">
                  Chat with{" "}
                  {chat.user1_id === userId ? chat.user2_id : chat.user1_id}
                </h2>
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
