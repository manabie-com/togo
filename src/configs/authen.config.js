module.exports = {
  jwt: {
    secretKey : "This1s4Rand0m",
    expiresInMinutes: process.env.JWT_TIMEOUT || (3 * 30 * 24 * 60), // 8 hours in minutes
    algorithm: "HS256"
  },
};