import { NavLink } from "react-router";

import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import useIsLoggedIn from "@utils/useIsLoggedIn";
import { useTranslation } from "react-i18next";

/**
 *  See https://reactrouter.com/start/library/navigating
 */
export default function ButtonAppBar() {
  const isLoggedIn = useIsLoggedIn();
  const { t } = useTranslation("default");

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
            <Box sx={{ display: "flex", gap: "1rem" }}>
              {isLoggedIn ? (
                <>
                  <NavLink to="/account">{t("Account")}</NavLink>
                  <NavLink to="/logout">{t("Logout")}</NavLink>
                </>
              ) : (
                <>
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
