module.exports = (className, methods) => {
  const Middleware = require(`./${className}`);

  if (Array.isArray(methods)) {
    const middlewares = [];
    methods.forEach((method) => {
      middlewares.push((req, res, next) => {
        const middleware = new Middleware(req, res, next);
        return middleware[method]();
      });
    });

    return middlewares;
  } else {
    return (req, res, next) => {
      const middleware = new Middleware(req, res, next);
      return middleware[methods]();
    };
  }
};
