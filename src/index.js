const config = require('./config/constants');
const app = require('./server');

const server = app.listen(config.port, () => {
  console.log(`Server listening on port ${config.port}`);
});

const unexpectedErrorHandler = () => {
  if (server) {
    server.close(() => {
      process.exit(1);
    });
  } else {
    process.exit(1);
  }
};

// Uncaught Exception
process.on('uncaughtException', (exception) => {
  console.log(`UNCAUGHT EXCEPTION: ${exception.stack}`);

  unexpectedErrorHandler();
});

// Unhandled Rejection
process.on('unhandledRejection', (reason) => {
  console.log(`UNHANDLED REJECTION: ${reason}`);

  unexpectedErrorHandler();
});
