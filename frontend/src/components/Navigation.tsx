import { NavLink } from "react-router";

import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import { getCookie, USER_COOKIE } from "@utils/getCookie";

/**
 *  See https://reactrouter.com/start/library/navigating
 */
export default function ButtonAppBar() {
  const isLoggedIn = getCookie(USER_COOKIE) !== null;

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
                Home
              </NavLink>
              <NavLink to="/about">About</NavLink>
            </Box>
            <Box sx={{ display: "flex", gap: "1rem" }}>
              {isLoggedIn ? (
                <>
                  <NavLink to="/account">Account</NavLink>
                  <NavLink to="/logout">Logout</NavLink>
                </>
              ) : (
                <>
                  <NavLink to="/login">Login</NavLink>
                  <NavLink to="/register">Register</NavLink>
                </>
              )}
            </Box>
          </Toolbar>
        </AppBar>
      </Box>
    </nav>
  );
}
