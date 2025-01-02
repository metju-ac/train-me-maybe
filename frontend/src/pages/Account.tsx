import AccountForm from "@components/AccountForm";
import userService from "@services/userService";
import { useQuery } from "@tanstack/react-query";

export default function Account() {
  const { data, isLoading, isError } = useQuery({
    queryKey: [userService.getUserDetails.key],
    queryFn: userService.getUserDetails.fn,
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error while loading data</div>;
  }

  console.log("data", data);

  return <AccountForm user={data!} />;
}
