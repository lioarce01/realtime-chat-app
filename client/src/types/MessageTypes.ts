export interface Message {
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

export interface MessagesProps {
  messages: Message[];
  dbUserId: string | undefined;
}
