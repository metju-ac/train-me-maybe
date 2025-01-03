import { ToastBarContext } from "@components/ToastBarProvider";
import userService from "@services/userService";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";

export default function useUpdateUserDetails() {
  const { setToast } = useContext(ToastBarContext);

  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: [userService.updateUserDetails.key],
    mutationFn: userService.updateUserDetails.fn,
    onSuccess: () => {
      // Invalidate and refetch
      queryClient.invalidateQueries({
        queryKey: [userService.getUserDetails.key],
      });

      setToast("Details updated successfully!", "success");
    },
    onError: () => {
      setToast("Failed to update details!", "error");
    },
  });
}
