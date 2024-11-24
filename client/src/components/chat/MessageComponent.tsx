import React, { useEffect, useRef } from "react";
import { Message, MessagesProps } from "@/types/MessageTypes";
import { formatDate } from "@/lib/utils";

const MessageComponent: React.FC<MessagesProps> = ({ messages, dbUserId }) => {
  const messagesEndRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  return (
    <div className="flex-1 overflow-y-auto p-4">
      {messages && messages.length > 0 ? (
        messages.map((message: Message, index: number) => {
          // Determina si es un mensaje enviado por el usuario actual
          const isOwnMessage = message.sender?.id === dbUserId;

          return (
            <div
              key={message.id || `message-${index}`}
              className={`mb-4 flex ${
                isOwnMessage ? "justify-end" : "justify-start"
              }`}
            >
              <div
                className={`max-w-sm p-3 rounded-lg shadow ${
                  isOwnMessage
                    ? "bg-blue-600 text-white"
                    : "bg-neutral-800 text-white"
                }`}
              >
                <p className="text-sm">{message.content}</p>
                <p className="text-xs text-neutral-400 text-right mt-1">
                  {formatDate(message?.created_at)}
                </p>
              </div>
            </div>
          );
        })
      ) : (
        <p className="text-center text-neutral-500">Loading messages...</p>
      )}
      <div ref={messagesEndRef} />
    </div>
  );
};

export default MessageComponent;
