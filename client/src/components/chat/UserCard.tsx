import React from "react";
import { Card, CardContent, CardFooter } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

interface User {
  id: string;
  username: string;
  email: string;
  profile_pic: string;
}

interface UserCardProps {
  user: User;
  onSendMessage: (userId: string) => void;
}

export const UserCard: React.FC<UserCardProps> = ({ user, onSendMessage }) => {
  return (
    <Card className="w-full max-w-sm">
      <CardContent className="flex items-center space-x-4 p-6">
        <Avatar className="h-12 w-12">
          <AvatarImage src={user.profile_pic} alt={user.username} />
          <AvatarFallback>
            {user.username.slice(0, 2).toUpperCase()}
          </AvatarFallback>
        </Avatar>
        <div className="flex-1 space-y-1">
          <h3 className="font-semibold">{user.username}</h3>
          <p className="text-sm text-muted-foreground">{user.email}</p>
        </div>
      </CardContent>
      <CardFooter>
        <Button
          className="w-full bg-neutral-800 hover:bg-neutral-700"
          onClick={() => onSendMessage(user.id)}
        >
          Send Message
        </Button>
      </CardFooter>
    </Card>
  );
};
