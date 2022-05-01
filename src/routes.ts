import healthCheck from './healthCheck/healthCheck.controller';
import user from './user/user.controller';
import task from './task/task.controller';

const routes = [...healthCheck, ...user, ...task];

export default routes;
