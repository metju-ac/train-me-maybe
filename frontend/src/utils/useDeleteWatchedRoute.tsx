import ToastBarContext from "@components/ToastBarContext";
import watchedRouteService from "@services/watchedRouteService";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import { useTranslation } from "react-i18next";
import printError from "./printError";

export default function useDeleteWatchedRoute() {
  const queryClient = useQueryClient();

  const { setToast } = useContext(ToastBarContext);

  const { t } = useTranslation("default");

  return useMutation({
    mutationKey: [watchedRouteService.deleteWatchedRoute.key],
    mutationFn: watchedRouteService.deleteWatchedRoute.fn,
    onSuccess: () => {
      // Invalidate and refetch
      void queryClient.invalidateQueries({
        queryKey: [watchedRouteService.getWatchedRoutes.key],
      });

      setToast(t("Route deleted successfully!"), "success");
    },
    onError: (err) => {
      printError(setToast, err, t("Failed to delete route"));
    },
  });
}
