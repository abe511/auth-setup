import { redirect } from "react-router-dom";
import { getUserProfile } from "../api/auth";


  const userProfileLoader = async () => {
  const token = localStorage.getItem("bearerToken");
  if (!token) {
    return redirect("/");
  }

  console.log("profile loader token:", token); // REMOVE

  try {
    const response = await getUserProfile(token);
    return response;
  } catch(err) {
    console.log("profile loader error:", err); // REMOVE
    localStorage.removeItem("bearerToken");
    return redirect("/");
  }
};


export default userProfileLoader;