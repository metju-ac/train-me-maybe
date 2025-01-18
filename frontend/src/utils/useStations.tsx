import { AvailableLanguage } from "@/i18n";
import staticDataService from "@services/staticDataService";
import { useQuery } from "@tanstack/react-query";
import { useTranslation } from "react-i18next";

export default function useStations() {
  const { i18n } = useTranslation();
  const lang = i18n.language;

  return useQuery({
    queryKey: [staticDataService.getStations.key, lang],
    queryFn: () => staticDataService.getStations.fn(lang as AvailableLanguage),
  });
}
