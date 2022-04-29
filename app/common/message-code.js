const responseMessage = code => {
  return { message: code }
}

module.exports = {
  I001: 'Logout successful',
  E001: 'Username or password incorrect',
  E002: 'User\'s tasks are limited',
  E003: 'Token is invalid',
  responseMessage
}
