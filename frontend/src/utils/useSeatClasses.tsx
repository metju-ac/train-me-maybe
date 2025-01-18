import staticDataService from "@services/staticDataService";
import { useQuery } from "@tanstack/react-query";

type Args = Parameters<typeof staticDataService.getSeatClasses.fn>[0];

export default function useSeatClasses(
  params: Omit<Args, "routeId"> & { routeId?: Args["routeId"] | null }
) {
  return useQuery({
    queryKey: [
      staticDataService.getSeatClasses.key,
      params.fromStationId,
      params.toStationId,
      params.routeId,
    ],
    queryFn: () => staticDataService.getSeatClasses.fn(params as Args),
    enabled: Boolean(
      params.fromStationId && params.toStationId && params.routeId
    ),
  });
}
