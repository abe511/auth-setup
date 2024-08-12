import Login from "../login/login-component";
import "./main-content-style.css";


const MainContent = () => {
  return (
    <main>
      <div id="login-modal-wrapper">
        <Login />
      </div>
      <div id="main-left"></div>
      <div id="main-top"></div>
      <div id="main-center"></div>
      <div id="main-right"></div>
      <div id="main-bottom"></div>
    </main>
  )
};

export default MainContent;