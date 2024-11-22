import React, { useState, useEffect, useRef, useCallback } from "react";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import {
  useGetChatByIdQuery,
  useSendMessageMutation,
} from "@/redux/api/chatApi";
import ChatInput from "./Input";
import ReconnectingWebSocket from "reconnecting-websocket";
import { useGetUserByIdQuery } from "@/redux/api/userApi";

interface ChatProps {
  chatId: string | null;
  dbUserId: string | undefined;
}

interface Message {
  id: string;
  sender: {
    id: string;
    username: string;
    profile_pic: string;
  };
  receiver: {
    id: string;
    username: string;
    profile_pic: string;
  };
  content: string;
  created_at: string;
  chat_id: string;
}

const Chat: React.FC<ChatProps> = ({ chatId, dbUserId }) => {
  const { data, isLoading, error } = useGetChatByIdQuery(chatId ?? "", {
    skip: !chatId,
  });

  const receiverUser =
    data?.chat?.user1_id === dbUserId
      ? data?.chat?.user2_id
      : data?.chat?.user1_id;

  const { data: otherUser } = useGetUserByIdQuery(receiverUser ?? "", {
    skip: !receiverUser,
  });

  const [sendMessage] = useSendMessageMutation();
  const [inputMessage, setInputMessage] = useState("");
  const [messages, setMessages] = useState<Message[]>([]);
  const wsRef = useRef<ReconnectingWebSocket | null>(null);
  const messagesEndRef = useRef<HTMLDivElement>(null);

  const formatDate = (dateString: string) => {
    const date = new Date(dateString);
    if (isNaN(date.getTime())) {
      return "Invalid Date";
    }
    return date.toISOString();
  };

  useEffect(() => {
    if (data?.chat?.messages) {
      setMessages(data.chat.messages);
    }
  }, [data]);

  const handleNewMessage = useCallback((newMessage: Message) => {
    setMessages((prevMessages) => [...prevMessages, newMessage]);
  }, []);

  useEffect(() => {
    if (!chatId || !dbUserId) return;

    const wsUrl = `ws://${window.location.hostname}:8080/ws/${dbUserId}`;
    const ws = new ReconnectingWebSocket(wsUrl);
    wsRef.current = ws;

    ws.onopen = () => {
      console.log("WebSocket Connected");
    };

    ws.onmessage = (event) => {
      try {
        const newMessage = JSON.parse(event.data);
        console.log("received message:", newMessage);
        if (newMessage.chat_id === chatId) {
          newMessage.created_at = formatDate(newMessage.created_at);
          handleNewMessage(newMessage);
        }
      } catch (error) {
        console.log("Received plain text message:", event.data);
        if (typeof event.data === "string") {
          const newMessage: Message = {
            id: Date.now().toString(),
            sender: { id: "", username: "", profile_pic: "" },
            receiver: { id: "", username: "", profile_pic: "" },
            content: event.data,
            created_at: new Date().toLocaleString(),
            chat_id: chatId,
          };
          if (newMessage.chat_id === chatId) {
            handleNewMessage(newMessage);
          }
        } else {
          console.error("Error parsing WebSocket message:", error);
        }
      }
    };

    ws.onclose = () => {
      console.log("WebSocket Disconnected - Reconnecting...");
    };

    ws.onerror = (error) => {
      console.error("WebSocket Error:", error);
    };

    return () => {
      if (wsRef.current) {
        wsRef.current.close();
      }
    };
  }, [chatId, dbUserId, handleNewMessage]);

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  const handleSendMessage = async () => {
    if (inputMessage.trim() === "" || !dbUserId || !chatId) return;

    try {
      const messageData = {
        sender_id: dbUserId,
        receiver_id: receiverUser,
        content: inputMessage,
      };

      const response = await sendMessage(messageData).unwrap();

      handleNewMessage(response);

      setInputMessage("");
    } catch (error) {
      console.error("Failed to send message:", error);
    }
  };

  const displayDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleString();
  };

  if (!chatId) {
    return (
      <div className="flex-1 flex items-center justify-center text-neutral-500">
        Select a chat to start messaging
      </div>
    );
  }

  if (isLoading) {
    return (
      <div className="flex-1 flex items-center justify-center text-neutral-500">
        Loading messages...
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex-1 flex items-center justify-center text-red-500">
        Failed to load messages. Please try again later.
      </div>
    );
  }

  const receiverData = otherUser?.user;

  return (
    <div className="flex flex-col h-full">
      <div className="p-4 bg-neutral-900 text-white flex items-center justify-between border-b border-neutral-800">
        <div className="flex items-center space-x-3">
          <Avatar>
            <AvatarImage
              src={
                receiverData?.profile_pic ||
                `https://api.dicebear.com/6.x/initials/svg?seed=${
                  receiverData?.username || "unknown"
                }`
              }
            />
            <AvatarFallback>
              {receiverData?.username?.charAt(0).toUpperCase() || "?"}
            </AvatarFallback>
          </Avatar>
          <h1 className="text-lg font-semibold">
            {receiverData?.username || "Unknown User"}
          </h1>
        </div>
        <span className="text-sm text-neutral-400">Online</span>
      </div>
      <div className="flex-1 overflow-y-auto p-4">
        {messages?.map((message: Message, index: number) => (
          <div
            key={message.id || `message-${index}`}
            className={`mb-4 flex ${
              message.sender?.id === dbUserId ? "justify-end" : "justify-start"
            }`}
          >
            <div
              className={`max-w-sm p-3 rounded-lg shadow ${
                message.sender?.id === dbUserId
                  ? "bg-blue-600 text-white"
                  : "bg-neutral-800 text-white"
              }`}
            >
              <p className="text-sm">{message.content}</p>
              <p className="text-xs text-neutral-400 text-right mt-1">
                {displayDate(message.created_at)}
              </p>
            </div>
          </div>
        ))}
        <div ref={messagesEndRef} />
      </div>
      <ChatInput
        inputMessage={inputMessage}
        setInputMessage={setInputMessage}
        handleSendMessage={handleSendMessage}
      />
    </div>
  );
};

export default Chat;
