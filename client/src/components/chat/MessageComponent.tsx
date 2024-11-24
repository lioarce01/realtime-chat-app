import React, { useEffect, useRef } from "react";
import { Message, MessagesProps } from "@/types/MessageTypes";
import { formatDate } from "@/lib/utils";

const MessageComponent: React.FC<MessagesProps> = ({ messages, dbUserId }) => {
  const messagesEndRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  return (
    <div className="flex-1 overflow-y-auto p-4 bg-neutral-900">
      {messages && messages.length > 0 ? (
        messages.map((message: Message, index: number) => {
          const isOwnMessage = message.sender?.id === dbUserId;

          return (
            <div
              key={message.id || `message-${index}`}
              className={`mb-4 flex ${
                isOwnMessage ? "justify-end" : "justify-start"
              }`}
            >
              <div
                className={`max-w-sm p-3 rounded-2xl shadow-lg ${
                  isOwnMessage
                    ? "bg-neutral-800 text-cyan-400 border border-cyan-400"
                    : "bg-neutral-800 text-pink-400 border border-pink-400"
                }`}
              >
                <p className="text-sm leading-relaxed">{message.content}</p>
                <p className="text-xs text-neutral-400 text-right mt-2">
                  {formatDate(message?.created_at)}
                </p>
              </div>
            </div>
          );
        })
      ) : (
        <p className="text-center text-gray-500 w-full h-full flex justify-center items-center">
          Loading messages...
        </p>
      )}
      <div ref={messagesEndRef} />
    </div>
  );
};

export default MessageComponent;
