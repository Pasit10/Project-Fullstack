import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import axiosInstant from "../../utils/axios";

const Home = () => {
  const [userData, setUserData] = useState(null);
  const [error, setError] = useState("");
  const navigate = useNavigate();

  useEffect(() => {
    const fetchProtectedData = async () => {
      try {
        // const user = auth.currentUser;
        // if (!user) {
        //   navigate("/"); // Redirect if not logged in
        //   return;
        // }

        // Get Firebase JWT token
        // const token = await getIdToken(user);

        // Fetch protected data from FastAPI
        const response = await axiosInstant("/user", {
          withCredentials: true,
        })

        if (response.status !== 200) throw new Error("Failed to fetch protected data");

        const data = await response.data;
        setUserData(data);
      } catch (err) {
        console.error("Error:", err);
        setError("Unauthorized access. Please log in again.");
      }
    };

    fetchProtectedData();
  }, [navigate]);

  // Logout function
  const handleLogout = async () => {
    navigate("/"); // Redirect to login
    const response = await axiosInstant.get("/logout" , {withCredentials:true})
    if (response.status !== 200) {
      throw new Error("Failed to fetch protected data");
    }
  };

  return (
    <div>
      <h2>Home Page</h2>
      {error ? (
        <p style={{ color: "red" }}>{error}</p>
      ) : userData ? (
        <div>
          <p>Welcome, authenticated user!</p>
          <pre>{JSON.stringify(userData, null, 2)}</pre>
        </div>
      ) : (
        <p>Loading...</p>
      )}
      <button onClick={handleLogout}>Logout</button>
    </div>
  );
};

export default Home;
