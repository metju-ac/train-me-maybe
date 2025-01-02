import AccountForm from "@components/AccountForm";
import { CircularProgress } from "@mui/material";
import useIsLoggedIn from "@utils/useIsLoggedIn";
import useUserDetails from "@utils/useUserDetails";

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

  const { data, isLoading, isError } = useUserDetails();

  if (isLoading) {
    return <CircularProgress />;
  }

  if (isError) {
    return <div>Error while loading data</div>;
  }

  console.log("data", data);

  return <AccountForm user={data!} />;
}
