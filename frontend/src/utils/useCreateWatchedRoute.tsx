import ToastBarContext from "@components/ToastBarContext";
import watchedRouteService from "@services/watchedRouteService";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import { useNavigate } from "react-router";

export default function useCreateWatchedRoute() {
  const queryClient = useQueryClient();
  const { setToast } = useContext(ToastBarContext);
  const navigate = useNavigate();

  return useMutation({
    mutationKey: [watchedRouteService.createWatchedRoute.key],
    mutationFn: watchedRouteService.createWatchedRoute.fn,
    onSuccess: () => {
      // Invalidate and refetch
      void queryClient.invalidateQueries({
        queryKey: [watchedRouteService.getWatchedRoutes.key],
      });
      setToast("Route created successfully!", "success");
      void navigate("/");
    },
    onError: (err) => {
      setToast("Failed to create route: " + err.message, "error");
    },
  });
}
