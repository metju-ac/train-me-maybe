import staticDataService from "@services/staticDataService";
import { useQuery } from "@tanstack/react-query";

export default function useTariffs() {
  return useQuery({
    queryKey: [staticDataService.getTariffs.key],
    queryFn: staticDataService.getTariffs.fn,
  });
}
