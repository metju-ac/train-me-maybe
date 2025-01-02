import watchedRouteService from "@services/watchedRouteService";
import { useQuery } from "@tanstack/react-query";

export default function useWatchedRoutes() {
  return useQuery({
    queryKey: [watchedRouteService.getWatchedRoutes.key],
    queryFn: watchedRouteService.getWatchedRoutes.fn,
  });
}
