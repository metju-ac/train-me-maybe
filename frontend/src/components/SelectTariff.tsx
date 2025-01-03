import {
  FormControl,
  InputLabel,
  MenuItem,
  Select,
  SelectProps,
} from "@mui/material";
import { Tariff } from "@services/staticDataService";

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
  return (
    <FormControl fullWidth>
      <InputLabel id="select-tariff">Tariff</InputLabel>
      <Select
        labelId="select-tariff"
        value={selectedTariff?.key ?? ""}
        label="Tariff"
        onChange={(e) =>
          { setSelectedTariff(
            tariffs.find((t) => t.key === e.target.value) ?? null
          ); }
        }
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
