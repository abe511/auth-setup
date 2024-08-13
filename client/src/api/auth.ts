
const API_URL = import.meta.env.VITE_API_URL;

export const loginUser = async (email: string, password: string): Promise<string | Error | null> => {

  console.log("login user request:", email, password);

  try {
    const response = await fetch(`${API_URL}/login`, {
      method: "POST",
      headers: {
        "Content-type": "application/json",
      },
      body: JSON.stringify({email, password}),
    });
    if(!response.ok) {
      throw new Error("Failed to login")
    }
    const data = await response.json();

    console.log("login user.", data); // REMOVE

    return data.accessToken;
  } catch (err) {
    if (err instanceof Error) {
      return err.message;
    }
    return null;
  }
};

export const getUserProfile = async (token: string): Promise<string | Error | null> => {
  try {
    const response = await fetch(`${API_URL}/profile`, {
      headers: {
        "Authorization": token,
      }
    });

    if(!response.ok) {
      throw new Error("Failed to get user profile data")
    }
    const data = await response.json();

    console.log("get user profile.", data); // REMOVE

    return data;
  } catch(err) {
    if (err instanceof Error) {
      return err.message;
    }
    return new Error("Unknown error"); // CHECK THIS RETURN VALUE
  }
}