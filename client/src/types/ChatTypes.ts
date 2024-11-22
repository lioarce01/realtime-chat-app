export interface ChatProps {
  chatId: string | null;
  dbUserId: string | undefined;
}

export interface ChatInputProps {
  inputMessage: string;
  setInputMessage: React.Dispatch<React.SetStateAction<string>>;
  handleSendMessage: () => void;
}
