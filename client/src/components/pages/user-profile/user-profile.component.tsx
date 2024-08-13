import { useLoaderData } from "react-router-dom";

type User = {
  email: string,
  username: string,
}

const UserProfile = () => {
  const userData = useLoaderData() as User;

  return (
    <section>
      <p>{userData.username}</p>
      <p>{userData.email}</p>
    </section>
  )
}

export default UserProfile;