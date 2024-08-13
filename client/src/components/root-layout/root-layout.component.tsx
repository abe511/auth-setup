import { Outlet } from "react-router-dom";
import Header from "../header/header.component";
import Footer from "../footer/footer.component";
import Background from "./background.component";

import "./root-layout.style.css";

function RootLayout() {
  return (
    <>
      <Header />
      <main>
        <Background />
        <div id="content">
          <Outlet />
        </div>
      </main>
      <Footer />
    </>
  )
}

export default RootLayout;
