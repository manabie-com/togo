module.exports = {
  authService: {
    jwt: 'jwt',
  },
  jwt: {
    accessTokenExpiration: 31536000000, // 1 year
    accessType: 'access',
    accessSecretKey: 'ghaHGTG$%#!#@',
  },
  saltRoundPassword: 8,
  port: 5050,
  testUser: {
    username: 'test_user',
    password: 'test_pass',
  },
}