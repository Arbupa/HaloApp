import Button from "@material-ui/core/Button";
import LogoutIcon from "@material-ui/icons/ExitToApp";
import React from "react";
import "../styles/MaterialStyle.css";

const Material = () => {
  return (
    <div>
      <h2>Funcionando</h2>
      <div>
        <Button
          variant="contained"
          color="primary"
          onClick={() => alert("Wasuuuuuuuup")}
        >
          Hola, prrus
        </Button>
      </div>
      <Button
        id="botonOut"
        startIcon={<LogoutIcon />}
        color="secondary"
        variant="contained"
      >
        Logout
      </Button>
    </div>
  );
};

export default Material;
