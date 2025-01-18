import {
  AvailableLanguage,
  availableLanguages,
  languageToFlagUrl,
} from "@/i18n";
import { Autocomplete, AutocompleteProps, Box, TextField } from "@mui/material";
import i18next from "i18next";

export const LanguageSwitcher = (
  props: Partial<AutocompleteProps<AvailableLanguage, false, true, false>>
) => {
  // const { t } = useTranslation("default");

  const currentLanguage = i18next.language as AvailableLanguage;

  return (
    <Autocomplete
      {...props}
      disablePortal
      disableClearable
      id="language-switcher"
      // title={t("Choose language")}
      options={availableLanguages}
      value={currentLanguage}
      onChange={(_, value) => {
        void i18next.changeLanguage(value ?? undefined);
      }}
      renderInput={(params) => (
        <TextField
          {...params}
          // label={t("Choose language")}
          slotProps={{
            input: {
              ...params.InputProps,
              startAdornment: (
                <img
                  loading="lazy"
                  width="20"
                  src={languageToFlagUrl[currentLanguage]}
                />
              ),
              style: {
                paddingTop: 0,
                paddingBottom: 0,
                color: "white",
              },
            },
          }}
        />
      )}
      renderOption={(props, option) => (
        <Box
          component="li"
          sx={{ "& > img": { mr: 2, flexShrink: 0 } }}
          {...props}
        >
          <img loading="lazy" width="20" src={languageToFlagUrl[option]} />
          {option}
        </Box>
      )}
    />
  );
};
