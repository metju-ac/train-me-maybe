import { NavLink } from "react-router";

import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";

/**
 *  See https://reactrouter.com/start/library/navigating
 */
export default function ButtonAppBar() {
  return (
    <nav>
      <Box sx={{ flexGrow: 1 }}>
        <AppBar position="static">
          <Toolbar
            sx={{
              display: "flex",
              justifyContent: "space-between",
              minWidth: "800px",
            }}
          >
            <Box sx={{ display: "flex", gap: "1rem" }}>
              <NavLink to="/" end>
                Home
              </NavLink>
              <NavLink to="/about">About</NavLink>
            </Box>
            <NavLink to="/account">Account</NavLink>
          </Toolbar>
        </AppBar>
      </Box>
    </nav>
  );
}
