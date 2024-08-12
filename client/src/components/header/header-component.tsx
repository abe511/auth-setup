import Button from "../button/button-component";
import "./header-style.css"
import UserIconSvg from "./user-icon-svg-component";

const Header = () => {
  return (
    <header>
      <UserIconSvg width={48} height={48} />
      <nav>
        <button>Home</button>
        <button>Contact</button>
        <button>About</button>
      </nav>
      <Button type="button" text="Login" className="primary" />
    </header>
  )
};

export default Header;