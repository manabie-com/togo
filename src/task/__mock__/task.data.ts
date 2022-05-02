export const createTaskPayload = {
  userId: 'username',
  name: '_task'
};

export const createTaskPayloadFn = (payload: any) => ({
  ...payload,
  name: '_task'
});
