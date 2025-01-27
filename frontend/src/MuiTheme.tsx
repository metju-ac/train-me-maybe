import { createTheme } from "@mui/material";

const theme = createTheme({
  components: {
    MuiTooltip: {
      defaultProps: {
        arrow: true,
        enterTouchDelay: 0,
      },
    },
  },
});

export default theme;
