"use client";

import React, { useState } from "react";
import { motion } from "framer-motion";
import { useGetUserByIdQuery } from "@/redux/api/userApi";
import AuthSyncWrapper from "@/components/auth/AuthSyncWrapper";
import { useAuth0 } from "@auth0/auth0-react";
import Chat from "@/components/chat/Chat";
import ChatInput from "@/components/chat/Input";
import Sidebar from "@/components/chat/SidebarChat";

interface ChatContact {
  id: string;
  user1_id: string;
  user2_id: string;
  created_at?: string;
  updated_at?: string;
  messages?: any[];
}

const ChatPage = () => {
  const { user } = useAuth0();
  const [selectedChatId, setSelectedChatId] = useState<string | null>(null);
  const { data: dbUser } = useGetUserByIdQuery(user?.sub ?? "", {
    skip: !user?.sub,
  });

  const dbUserId = dbUser?.user?.id;

  return (
    <AuthSyncWrapper>
      <div className="flex items-center justify-center w-full h-screen bg-[#0c0c0c] text-white">
        <motion.div
          initial={{ opacity: 0, scale: 0.9 }}
          animate={{ opacity: 1, scale: 1 }}
          transition={{ duration: 0.5 }}
          className="h-full w-full md:h-[95%] md:w-[90%] lg:w-[80%] rounded-sm shadow-md shadow-neutral-950 overflow-hidden"
        >
          <div className="flex justify-between w-full h-full">
            {/* Sidebar */}
            <Sidebar
              userId={user?.sub!}
              setSelectedChatId={setSelectedChatId}
              selectedChatId={selectedChatId}
            />
            {/* Chat Window */}
            <div className="flex-1 flex flex-col bg-neutral-950">
              {/* Chat Messages */}
              <Chat chatId={selectedChatId} dbUserId={dbUserId} />

              {/* Input Bar */}
              {/* <ChatInput
                inputMessage={inputMessage}
                setInputMessage={setInputMessage}
                handleSendMessage={handleSendMessage}
              /> */}
            </div>
          </div>
        </motion.div>
      </div>
    </AuthSyncWrapper>
  );
};

export default ChatPage;
