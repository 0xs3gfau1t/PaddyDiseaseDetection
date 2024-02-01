const API_ROOT = "http://127.0.0.1:3000/api";

const endpoints = {
  auth: {
    root: `${API_ROOT}/auth`,
    signup: `${API_ROOT}/auth/signup`,
    logout: `${API_ROOT}/auth/logout`,
    login: `${API_ROOT}/auth/login`,
  },
};

export default endpoints;
