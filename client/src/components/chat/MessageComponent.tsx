import { Message, MessagesProps } from "@/types/MessageTypes";
import React, { useEffect, useRef } from "react";

const MessageComponent: React.FC<MessagesProps> = ({ messages, dbUserId }) => {
  const displayDate = (dateString: string) => {
    const date = new Date(dateString);
    return date.toLocaleString();
  };

  const messagesEndRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  return (
    <div className="flex-1 overflow-y-auto p-4">
      {messages
        ? messages.map((message: Message, index: number) => (
            <div
              key={message.id || `message-${index}`}
              className={`mb-4 flex ${
                message.sender?.id === dbUserId
                  ? "justify-end"
                  : "justify-start"
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
          ))
        : "loading messages..."}
      <div ref={messagesEndRef} />
    </div>
  );
};

export default MessageComponent;
