import { useNavigate } from "react-router-dom";
import React, { useState } from "react";
import '../styles/Auth.css';
import axios from "axios";


const Auth: React.FC = () => {
  const [credentials, setCredentials] = useState({ username: "", password: "" });
  const [authToken, setAuthToken] = useState<string | null>(null);
  const navigate = useNavigate();


  const handleLogin = async () => {
    try { // TODO: sep api env
      const response = await axios.post("http://localhost:8080/api/v1/auth/login", credentials);
      const data = response.data;

      if (data !== null && typeof data === 'object' && 'success' in data) {
        if (data.success) {
          console.log("Login successful");
          setAuthToken("token_here");
          navigate("/animals");
        } else {
          console.error("Login failed: Invalid credentials");
        }
      } else {
        console.error("Unexpected response format", data);
      }
    } catch (error) {
      console.error("Login failed", error);
    }
  };

  return (
      <div className="auth-container">
        <h1>Login</h1>
        <input
            type="text"
            value={credentials.username}
            onChange={(e) => setCredentials({...credentials, username: e.target.value})}
            placeholder="Username"
        />
        <input
            type="password"
            value={credentials.password}
            onChange={(e) => setCredentials({...credentials, password: e.target.value})}
            placeholder="Password"
        />
        <button onClick={handleLogin}>Login</button>
        {authToken && <p>Authenticated!</p>}
      </div>
  );
};


export default Auth;
