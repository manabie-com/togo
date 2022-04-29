module.exports = {
  roots: ['<rootDir>/src'],
  setupFilesAfterEnv: ['<rootDir>/src/__jest__/setup.ts'],
  transform: {
    '^.+\\.ts$': 'ts-jest'
  },
  preset: 'ts-jest',
  testEnvironment: 'node',
  testMatch: ['**/?(*.)(spec|test).ts'],
  moduleFileExtensions: ['ts', 'js', 'json', 'node'],
  globals: {
    'ts-jest': {
      diagnostics: false
    }
  },
  coverageThreshold: {
    global: {
      branches: 80,
      functions: 80,
      lines: 80
    }
  },
  collectCoverageFrom: [
    'src/**/*.{ts,js}',
    '!src/**/*.(interface|constant|model|enum|type).{ts,js}',
    '!**/__mocks__/**',
    '!**/node_modules/**'
  ]
};
