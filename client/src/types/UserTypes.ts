export interface ReceiverProfileProps {
  receiverData: any;
}

export interface SidebarProps {
  userId: string;
  setSelectedChatId: (chatId: string | null) => void;
  selectedChatId: string | null;
}
