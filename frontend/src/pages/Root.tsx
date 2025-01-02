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
      alignItems="center"
      sx={{ height: "100vh" }}
    >
      <Box component="header" marginY={2}>
        <Navigation />
      </Box>
      <main>
        <Outlet />
      </main>
      <footer>
        <Footer />
      </footer>
    </Grid2>
  );
}
