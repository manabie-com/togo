const Logger = require("./libs/logger");

const bannerLogger = require("./libs/banner");

const winstonLoader = require("./loaders/winstonLoader");
const mongooseLoader = require("./loaders/mongooseLoader");
const expressLoader = require("./loaders/expressLoader");
const swaggerLoader = require("./loaders/swaggerLoader");

const log = new Logger(__filename);

async function initApp() {
  // logging
  winstonLoader();

  // Database
  await mongooseLoader();

  // express
  const app = expressLoader();

  // swagger
  swaggerLoader(app);
}

initApp()
  .then(() => bannerLogger(log))
  .catch((error) => log.error("Application is crashed: " + error));
