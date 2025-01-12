import { NavLink } from "react-router";

import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import useIsLoggedIn from "@utils/useIsLoggedIn";
import useScreenWidth from "@utils/useScreenWidth";
import { useTranslation } from "react-i18next";
import { LanguageSwitcher } from "./LanguageSwitcher";

/**
 *  See https://reactrouter.com/start/library/navigating
 */
export default function ButtonAppBar() {
  const isLoggedIn = useIsLoggedIn();
  const { t } = useTranslation("default");
  const width = useScreenWidth();

  return (
    <nav>
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position="static">
          <Toolbar
            sx={{
              display: "flex",
              justifyContent: "space-between",
            }}
          >
            <Box sx={{ display: "flex", gap: "1rem" }}>
              <NavLink to="/" end>
                {t("Home")}
              </NavLink>
              <NavLink to="/about">{t("About")}</NavLink>
            </Box>
            <Box sx={{ display: "flex", gap: "1rem", alignItems: "center" }}>
              {isLoggedIn ? (
                <>
                  {width > 500 ? <LanguageSwitcher /> : null}
                  <NavLink to="/account">{t("Account")}</NavLink>
                  <NavLink to="/logout">{t("Logout")}</NavLink>
                </>
              ) : (
                <>
                  {width > 500 ? <LanguageSwitcher /> : null}
                  <NavLink to="/login">{t("Login")}</NavLink>
                  <NavLink to="/register">{t("Register")}</NavLink>
                </>
              )}
            </Box>
          </Toolbar>
        </AppBar>
      </Box>
    </nav>
  );
}
