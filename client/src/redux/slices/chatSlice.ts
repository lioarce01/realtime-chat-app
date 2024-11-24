import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Message } from "@/types/MessageTypes";

interface ChatState {
  messages: { [chatId: string]: Message[] };
}

const initialState: ChatState = {
  messages: {},
};

const chatSlice = createSlice({
  name: "chat",
  initialState,
  reducers: {
    setMessages: (
      state,
      action: PayloadAction<{ chatId: string; messages: Message[] }>
    ) => {
      const { chatId, messages } = action.payload;
      state.messages[chatId] = messages;
    },
    addMessage: (
      state,
      action: PayloadAction<{ chatId: string; message: Message }>
    ) => {
      const { chatId, message } = action.payload;
      if (!state.messages[chatId]) {
        state.messages[chatId] = [];
      }
      state.messages[chatId].push(message);
    },
    updateMessage: (
      state,
      action: PayloadAction<{
        chatId: string;
        messageId: string;
        updates: Partial<Message>;
      }>
    ) => {
      const { chatId, messageId, updates } = action.payload;
      const messageIndex = state.messages[chatId]?.findIndex(
        (msg) => msg.id === messageId
      );
      if (messageIndex !== undefined && messageIndex !== -1) {
        state.messages[chatId][messageIndex] = {
          ...state.messages[chatId][messageIndex],
          ...updates,
        };
      }
    },
  },
});

export const { setMessages, addMessage, updateMessage } = chatSlice.actions;
export default chatSlice.reducer;
