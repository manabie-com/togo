import hapi from '@hapi/hapi';

const ping: hapi.ServerRoute = {
  method: 'GET',
  path: '/ping',
  options: {
    description: 'Pongs back',
    notes: 'To check is service pongs on a ping',
    tags: ['api'],
    handler: () => 'pong !!!'
  }
};

const healthCheckController: hapi.ServerRoute[] = [ping];

export default healthCheckController;
