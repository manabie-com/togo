const httpStatus = require("http-status");

const catchAsync = require("../../utils/catch-async");
const { userService, authService } = require("../services");

const register = catchAsync(async (req, res) => {
  const user = await userService.createUser(req.body);

  const token = user.generateAuthToken();

  res.status(httpStatus.CREATED).send({ user, token });
});

const login = catchAsync(async (req, res) => {
  const { email, password } = req.body;
  const user = await authService.login(email, password);

  const token = await user.generateAuthToken();

  res.send({ user, token });
});

module.exports = {
  register,
  login,
};
