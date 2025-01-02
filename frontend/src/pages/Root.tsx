import Footer from "@components/Footer";
import Navigation from "@components/Navigation";
import { Box, Grid2 } from "@mui/material";
import { Outlet } from "react-router";

/**
 * See https://reactrouter.com/start/library/routing#index-routes
 */
export default function Root() {
  return (
    <Grid2
      container
      direction="column"
      justifyContent="space-between"
      alignItems="flex-start"
      padding={2}
      minHeight="100vh"
      sx={{ backgroundColor: "#f6f6f6" }}
    >
      <Box component="header" marginY={2} width="100%">
        <Navigation />
      </Box>
      <Box
        component="main"
        sx={{ flexGrow: 1, width: "100%", overflow: "scroll" }}
      >
        <Outlet />
      </Box>
      <Box component="footer" alignSelf="center">
        <Footer />
      </Box>
    </Grid2>
  );
}
