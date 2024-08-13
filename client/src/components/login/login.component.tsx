import Button from "../button/button.component";
import "./login.style.css";

const Login = () => {
  return (
    <div id="login-modal-wrapper">
      <article id="login-form">
        <h2>Login</h2>
        <form>
          <fieldset>
            <legend>Email</legend>
            <input id="email" name="email" type="email" />
          </fieldset>
          <fieldset>
            <legend>Password</legend>
            <input id="password" name="password" type="password" />
          </fieldset>
          <div id="login-button-wrapper">
            <Button id="login-cancel-btn" type="button" text="Cancel" className="secondary" />
            <Button id="login-submit-btn" type="button" text="Login" className="primary" />
          </div>
        </form>
      </article>
    </div>
  )
};

export default Login;