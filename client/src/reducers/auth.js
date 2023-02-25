import { createSlice } from "@reduxjs/toolkit";

if (!document.cookie.includes("check=")) {
  localStorage.removeItem("state");
}

const prevState = JSON.parse(localStorage.getItem("state"));
const initialState = prevState
  ? { user: prevState, isAuth: true, isLogin: false, isReg: false }
  : { user: null, isAuth: false, isLogin: false, isReg: false };

const authSlice = createSlice({
  name: "auth",
  initialState,
  reducers: {
    SET_USER: (state, action) => {
      state.user = action.payload;
      state.isAuth = true;
      localStorage.setItem("state", JSON.stringify(state.user));
    },
    REMOVE_USER: (state) => {
      state.user = null;
      state.isAuth = false;
      localStorage.removeItem("state");
    },
    TOGGLE_LOGIN: (state) => {
      state.isReg = false;
      state.isLogin = !state.isLogin;
    },
    TOGGLE_REG: (state) => {
      state.isLogin = false;
      state.isReg = !state.isReg;
    },
  },
});

export const { SET_USER, REMOVE_USER, TOGGLE_LOGIN, TOGGLE_REG } =
  authSlice.actions;
export default authSlice.reducer;
