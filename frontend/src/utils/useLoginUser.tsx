import { ToastBarContext } from "@components/ToastBarProvider";
import userService from "@services/userService";
import { useMutation } from "@tanstack/react-query";
import { useContext } from "react";
import { useNavigate } from "react-router";

export default function useLoginUser() {
  const { setToast } = useContext(ToastBarContext);
  let navigate = useNavigate();

  return useMutation({
    mutationKey: [userService.loginUser.key],
    mutationFn: userService.loginUser.fn,
    onSuccess: (data) => {
      console.log("User logged in successfully", data);
      setToast("Successfully logged in", "success");
      navigate("/account");
    },
    onError: (error) => {
      console.error("Error logging in", error);
      setToast("Error logging in: " + error.message, "error");
    },
  });
}
