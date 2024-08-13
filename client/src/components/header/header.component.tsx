import { Link } from "react-router-dom";
import Button from "../button/button.component";
import UserIconSvg from "./user-icon-svg.component";

import "./header.style.css";

const Header = () => {
  return (
    <header>
      <Link to="profile">
        <UserIconSvg />
      </Link>
      <nav>
        <Link to="/">
          <Button id="nav-btn-home" type="button" text="Home" className="nav-btn secondary active" />
        </Link>
        <Link to="contact">
          <Button id="nav-btn-contact" type="button" text="Contact" className="nav-btn secondary" />
        </Link>
        <Link to="about">
          <Button id="nav-btn-about" type="button" text="About" className="nav-btn secondary" />
        </Link>
      </nav>
      <Link to="login">
        <Button type="button" text="Login" className="primary" />
      </Link>
    </header>
  )
};

export default Header;