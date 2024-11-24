export interface ReceiverProfileProps {
  receiverData: any;
  isLoading: any;
}

export interface SidebarProps {
  chats?: any[];
  userId: string;
  setSelectedChatId: (chatId: string | null) => void;
  selectedChatId: string | null;
  dbUser?: any;
}

export interface UserMenuProps {
  dbUser: any;
}
