import { Button, Tooltip } from "@mui/material";
import { PropsWithChildren } from "react";

export default function SubmitButton({
  isValid,
  tooltipTitle,
  children,
}: PropsWithChildren<{
  isValid: boolean;
  tooltipTitle?: string;
}>) {
  return (
    <Tooltip arrow title={isValid ? "" : tooltipTitle}>
      <span>
        <Button
          type="submit"
          variant="contained"
          color="primary"
          disabled={!isValid}
        >
          {children}
        </Button>
      </span>
    </Tooltip>
  );
}
