module.exports = {
  testEnvironment: 'node',
  testEnvironmentOptions: {
    NODE_ENV: 'test',
  },
  restoreMocks: true,
  coveragePathIgnorePatterns: ['node_modules', 'backend/config', 'backend/app/app.js', 'tests'],
  coverageReporters: ['text', 'lcov', 'clover', 'html'],
};
