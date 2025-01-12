import ToastBarContext from "@components/ToastBarContext";
import watchedRouteService from "@services/watchedRouteService";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router";
import printError from "./printError";

export default function useCreateWatchedRoute() {
  const queryClient = useQueryClient();
  const { setToast } = useContext(ToastBarContext);
  const navigate = useNavigate();
  const { t } = useTranslation("default");

  return useMutation({
    mutationKey: [watchedRouteService.createWatchedRoute.key],
    mutationFn: watchedRouteService.createWatchedRoute.fn,
    onSuccess: () => {
      // Invalidate and refetch
      void queryClient.invalidateQueries({
        queryKey: [watchedRouteService.getWatchedRoutes.key],
      });
      setToast(t("Route created successfully!"), "success");
      void navigate("/");
    },
    onError: (err) => {
      printError(setToast, err, t("Failed to create route"));
    },
  });
}
