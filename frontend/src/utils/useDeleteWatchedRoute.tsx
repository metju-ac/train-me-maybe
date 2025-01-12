import ToastBarContext from "@components/ToastBarContext";
import watchedRouteService from "@services/watchedRouteService";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import printError from "./printError";

export default function useDeleteWatchedRoute() {
  const queryClient = useQueryClient();

  const { setToast } = useContext(ToastBarContext);

  return useMutation({
    mutationKey: [watchedRouteService.deleteWatchedRoute.key],
    mutationFn: watchedRouteService.deleteWatchedRoute.fn,
    onSuccess: () => {
      // Invalidate and refetch
      void queryClient.invalidateQueries({
        queryKey: [watchedRouteService.getWatchedRoutes.key],
      });

      setToast("Route deleted successfully!", "success");
    },
    onError: (err) => {
      printError(setToast, err, "Failed to delete route");
    },
  });
}
