import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button, TextField } from "@mui/material";
import axios from "axios";
import { useDispatch } from "react-redux";
import { SET_USER, TOGGLE_LOGIN, TOGGLE_REG } from "../../reducers/auth";

const Login = () => {
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    const config = {
      header: {
        "Content-Type": "application/json",
      },
    };

    try {
      const { data } = await axios.post("/api/v1/auth/login", { username, password }, config);
      if (data.data) {
        dispatch(SET_USER(data.data));
        dispatch(TOGGLE_LOGIN());
        navigate("/");
      }
    } catch (error) {
      console.log(error.response.data.error);
      setError(error.response.data.error);
      setTimeout(() => {
        setError("");
      }, 4000);
    }
  };

  useEffect(() => {
    let modal = document.querySelector(".auth-con");
    window.onclick = function (event) {
      if (event.target === modal) {
        dispatch(TOGGLE_LOGIN());
      }
    };
  }, [dispatch]);

  return (
    <div className="auth-con cen-grid" style={{ zIndex: "9999" }}>
      <div className="auth-modal login">
        <div className="logo">
          <h1 style={{ letterSpacing: "-1px" }} onClick={() => navigate("/")}>
            FLENJO
          </h1>
        </div>
        <div className="title">
          <h1>Login</h1>
        </div>
        <form className="auth-form" autoComplete="off" onSubmit={handleSubmit}>
          <div className="input-con">
            <TextField
              variant="outlined"
              size="small"
              name="username"
              label="Username"
              fullWidth
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
            />
          </div>
          <div className="input-con">
            <TextField
              variant="outlined"
              size="small"
              name="password"
              label="Password"
              type="password"
              fullWidth
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>

          {error && <span className="error-msg">{error}</span>}

          <div className="cen-grid">
            <Button variant="contained" color="secondary" type="submit">
              Login
            </Button>
          </div>
        </form>
        <p>
          New to Flenjo?{" "}
          <span style={{ color: "#ed5a6b", cursor: "pointer" }} onClick={() => dispatch(TOGGLE_REG())}>
            Create account
          </span>
        </p>
      </div>
    </div>
  );
};

export default Login;
