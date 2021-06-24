import "../styles/index.css";
import React from "react";
import AppBar from "@material-ui/core/AppBar";
import Toolbar from "@material-ui/core/Toolbar";
import logo from "../images/logo3.png";
import { Link } from "react-router-dom";

const NavBar = () => {
  return (
    <header>
      <AppBar position="fixed">
        <Toolbar>
          <div className="logo-div">
            <Link to="/">
              <img src={logo} alt="" />
            </Link>
          </div>
          <div className="button-headers">
            <ul>
              <Link to="/missions">
                <li class="aLinkTag" href="/missions">
                  missions
                </li>
              </Link>
              <Link to="/weapons">
                <li class="aLinkTag" href="/weapons">
                  weapons
                </li>
              </Link>
              <Link to="/vehicles">
                <li class="aLinkTag" href="/vehicles">
                  vehicles
                </li>
              </Link>
              <Link to="/enemies">
                <li class="aLinkTag" href="/enemies">
                  enemies
                </li>
              </Link>
              <Link to="/skulls">
                <li class="aLinkTag" href="/skulls">
                  skulls
                </li>
              </Link>
              <Link to="/maps">
                <li class="aLinkTag" href="">
                  maps
                </li>
              </Link>
            </ul>
          </div>
        </Toolbar>
      </AppBar>
    </header>
  );
};

export default NavBar;
