import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

function Login() {
  const [message, setMessage] = useState("");
  const [formData, setFormData] = useState({
    username: "",
    password: "",
    yourotp: "",
  });

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData((prevState) => ({ ...prevState, [name]: value }));
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    console.log(formData);

    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        username: formData.username,
        password: formData.password,
        yourotp: formData.yourotp,
      }),
    };
    fetch("http://localhost:8080/login", requestOptions)
      .then((response) => response.json())
      .then((result) => {
        alert(result);
        setMessage(result.status);
      });
  };

  return (
    <>
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src={viteLogo} className="logo" alt="Vite logo" />
        </a>
        <a href="https://react.dev" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Login</h1>
      <h3 style={{ color: "red" }}>{message}</h3>
      <form onSubmit={handleSubmit}>
        <div className="card">
          <div className="myform">
            Username:{" "}
            <input
              type="text"
              size="40"
              name="username"
              value={formData.username}
              onChange={handleChange}
            ></input>
          </div>

          <div className="myform">
            Password:{" "}
            <input
              type="password"
              size="40"
              name="password"
              value={formData.password}
              onChange={handleChange}
            ></input>
          </div>

          <div className="myform">
            Your OTP:{" "}
            <input
              type="text"
              size="40"
              name="yourotp"
              value={formData.yourotp}
              onChange={handleChange}
            ></input>
          </div>

          <div className="myform">
            <button>Login</button>
          </div>
        </div>
      </form>


    </>
  );
}

export default Login;
