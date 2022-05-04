import healthCheck from './healthCheck/healthCheck.controller';
import user from './user/user.controller';

const routes = [...healthCheck, ...user];

export default routes;
