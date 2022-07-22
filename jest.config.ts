export default {
  clearMocks: true,
  collectCoverage: true,
  coverageDirectory: "coverage",
  coverageProvider: "v8",
  preset: 'ts-jest',
  testEnvironment: "node",
  testPathIgnorePatterns: [
    "/node_modules/",
    "/bin/",
  ],
  transform: {
	  '^.+\\.ts?$': 'ts-jest',
  },
  transformIgnorePatterns: [
    "/node_modules/",
    "\\.pnp\\.[^\\/]+$",
  ],
};
