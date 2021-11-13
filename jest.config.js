module.exports = {
  testMatch: ['**/__tests__/**/*.spec.js'],
  clearMocks: true,
  verbose: true,
  transform: {},
  moduleNameMapper: {
    // 'config/(.*)': '<rootDir>/src/config/$1',
    // 'utils/(.*)': '<rootDir>/src/utils/$1',
    // 'routes/(.*)': '<rootDir>/src/routes/$1',
    // 'modules/(.*)': '<rootDir>/src/modules/$1',
    // 'domain/(.*)': '<rootDir>/src/domain/$1',
    // 'infrastructure/(.*)': '<rootDir>/src/infrastructure/$1',
    // 'src/(.*)': '<rootDir>/src/$1',
    // 'middlewares/(.*)': '<rootDir>/src/middlewares/$1',
    'models/(.*)': '<rootDir>/src/infrastructure/models/$1',
  },
  moduleDirectories: [
    ".",
    "src",
    "node_modules",
  ]
};
