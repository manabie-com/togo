module.exports = {
  testEnvironment: 'node',
  globals: {
    __APP__: true,
  },
  globalSetup: '<rootDir>/test/globalSetup.ts',
  globalTeardown: '<rootDir>/test/globalTeardown.ts',
  testTimeout: 30000,
  maxWorkers: 1,
  moduleFileExtensions: ['js', 'json', 'ts'],
  moduleDirectories: ['node_modules', 'src'],
  modulePaths: ['src', 'node_modules'],
  preset: 'ts-jest',
  testRegex: '.spec.ts$',
  transform: {
    '~.+\\.(t|j)s$': 'ts-jest',
  },
  moduleNameMapper: {
    '^src/(.*)': '<rootDir>/src/$1',
  },
  coverageDirectory: '../coverage',
  testEnvironment: 'node',
};
