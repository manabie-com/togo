import axiosClient from './axiosClient';

const taskApi = {
  getAll(token) {
    return axiosClient.get('/tasks', {
      headers: {
        Auhoriration: token,
      }
    });
  },

  add(data) {
    return axiosClient.post('/tasks', data);
  },
};

export default taskApi;