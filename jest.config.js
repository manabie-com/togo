/* eslint-disable @typescript-eslint/no-var-requires */
const { pathsToModuleNameMapper } = require("ts-jest");

const { compilerOptions } = require("./tsconfig.json");

const moduleNameMapper = pathsToModuleNameMapper(compilerOptions.paths);

module.exports = {
   coverageProvider: "v8",
   moduleNameMapper,
   modulePaths: ["<rootDir>"],
   roots: ["<rootDir>"],
   testEnvironment: "node",
   testMatch: ["**/*.(steps|test).(ts|js)"],
   transform: {
      "\\.ts$": ["ts-jest"],
   },
};
