import AccountForm from "@components/AccountForm";
import { CircularProgress } from "@mui/material";
import useTariffs from "@utils/useTariffs";
import useUserDetails from "@utils/useUserDetails";
import { useTranslation } from "react-i18next";

export default function Account() {
  const { t } = useTranslation("default");

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
    return <div>{t("Error while loading data")}</div>;
  }

  return <AccountForm user={data!} tariffs={tariffs!} />;
}
