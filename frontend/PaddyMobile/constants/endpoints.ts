const API_ROOT = 'http://192.168.10.104:3000/api';

const endpoints = {
  auth: {
    root: `${API_ROOT}/auth`,
    signup: `${API_ROOT}/auth/signup`,
    logout: `${API_ROOT}/auth/logout`,
    login: `${API_ROOT}/auth/login`,
  },
  profile: {
    root: `${API_ROOT}/profile`,
  },
  uploadImage: `${API_ROOT}/upload`,
  disease: `${API_ROOT}/disease`,
};

export default endpoints;
