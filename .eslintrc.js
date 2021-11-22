module.exports = {
  env: {
    browser: true,
    commonjs: true,
    es2021: true,
  },
  extends: ["airbnb-base", "prettier"],
  parserOptions: {
    ecmaVersion: 13,
  },
  rules: {},
  overrides: [
    {
      files: ["bin/*.js", "lib/*.js"],
      excludedFiles: "*.test.js",
      rules: {
        quotes: ["error", "single"],
      },
    },
  ],
};
