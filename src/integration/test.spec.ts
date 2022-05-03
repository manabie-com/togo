import axios from 'axios';

const createFakeUserName = (length) => {
  let result = '';
  const characters =
    'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
  const charactersLength = characters.length;
  for (let i = 0; i < length; i++) {
    result += characters.charAt(Math.floor(Math.random() * charactersLength));
  }
  return result;
};

describe('integration testing', () => {
  const axiosInstance = axios.create({ baseURL: 'http://localhost:3000' });
  it('Should create user and create task successfully', async (done) => {
    // create new user
    const createUserResponse = await axiosInstance.post('/user', {
      username: createFakeUserName(8),
      password: 'password',
      configuration: {
        limit: 1,
        type: 'DAILY'
      }
    });

    const {
      data: {
        data: { userId }
      }
    } = createUserResponse;

    //create new task
    const createTaskResponse = await axiosInstance.post(
      `/user/${userId}/task`,
      {
        name: '_task'
      }
    );

    setTimeout(async () => {
      //get list task
      const getTasksResponse = await axiosInstance.get(`/user/${userId}/tasks`);

      const {
        data: { data: tasks }
      } = getTasksResponse;

      expect(createUserResponse.status).toEqual(201);
      expect(createTaskResponse.status).toEqual(201);
      expect(getTasksResponse.status).toEqual(200);
      expect(tasks).toEqual(
        expect.arrayContaining([
          expect.objectContaining({ userId, status: 'DONE' })
        ])
      );
      done();
    }, 2000);
  });

  it('Should create task with the status is FAILED when the limit configuration < count', async (done) => {
    // create new user
    const createUserResponse = await axiosInstance.post('/user', {
      username: createFakeUserName(8),
      password: 'password',
      configuration: {
        limit: 1,
        type: 'DAILY'
      }
    });

    const {
      data: {
        data: { userId }
      }
    } = createUserResponse;

    //create new task
    const createFirstTaskResponse = axiosInstance.post(`/user/${userId}/task`, {
      name: '_firstTask'
    });
    const createSecondTaskResponse = axiosInstance.post(
      `/user/${userId}/task`,
      {
        name: '_secondTask'
      }
    );

    await Promise.all([createFirstTaskResponse, createSecondTaskResponse]);

    setTimeout(async () => {
      //get list task
      const getTasksResponse = await axiosInstance.get(
        `/user/${userId}/tasks?status=ALL`
      );

      const {
        data: { data: tasks }
      } = getTasksResponse;

      expect(createUserResponse.status).toEqual(201);
      expect(getTasksResponse.status).toEqual(200);
      expect(tasks).toEqual(
        expect.arrayContaining([
          expect.objectContaining({ userId, status: 'DONE' }),
          expect.objectContaining({ userId, status: 'FAILED' })
        ])
      );
      done();
    }, 2000);
  });
});
