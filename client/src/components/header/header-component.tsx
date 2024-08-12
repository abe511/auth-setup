import Button from "../button/button-component";
import UserIconSvg from "./user-icon-svg-component";
import "./header-style.css"

const Header = () => {
  return (
    <header>
      <UserIconSvg width={48} height={48} />
      <nav>
        <Button id="nav-btn-home" type="button" text="Home" className="nav-btn secondary active" />
        <Button id="nav-btn-contact" type="button" text="Contact" className="nav-btn secondary" />
        <Button id="nav-btn-about" type="button" text="About" className="nav-btn secondary" />
      </nav>
      <Button type="button" text="Login" className="primary" />
    </header>
  )
};

export default Header;