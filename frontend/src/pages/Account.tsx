import AccountForm from "@components/AccountForm";
import userService from "@services/userService";
import { useQuery } from "@tanstack/react-query";
import useIsLoggedIn from "@utils/useIsLoggedIn";

export default function Account() {
  const isLoggedIn = useIsLoggedIn();

  if (!isLoggedIn) {
    return (
      <>
        <h1>You are not logged in</h1>
        <p>To use the app you need to register.</p>
      </>
    );
  }

  const { data, isLoading, isError } = useQuery({
    queryKey: [userService.getUserDetails.key],
    queryFn: userService.getUserDetails.fn,
    retry: 2,
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
