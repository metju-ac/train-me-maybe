import {
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  SelectProps,
} from "@mui/material";
import { Tariff } from "@services/staticDataService";
import { useTranslation } from "react-i18next";

export default function SelectTariff({
  tariffs,
  selectedTariff,
  setSelectedTariff,
  ...props
}: {
  tariffs: Tariff[];
  selectedTariff: Tariff | null;
  setSelectedTariff: (tariff: Tariff | null) => void;
} & SelectProps<string>) {
  const { t } = useTranslation("default");

  return (
    <FormControl fullWidth>
      <InputLabel id="select-tariff">{t("Tariff")}</InputLabel>
      <Select
        labelId="select-tariff"
        value={selectedTariff?.key ?? ""}
        label={t("Tariff")}
        onChange={(e) => {
          setSelectedTariff(
            tariffs.find((t) => t.key === e.target.value) ?? null
          );
        }}
        {...props}
      >
        {tariffs.map((tariff) => (
          <MenuItem key={tariff.key} value={tariff.key}>
            {tariff.key} ({tariff.value})
          </MenuItem>
        ))}
      </Select>
    </FormControl>
  );
}
