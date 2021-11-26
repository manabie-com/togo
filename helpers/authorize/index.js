const MOCK_TOKEN = require("../constants");

const validAuthorized = (token) => {
  const trimToken = token?.trim();
  const validToken = trimToken?.length && token === MOCK_TOKEN;
  if (!validToken) {
    return false;
  }
  return true;
};

module.exports = validAuthorized;
