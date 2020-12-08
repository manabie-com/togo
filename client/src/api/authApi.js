import axiosClient from './axiosClient';

const authApi = {
  login(data) {
    return axiosClient.post('/auth', data);
  },
};

export default authApi;