import { createBrowserRouter } from "react-router-dom";
import Home from "./components/home/home.component";
import About from "./components/pages/about/about.component";
import Contact from "./components/pages/contact/contact.component";
import Login from "./components/login/login.component";
import UserProfile from "./components/pages/user-profile/user-profile.component";
import RootLayout from "./components/root-layout/root-layout.component";
import loginAction from "./routes/loginAction";
import userProfileLoader from "./routes/userProfileLoader";


const router = createBrowserRouter([
  {
    path: "/",
    element: <RootLayout />,
    children: [
      {
        index: true,
        element: <Home />
      },
      {
        path: "contact",
        element: <Contact />
      },
      {
        path: "about",
        element: <About />
      },
      {
        path: "login",
        element: <Login />,
        action: loginAction,
      },
      {
        path: "profile",
        element: <UserProfile />,
        loader: userProfileLoader,
      }
    ]
  }
]);

export default router;