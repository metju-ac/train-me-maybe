import ToastBarContext from "@components/ToastBarContext";
import userService from "@services/userService";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import { useNavigate } from "react-router";

export default function useRegisterUser() {
  const queryClient = useQueryClient();

  const { setToast } = useContext(ToastBarContext);
  const navigate = useNavigate();

  return useMutation({
    mutationKey: [userService.registerUser.key],
    mutationFn: userService.registerUser.fn,
    onSuccess: (data) => {
      // Invalidate and refetch
      void queryClient.invalidateQueries({
        queryKey: [userService.getUserDetails.key],
      });
      console.log("User registered successfully", data);
      setToast("Successfully registered user", "success");
      void navigate("/account");
    },
    onError: (error) => {
      console.error("Error registering user", error);
      setToast("Error registering user: " + error.message, "error");
    },
  });
}
