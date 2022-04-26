const httpStatus = require('http-status');
const ApiError = require('../utils/ApiError');
const catchAsync = require('../utils/catchAsync');
const { userService } = require('../services');
const response = require('../utils/responseTemp');

/**
 * Get own user
 */
const getMyProfile = catchAsync(async (req, res) => {
  const user = await userService.getUserByPk(req.user.id);
  if (!user) {
    throw new ApiError(httpStatus.NOT_FOUND, 'User not found');
  }
  res.send(response(httpStatus.OK, 'Get profile success', user));
});

/**
 * Update user
 */
const updateUser = catchAsync(async (req, res) => {
  const user = await userService.updateUserByPk(req.user.id, req.body);
  res.send(response(httpStatus.OK, 'Update user success', user));
});

/**
 * Delete an user
 */
const deleteUser = catchAsync(async (req, res) => {
  await userService.deleteUserByPk(req.params.userId);
  res.send(response(httpStatus.OK, 'Delete user success'));
});

module.exports = {
  updateUser,
  deleteUser,
  getMyProfile,
};
