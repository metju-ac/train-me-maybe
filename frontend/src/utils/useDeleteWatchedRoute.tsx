import watchedRouteService from "@services/watchedRouteService";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export default function useDeleteWatchedRoute() {
  const queryClient = useQueryClient();

  return useMutation({
    mutationKey: [watchedRouteService.deleteWatchedRoute.key],
    mutationFn: watchedRouteService.deleteWatchedRoute.fn,
    onSuccess: () => {
      // Invalidate and refetch
      queryClient.invalidateQueries({
        queryKey: [watchedRouteService.getWatchedRoutes.key],
      });
    },
  });
}
