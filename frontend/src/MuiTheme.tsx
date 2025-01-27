import { createTheme } from "@mui/material";

const theme = createTheme({
  components: {
    MuiTooltip: {
      defaultProps: {
        arrow: true,
        enterTouchDelay: 0,
        leaveTouchDelay: 5000,
      },
    },
  },
});

export default theme;
