import { ToastBarContext } from "@components/ToastBarProvider";
import userService from "@services/userService";
import { useMutation } from "@tanstack/react-query";
import { useContext } from "react";

export default function useUpdateUserDetails() {
  const { setToast } = useContext(ToastBarContext);

  return useMutation({
    mutationKey: [userService.updateUserDetails.key],
    mutationFn: userService.updateUserDetails.fn,
    onSuccess: () => {
      setToast("Details updated successfully!", "success");
    },
    onError: () => {
      setToast("Failed to update details!", "error");
    },
  });
}
