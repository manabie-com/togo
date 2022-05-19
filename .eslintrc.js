module.exports = {
   env: {
      jest: true,
      node: true,
   },
   extends: [
      "eslint:recommended",
      "plugin:@typescript-eslint/recommended",
      "airbnb-base",
      "prettier",
      "plugin:prettier/recommended",
      "plugin:import/recommended",
      "plugin:import/typescript",
   ],
   parser: "@typescript-eslint/parser",
   plugins: ["@typescript-eslint"],
   root: true,
   rules: {
      "prettier/prettier": ["error"],
      "no-promise-executor-return": ["warn"],
      "no-unused-vars": 0,
      "no-shadow": 0,
      "@typescript-eslint/no-unused-vars": ["warn"],
      "import/extensions": [
         "error",
         "ignorePackages",
         {
            ts: "never",
         },
      ],
      "import/no-extraneous-dependencies": [0],
      "import/no-named-as-default-member": 0,
      "import/no-unresolved": [2, { amd: true, commonjs: true }],
      "import/order": [
         "error",
         {
            alphabetize: { caseInsensitive: true, order: "asc" },
            groups: ["builtin", "external", "internal", ["index", "sibling", "parent", "object"]],
            "newlines-between": "always",
            // define material-ui group that will appear separately after other main externals
            pathGroups: [{ group: "external", pattern: "@/**", position: "after" }],
            // default is builtin, external; but we want to divide up externals into groups also
            pathGroupsExcludedImportTypes: ["builtin"],
         },
      ],
      "import/prefer-default-export": 0,
      "no-console": [
         1,
         {
            allow: ["error", "info"],
         },
      ],
      "prettier/prettier": ["error"],
      quotes: [
         2,
         "double",
         {
            allowTemplateLiterals: true,
            avoidEscape: true,
         },
      ],
      "sort-keys": [1, "asc", { caseSensitive: true, natural: true }],
   },
   settings: {
      "import/extensions": [".js", ".jsx", ".ts", ".tsx"],
      "import/parsers": {
         "@typescript-eslint/parser": [".ts", ".tsx"],
      },
      "import/resolver": {
         node: {
            extensions: [".js", ".jsx", ".ts", ".tsx"],
         },
         typescript: {},
      },
   },
};
