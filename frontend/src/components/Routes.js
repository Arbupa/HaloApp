import React from "react";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import Counter from "./Counter";

export default function Routes() {
  return (
    <Router>
      <div>
        <nav>
          <ul>
            <li>
              <Link to="/">Halo App</Link>
            </li>
            <li>
              <Link to="/missions">Missions</Link>
            </li>
            <li>
              <Link to="/weapons">Weapons</Link>
            </li>
            <li>
              <Link to="/vehicles"> Vehicles</Link>
            </li>
            <li>
              <Link to="/enemies">Enemies</Link>
            </li>
            <li>
              <Link to="/skulls">Skulls</Link>
            </li>
            <li>
              <Link to="/maps">Maps</Link>
            </li>
          </ul>
        </nav>

        {/* A <Switch> looks through its children <Route>s and
            renders the first one that matches the current URL. */}
        <Switch>
          <Route path="/">
            <Home />
          </Route>
          <Route path="/missions">
            <Counter />
          </Route>
          <Route path="/maps">
            <Users />
          </Route>
          <Route path="/weapons">
            <Users />
          </Route>
          <Route path="/vehicles">
            <Users />
          </Route>
          <Route path="/enemies">
            <Users />
          </Route>
          <Route path="/skulls">
            <Users />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

function Home() {
  // Aquí agregar los respectivos componentes que renderizan su parte
  return <h2>Home</h2>;
}

function About() {
  return <h2>About</h2>;
}

function Users() {
  return <h2>Users</h2>;
}
