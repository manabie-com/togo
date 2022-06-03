module.exports = {
  db: {
    url: process.env.DB_URL || "mongodb://localhost:27017/togo",
    urlTest: process.env.DB_TEST_URL || "mongodb://localhost:27017/togo-test",
  },
};