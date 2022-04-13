import startApp from './app';

module.exports = async () => {
  const app = await startApp();
  (global as any).__APP__ = app;
};
