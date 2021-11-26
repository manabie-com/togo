const login = (username, password) => {
  const trimUsername = username?.trim();
  const trimPassword = password?.trim();
  const lengthValidation = trimUsername?.length && trimPassword?.length;
  const verifyAccount =
    trimUsername === "firstUser" && trimPassword === "example";

  if (!lengthValidation || !verifyAccount) {
    return false;
  }

  return true;
};

module.exports = {
  login,
};
