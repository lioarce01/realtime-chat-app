import React, { useState, useEffect, useRef, useCallback } from "react";
import {
  useGetChatByIdQuery,
  useGetMessagesByChatIdQuery,
  useSendMessageMutation,
} from "@/redux/api/chatApi";
import ChatInput from "./Input";
import ReconnectingWebSocket from "reconnecting-websocket";
import { useGetUserByIdQuery } from "@/redux/api/userApi";
import MessageComponent from "./MessageComponent";
import ReceiverProfile from "./ReceiverProfile";
import { Message } from "@/types/MessageTypes";
import { ChatProps } from "@/types/ChatTypes";
import { formatDate } from "@/lib/utils";

const Chat: React.FC<ChatProps> = ({ chatId, dbUserId }) => {
  const { data } = useGetChatByIdQuery(chatId ?? "", {
    skip: !chatId,
  });

  const { data: chatMessages, isLoading } = useGetMessagesByChatIdQuery(
    chatId ?? "",
    {
      skip: !chatId,
    }
  );

  const receiverUser =
    data?.chat?.user1_id !== dbUserId
      ? data?.chat?.user1_id
      : data?.chat?.user2_id;

  const { data: otherUser } = useGetUserByIdQuery(receiverUser ?? "", {
    skip: !receiverUser,
  });

  const [sendMessage] = useSendMessageMutation();
  const [inputMessage, setInputMessage] = useState("");
  const [messages, setMessages] = useState<Message[]>([]);
  const wsRef = useRef<ReconnectingWebSocket | null>(null);

  useEffect(() => {
    if (chatMessages?.messages) {
      setMessages(chatMessages.messages);
    }
  }, [chatMessages]);

  const handleNewMessage = useCallback(
    (newMessage: Message) => {
      setMessages((prevMessages) => {
        const updatedMessage = {
          ...newMessage,
          sender:
            newMessage.sender?.id === dbUserId
              ? { id: dbUserId, username: dbUserId, profile_pic: "" }
              : newMessage.sender,
        };
        return [...prevMessages, updatedMessage];
      });
    },
    [dbUserId]
  );

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
          newMessage.created_at = newMessage.created_at;
          newMessage.sender =
            newMessage.sender_id === dbUserId
              ? { id: dbUserId, username: dbUserId, profile_pic: "" }
              : newMessage.sender;
          handleNewMessage(newMessage);
        }
      } catch (error) {
        console.log("Received plain text message:", event.data);
        if (typeof event.data === "string") {
          const newMessage: Message = {
            sender: { id: dbUserId, username: dbUserId, profile_pic: "" },
            receiver: { id: receiverUser ?? "", username: "", profile_pic: "" },
            content: event.data,
            created_at: new Date().toISOString(),
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
  }, [chatId, dbUserId, handleNewMessage, receiverUser]);

  const handleSendMessage = async () => {
    if (inputMessage.trim() === "" || !dbUserId || !chatId) return;

    try {
      const messageData = {
        sender_id: dbUserId,
        receiver_id: receiverUser,
        content: inputMessage,
        chat_id: chatId,
      };

      // Enviar mensaje al backend
      await sendMessage(messageData).unwrap();

      // Limpiar el input después del envío
      setInputMessage("");
    } catch (error) {
      console.error("Failed to send message:", error);
    }
  };

  if (!chatId) {
    return (
      <div className="flex-1 flex items-center justify-center text-neutral-500">
        Select a chat to start messaging
      </div>
    );
  }

  return (
    <div className="flex flex-col h-full">
      <ReceiverProfile receiverData={otherUser?.user} isLoading={isLoading} />
      <MessageComponent messages={messages} dbUserId={dbUserId} />

      <ChatInput
        inputMessage={inputMessage}
        setInputMessage={setInputMessage}
        handleSendMessage={handleSendMessage}
      />
    </div>
  );
};

export default Chat;
