import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import NavBar from "./components/NavBar";
import Maps from "./components/Maps";
import Missions from "./components/Missions";
import Inicio from "./components/Inicio";
import Enemies from "./components/Enemies";
import Weapons from "./components/Weapons";
import Vehicles from "./components/Vehicles";
import Skulls from "./components/Skulls";

function App() {
  return (
    <Router>
      <div className="App">
        <NavBar />

        <Switch>
          <Route path="/" exact component={Inicio} />
          <Route path="/missions" component={Missions} />
          <Route path="/weapons" component={Weapons} />
          <Route path="/vehicles" component={Vehicles} />
          <Route path="/enemies" component={Enemies} />
          <Route path="/skulls" component={Skulls} />
          <Route path="/maps" component={Maps} />
        </Switch>
      </div>
    </Router>
  );
}

export default App;
