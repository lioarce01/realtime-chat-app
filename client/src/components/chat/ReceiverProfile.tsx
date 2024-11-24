import React from "react";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";
import { ReceiverProfileProps } from "@/types/UserTypes";

const ReceiverProfile: React.FC<ReceiverProfileProps> = ({
  receiverData,
  isLoading,
}) => {
  return (
    <div className="p-4 bg-neutral-900 text-white flex items-center justify-between border-b border-neutral-800">
      {isLoading ? (
        <div className="flex items-center space-x-3 animate-pulse">
          <div className="w-10 h-10 bg-neutral-700 rounded-full"></div>
          <div className="h-4 bg-neutral-700 rounded w-24"></div>
        </div>
      ) : (
        <div className="flex items-center space-x-3">
          <Avatar>
            <AvatarImage
              src={
                receiverData?.profile_pic ||
                `https://api.dicebear.com/6.x/initials/svg?seed=${receiverData?.username}`
              }
              alt={`${receiverData?.username}'s profile picture`}
            />
            <AvatarFallback>
              {receiverData?.username?.charAt(0).toUpperCase()}
            </AvatarFallback>
          </Avatar>
          <h1 className="text-lg font-semibold">{receiverData?.username}</h1>
        </div>
      )}
      <span className="text-sm text-neutral-400">
        {isLoading ? (
          <div className="h-4 bg-neutral-700 rounded w-12 animate-pulse"></div>
        ) : (
          "Online"
        )}
      </span>
    </div>
  );
};

export default ReceiverProfile;
