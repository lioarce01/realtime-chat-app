import React from "react";
import { Avatar, AvatarFallback, AvatarImage } from "../ui/avatar";

interface ReceiverProfileProps {
  receiverData: any;
}

const ReceiverProfile: React.FC<ReceiverProfileProps> = ({ receiverData }) => {
  return (
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
  );
};

export default ReceiverProfile;
