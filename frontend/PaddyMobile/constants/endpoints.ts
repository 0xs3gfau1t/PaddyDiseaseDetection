import { BASE_API_URL } from '@env';
const API_ROOT = `http://192.168.1.73:3000/api`;

const endpoints = {
  dashboard: `${API_ROOT}/dashboard`,
  auth: {
    root: `${API_ROOT}/auth`,
    signup: `${API_ROOT}/auth/signup`,
    logout: `${API_ROOT}/auth/logout`,
    login: `${API_ROOT}/auth/login`,
  },
  profile: {
    root: `${API_ROOT}/profile`,
    changePassword: `${API_ROOT}/profile/change_password`,
  },
  map: `${API_ROOT}/heatmap`,
  uploadImage: `${API_ROOT}/upload`,
  diseases: `${API_ROOT}/uploads`,
  disease: `${API_ROOT}/upload`,
};

export default endpoints;
