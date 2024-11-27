import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

const API_URL = process.env.NEXT_PUBLIC_API_URL;

export const userApi = createApi({
  reducerPath: "userApi",
  baseQuery: fetchBaseQuery({ baseUrl: API_URL, credentials: "include" }),
  tagTypes: ["User"],
  endpoints: (builder) => ({
    getAllUsers: builder.query({
      query: (query) => {
        const queryString = query ? `?username=${query}` : "";
        return `/users${queryString}`;
      },
    }),
    registerUser: builder.mutation({
      query: (user) => {
        return {
          url: "/register",
          method: "POST",
          body: user,
        };
      },
    }),
    getUserById: builder.query({
      query: (id) => `/users/${id}`,
    }),
    getUserChats: builder.query({
      query: (id) => `/users/${id}/chats`,
      providesTags: (result, error, id) => [{ type: "User", id }],
    }),
  }),
});

export const {
  useGetAllUsersQuery,
  useRegisterUserMutation,
  useGetUserByIdQuery,
  useGetUserChatsQuery,
  useLazyGetUserByIdQuery,
} = userApi;
