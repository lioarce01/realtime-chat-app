export interface ReceiverProfileProps {
  receiverData: any;
  isLoading: any;
}

export interface SearchbarProps {
  onCreateChat: (userId: string) => void;
}

export interface SidebarProps {
  chats?: any[];
  userId: string;
  setSelectedChatId: (chatId: string | null) => void;
  selectedChatId: string | null;
  dbUser?: any;
  onCreateChat: (otherUserId: string) => void;
}

export interface UserMenuProps {
  dbUser: any;
}
