import ToastBarContext from "@components/ToastBarContext";
import userService from "@services/userService";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import printError from "./printError";

export default function useUpdateUserDetails() {
  const { setToast } = useContext(ToastBarContext);

  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: [userService.updateUserDetails.key],
    mutationFn: userService.updateUserDetails.fn,
    onSuccess: () => {
      // Invalidate and refetch
      void queryClient.invalidateQueries({
        queryKey: [userService.getUserDetails.key],
      });

      setToast("Details updated successfully!", "success");
    },
    onError: (err) => {
      printError(setToast, err, "Failed to update user details");
    },
  });
}
