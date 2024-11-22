import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

const API_URL = process.env.NEXT_PUBLIC_API_URL;

export const chatApi = createApi({
  reducerPath: "chatApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_URL }),
  tagTypes: ["Chat"],
  endpoints: (builder) => ({
    getChatById: builder.query({
      query: (id) => `/chats/${id}`,
      providesTags: (result, error, id) => [{ type: "Chat", id }],
    }),
    sendMessage: builder.mutation({
      query: (message) => ({
        url: "/send-message",
        method: "POST",
        body: message,
      }),
      invalidatesTags: (result, error, { chat_id }) => [
        { type: "Chat", id: chat_id },
      ],
    }),
  }),
});

export const { useGetChatByIdQuery, useSendMessageMutation } = chatApi;
