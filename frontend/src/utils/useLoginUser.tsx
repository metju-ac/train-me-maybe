import config from "@/config";
import ToastBarContext from "@components/ToastBarContext";
import userService from "@services/userService";
import { useMutation } from "@tanstack/react-query";
import { useContext } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router";
import printError from "./printError";

export default function useLoginUser() {
  const { setToast } = useContext(ToastBarContext);
  const navigate = useNavigate();

  const { t } = useTranslation("default");

  return useMutation({
    mutationKey: [userService.loginUser.key],
    mutationFn: userService.loginUser.fn,
    onSuccess: (data) => {
      console.log("User logged in successfully", data);
      setToast(t("Successfully logged in"), "success");
      void navigate(`/${config.urlPrefix}`);
    },
    onError: (error) => {
      console.error("Error logging in", error);
      printError(setToast, error, t("Error logging in"));
    },
  });
}
