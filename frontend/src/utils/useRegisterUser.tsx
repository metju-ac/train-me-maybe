import ToastBarContext from "@components/ToastBarContext";
import userService from "@services/userService";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router";
import printError from "./printError";

export default function useRegisterUser() {
  const queryClient = useQueryClient();

  const { setToast } = useContext(ToastBarContext);
  const navigate = useNavigate();

  const { t } = useTranslation("default");

  return useMutation({
    mutationKey: [userService.registerUser.key],
    mutationFn: userService.registerUser.fn,
    onSuccess: (data) => {
      // Invalidate and refetch
      void queryClient.invalidateQueries({
        queryKey: [userService.getUserDetails.key],
      });
      console.log("User registered successfully", data);
      setToast(t("Successfully registered user"), "success");
      void navigate("/account");
    },
    onError: (error) => {
      console.error("Error registering user", error);
      printError(setToast, error, t("Error registering user"));
    },
  });
}
