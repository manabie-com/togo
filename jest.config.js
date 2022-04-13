module.exports = {
  testPathIgnorePatterns: ['test'],
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
};
