// import { BASE_API_URL } from '@env';
// const API_ROOT = `http://192.168.30.44:3000/api`;

const endpoints = {
  dashboard: `/dashboard`,
  auth: {
    root: `/auth`,
    signup: `/auth/signup`,
    logout: `/auth/logout`,
    login: `/auth/login`,
  },
  profile: {
    root: `/profile`,
    changePassword: `/profile/change_password`,
  },
  map: `/heatmap`,
  uploadImage: `/upload`,
  diseases: `/uploads`,
  disease: `/upload`,
  diseaseStat: `/uploadStat`,
};

export default endpoints;
