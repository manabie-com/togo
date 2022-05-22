function getOsEnv(key) {
  if (typeof process.env[key] === "undefined") {
    throw new Error(`Environment variable ${key} is not set.`);
  }

  return process.env[key];
}

function getOsEnvOptional(key) {
  return process.env[key];
}

function toBool(value) {
  return value === "true";
}

function normalizePort(port) {
  const parsedPort = parseInt(port, 10);
  if (isNaN(parsedPort)) {
    // named pipe
    return port;
  }
  if (parsedPort >= 0) {
    // port number
    return parsedPort;
  }
  return false;
}

module.exports = {
  getOsEnv,
  getOsEnvOptional,
  toBool,
  normalizePort,
};
