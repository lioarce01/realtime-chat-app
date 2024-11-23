export interface ReceiverProfileProps {
  receiverData: any;
}

export interface SidebarProps {
  chats?: any[];
  userId: string;
  setSelectedChatId: (chatId: string | null) => void;
  selectedChatId: string | null;
}
