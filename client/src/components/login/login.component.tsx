import { Form } from "react-router-dom";
import Button from "../button/button.component";
import "./login.style.css";

const Login = () => {
  return (
    <div id="login-modal-wrapper">
      <article id="login-form">
        <h2>Login</h2>
        <Form method="post">
          <fieldset>
            <legend>Email</legend>
            <input id="email" name="email" type="email" required />
          </fieldset>
          <fieldset>
            <legend>Password</legend>
            <input id="password" name="password" type="password" required />
          </fieldset>
          <div id="login-button-wrapper">
            <Button id="login-cancel-btn" type="reset" text="Cancel" className="secondary" />
            <Button id="login-submit-btn" type="submit" text="Login" className="primary" />
          </div>
        </Form>
      </article>
    </div>
  )
};

export default Login;