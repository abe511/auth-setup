import { ActionFunctionArgs, redirect } from "react-router-dom";
import { loginUser } from "../api/auth";


const loginAction = async ({request}: ActionFunctionArgs) => {
  const formData = await request.formData();
  const email = formData.get("email");
  const password = formData.get("password");

console.log("login action. ", "email: ", email, "password: ", password); // REMOVE

try {
    const accessToken = await loginUser(email as string, password as string);
    // const accessToken = await loginUser("test@email.net", "test_pass");
    if(accessToken !== null && !(accessToken instanceof Error)) {
      localStorage.setItem("bearerToken", accessToken);
    }
    return redirect("/profile");
  } catch(err) {
    if (err instanceof Error) {
      return err.message;
    }
    return null;
  }
};

export default loginAction;