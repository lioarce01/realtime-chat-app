import React, { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useGetAllUsersQuery } from "@/redux/api/userApi";
import { UserCard } from "@/components/chat/UserCard";
import { Skeleton } from "@/components/ui/skeleton";
import { SearchbarProps } from "@/types/UserTypes";

const Searchbar: React.FC<SearchbarProps> = ({ onCreateChat }) => {
  const [search, setSearch] = useState("");
  const [triggerSearch, setTriggerSearch] = useState(false);
  const { data, isLoading, error } = useGetAllUsersQuery(search, {
    skip: !triggerSearch,
  });

  const handleSearchClick = () => {
    if (search.trim() !== "") {
      setTriggerSearch(true);
    }
  };

  const handleClearClick = () => {
    setSearch("");
    setTriggerSearch(false);
  };

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setSearch(e.target.value);
    setTriggerSearch(false);
  };

  const handleSendMessage = (userId: string) => {
    onCreateChat(userId);
  };

  return (
    <div className="w-full max-w-2xl mx-auto flex flex-col items-center justify-center px-6 space-y-4">
      <div className="flex w-full space-x-4">
        <Input
          type="text"
          value={search}
          onChange={handleInputChange}
          placeholder="Search user"
          className="border-neutral-700 focus:ring-neutral-700"
        />
        <Button
          className="bg-neutral-800 hover:bg-neutral-700"
          onClick={handleSearchClick}
        >
          Search user
        </Button>
        <Button
          className="bg-neutral-800 hover:bg-neutral-700"
          onClick={handleClearClick}
        >
          Clear
        </Button>
      </div>
      <div className="w-full">
        {isLoading && (
          <div className="space-y-3">
            <Skeleton className="h-12 w-full" />
            <Skeleton className="h-4 w-full" />
            <Skeleton className="h-4 w-3/4" />
          </div>
        )}
        {error && (
          <p className="text-red-500">Error: try with another username</p>
        )}
        {data && data.users === null && triggerSearch && (
          <p className="text-center">No users found.</p>
        )}
        {data && data.users && data.users.length > 0 && (
          <div className="space-y-4 overflow-y-auto max-h-72 px-2">
            {data.users.map((user: any) => (
              <UserCard
                key={user.id}
                user={user}
                onSendMessage={handleSendMessage}
              />
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default Searchbar;
