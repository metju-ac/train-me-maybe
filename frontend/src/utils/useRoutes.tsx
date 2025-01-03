import routeService from "@services/routeService";
import { useQuery } from "@tanstack/react-query";

export default function useRoutes(
  props: Parameters<typeof routeService.getRoutes.fn>[0]
) {
  return useQuery({
    queryKey: [
      routeService.getRoutes.key,
      props.fromStation.stationID,
      props.toStation.stationID,
      props.date,
    ],
    queryFn: () => routeService.getRoutes.fn(props),
  });
}
