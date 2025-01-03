import AccountForm from "@components/AccountForm";
import { CircularProgress } from "@mui/material";
import useTariffs from "@utils/useTariffs";
import useUserDetails from "@utils/useUserDetails";

export default function Account() {
  const { data, isLoading, isError } = useUserDetails();
  const {
    data: tariffs,
    isLoading: isLoadingTariffs,
    isError: isErrorTariffs,
  } = useTariffs();

  if (isLoading || isLoadingTariffs) {
    return <CircularProgress />;
  }

  if (isError || isErrorTariffs) {
    return <div>Error while loading data</div>;
  }

  return <AccountForm user={data!} tariffs={tariffs!} />;
}
