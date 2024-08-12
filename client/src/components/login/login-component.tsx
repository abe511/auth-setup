import "./login-style.css";

const Login = () => {
  return (
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
          <button id="login-cancel-btn"type="button">Cancel</button>
          <button id="login-submit-btn" type="button">Login</button>
        </div>
      </form>
    </article>
  )
};

export default Login;