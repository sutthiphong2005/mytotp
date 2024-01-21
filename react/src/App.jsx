import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import "./App.css";

function App() {
  const [count, setCount] = useState(0);

  const [formData, setFormData] = useState({
    username: "",
    password: "",
    repassword: "",
  });

  const [qrcode, setQRcode] = useState("");

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
        repassword: formData.repassword,
      }),
    };

    fetch("http://localhost:8080/register", requestOptions)
      .then((response) => response.json())
      .then((result) => {
        setQRcode(result.png);
        alert(result);
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

      <h1>Register</h1>
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
            Re-Password:{" "}
            <input
              type="password"
              size="40"
              name="repassword"
              value={formData.repassword}
              onChange={handleChange}
            ></input>
          </div>

          <div className="myform">
            <button onClick={() => setCount((count) => count + 1)}>
              Sign Up
            </button>
          </div>
        </div>
      </form>

      <div>
        <img src={qrcode} alt="qrcode" />
      </div>
    </>
  );
}

export default App;
