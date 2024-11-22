"use client";

import React from "react";
import { Button } from "../ui/button";
import { Send } from "lucide-react";
import { Input } from "../ui/input";

interface ChatInputProps {
  inputMessage: string;
  setInputMessage: React.Dispatch<React.SetStateAction<string>>;
  handleSendMessage: () => void;
}

const ChatInput: React.FC<ChatInputProps> = ({
  inputMessage,
  setInputMessage,
  handleSendMessage,
}) => {
  return (
    <div className="p-4 bg-neutral-900 flex items-center space-x-2">
      <Input
        type="text"
        placeholder="Type a message..."
        value={inputMessage}
        onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
          setInputMessage(e.target.value)
        }
        className="flex-1 bg-neutral-800 border-neutral-700 text-white placeholder-neutral-400 focus:ring-2 focus:ring-neutral-700"
        onKeyDown={(e: React.KeyboardEvent<HTMLInputElement>) => {
          if (e.key === "Enter") handleSendMessage();
        }}
      />
      <Button
        onClick={handleSendMessage}
        className="bg-blue-600 hover:bg-blue-700 text-white"
      >
        <Send className="h-4 w-4" />
      </Button>
    </div>
  );
};

export default ChatInput;
