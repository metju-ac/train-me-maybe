import AccountForm from "@components/AccountForm";
import { CircularProgress } from "@mui/material";
import useUserDetails from "@utils/useUserDetails";

export default function Account() {
  const { data, isLoading, isError } = useUserDetails();

  if (isLoading) {
    return <CircularProgress />;
  }

  if (isError) {
    return <div>Error while loading data</div>;
  }

  return <AccountForm user={data!} />;
}
