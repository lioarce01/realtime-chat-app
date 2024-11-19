import React from "react";

const chats = [
  {
    id: 1,
    sender: "You",
    message: "Hi there! How are you?",
    time: "10:00 AM",
    outgoing: true,
  },
  {
    id: 2,
    sender: "Alice",
    message: "Hey! I'm doing great, thanks for asking 😊",
    time: "10:01 AM",
    outgoing: false,
  },
  {
    id: 3,
    sender: "You",
    message: "Awesome to hear that! Let's catch up soon.",
    time: "10:03 AM",
    outgoing: true,
  },
  {
    id: 4,
    sender: "Alice",
    message: "Sure thing! Talk to you later.",
    time: "10:04 AM",
    outgoing: false,
  },
];

const contacts = [
  { id: 1, name: "Alice", lastMessage: "Sure thing! Talk to you later." },
  { id: 2, name: "Bob", lastMessage: "See you tomorrow!" },
  { id: 3, name: "Charlie", lastMessage: "Can we reschedule our meeting?" },
];

const ChatPage = () => {
  return (
    <div className="flex items-center justify-center w-full h-screen bg-[#0e0e0e] text-white">
      <div className="h-full w-full md:h-[95%] md:w-[80%] shadow-md shadow-neutral-900">
        <div className="flex justify-between w-full h-full">
          {/* Sidebar */}
          <div className="w-1/3 bg-[#1f1f1f] border-r border-neutral-800">
            <div className="p-4">
              <h1 className="text-2xl font-semibold">Chats</h1>
            </div>
            <div className="overflow-y-auto h-[calc(100%-4rem)]">
              {contacts.map((contact) => (
                <div
                  key={contact.id}
                  className="p-4 flex items-center border-b border-neutral-800 w-full space-x-3 cursor-pointer hover:bg-[#242424] transition-all duration-300"
                >
                  <div className="bg-neutral-700 rounded-full w-12 h-12">
                    {/*imagen */}
                  </div>
                  <div>
                    <h2 className="font-bold">{contact.name}</h2>
                    <p className="text-sm text-[#b0b1b1] truncate">
                      {contact.lastMessage}
                    </p>
                  </div>
                </div>
              ))}
            </div>
          </div>
          {/* Chat Window */}
          <div className="flex-1 flex flex-col">
            {/* Header */}
            <div className="p-4 bg-[#242424] text-white flex items-center justify-between border-b border-neutral-700">
              <h1 className="text-lg font-semibold">Alice</h1>
              <span className="text-sm text-gray-100">Online</span>
            </div>

            {/* Chat Messages */}
            <div className="flex-1 overflow-y-auto p-4 bg-[#242424]">
              {chats.map((chat) => (
                <div
                  key={chat.id}
                  className={`mb-4 flex ${
                    chat.outgoing ? "justify-end" : "justify-start"
                  }`}
                >
                  <div
                    className={`max-w-sm p-3 rounded-lg shadow ${
                      chat.outgoing
                        ? "bg-[#EEF5FC] text-[#333540]"
                        : "bg-[#F4F5F7] text-[#333540]"
                    }`}
                  >
                    <p className="text-sm">{chat.message}</p>
                    <p className="text-xs text-[#707880] text-right mt-1">
                      {chat.time}
                    </p>
                  </div>
                </div>
              ))}
            </div>

            {/* Input Bar */}
            <div className="p-4 bg-[#242424]  flex items-center">
              <input
                type="text"
                placeholder="Type a message..."
                className="flex-1 p-2 border border-neutral-700 rounded-lg outline-none bg-[#242424] text-neutral-100 focus:ring-2 focus:ring-neutral-700"
              />
              <button className="ml-4 px-4 py-2 bg-[#dfdfdf] text-neutral-800 rounded-lg hover:bg-[#b8b7b7] transition-all duration-300">
                Send
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ChatPage;
