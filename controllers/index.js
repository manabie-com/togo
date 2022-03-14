module.exports = (className, method) => {
  const Controller = require(`./${className}`);
  return (req, res, next) => {
    const controller = new Controller(req, res, next);
    return controller[method]();
  };
};
