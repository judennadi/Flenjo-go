import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { Button, TextField } from "@mui/material";
import axios from "axios";
import { useDispatch } from "react-redux";
import { SET_USER, TOGGLE_LOGIN, TOGGLE_REG } from "../../reducers/auth";

const Register = () => {
  const navigate = useNavigate();
  const dispatch = useDispatch();
  // const { dispatch } = useContext(AuthContext);
  const [name, setName] = useState("");
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [error, setError] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (password !== confirmPassword) {
      setError("Passwords do not match");
      setTimeout(() => {
        setError("");
      }, 4000);
      setPassword("");
      setConfirmPassword("");
      return;
    }

    const config = {
      header: {
        "Content-Type": "application/json",
      },
    };

    try {
      const { data } = await axios.post("/api/v1/auth/register", { name, username, email, password }, config);
      if (data.data) {
        dispatch(SET_USER(data.data));
        dispatch(TOGGLE_REG());
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
        dispatch(TOGGLE_REG());
      }
    };
  }, [dispatch]);

  return (
    <div className="auth-con cen-grid" style={{ zIndex: "9999" }}>
      <div className="auth-modal">
        <div className="logo">
          <h1 style={{ letterSpacing: "-1px" }} onClick={() => navigate("/")}>
            FLENJO
          </h1>
        </div>
        <div className="title">
          <h1>Register</h1>
        </div>
        <form className="auth-form" autoComplete="off" onSubmit={handleSubmit}>
          <div className="input-con">
            <TextField
              variant="outlined"
              size="small"
              name="name"
              label="Full Name"
              fullWidth
              value={name}
              onChange={(e) => setName(e.target.value)}
              required
            />
          </div>
          <div className="input-con">
            <TextField
              variant="outlined"
              size="small"
              name="email"
              label="Email"
              type="email"
              fullWidth
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
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
          <div className="input-con">
            <TextField
              variant="outlined"
              size="small"
              name="confirm_Password"
              label="Confirm Password"
              type="password"
              fullWidth
              value={confirmPassword}
              onChange={(e) => setConfirmPassword(e.target.value)}
              required
            />
          </div>

          {error && <span className="error-msg">{error}</span>}

          <div className="cen-grid">
            <Button variant="contained" color="secondary" type="submit">
              Register
            </Button>
          </div>
        </form>
        <p>
          Already a member?{" "}
          <span style={{ color: "#ed5a6b", cursor: "pointer" }} onClick={() => dispatch(TOGGLE_LOGIN())}>
            Login
          </span>
        </p>
      </div>
    </div>
  );
};

export default Register;
