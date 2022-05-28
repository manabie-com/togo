module.exports = {
  parser: '@typescript-eslint/parser',
  parserOptions: {
    project: 'tsconfig.json',
    sourceType: 'module',
  },
  plugins: ['@typescript-eslint/eslint-plugin', 'eslint-plugin-tsdoc', 'security'],
  extends: [
    'plugin:@typescript-eslint/eslint-recommended',
    'plugin:@typescript-eslint/recommended',
    'prettier',
    'prettier/@typescript-eslint',
    'plugin:prettier/recommended',
    'plugin:security/recommended',
  ],
  root: true,
  env: {
    node: true,
    jest: true,
  },
  rules: {
    // Disabled Rules
    '@typescript-eslint/explicit-function-return-type': 'off',
    '@typescript-eslint/explicit-member-accessibility': 'off',
    '@typescript-eslint/interface-name-prefix': 'off',
    '@typescript-eslint/no-empty-interface': 'off',
    '@typescript-eslint/no-inferrable-types': 'off',
    '@typescript-eslint/no-magic-numbers': 'off',
    '@typescript-eslint/no-namespace': 'off',
    '@typescript-eslint/no-non-null-assertion': 'off',
    '@typescript-eslint/no-type-alias': 'off',
    '@typescript-eslint/no-require-imports': 'off',
    '@typescript-eslint/no-object-literal-type-assertion': 'off',
    '@typescript-eslint/prefer-interface': 'off',
    '@typescript-eslint/no-var-requires': 'off',
    '@typescript-eslint/no-parameter-properties': 'off',
    '@typescript-eslint/no-explicit-any': 'off',
    '@typescript-eslint/camelcase': 'off',
    '@typescript-eslint/explicit-module-boundary-types': 'off',
    '@typescript-eslint/ban-types': 'off',
    // Enabled rules
    '@typescript-eslint/adjacent-overload-signatures': 'error',
    '@typescript-eslint/array-type': 'error',
    '@typescript-eslint/await-thenable': 'error',
    '@typescript-eslint/member-ordering': 'error',
    '@typescript-eslint/no-unnecessary-type-assertion': 'error',
    '@typescript-eslint/no-unused-vars': 'error',
    '@typescript-eslint/no-unnecessary-qualifier': 'error',
    '@typescript-eslint/no-misused-new': 'error',
    '@typescript-eslint/no-for-in-array': 'error',
    '@typescript-eslint/prefer-function-type': 'error',
    '@typescript-eslint/promise-function-async': 'error',
    '@typescript-eslint/restrict-plus-operands': 'error',
    '@typescript-eslint/type-annotation-spacing': 'error',
    '@typescript-eslint/unified-signatures': 'error',
    '@typescript-eslint/naming-convention': [
      'error',
      {
        'selector': 'enumMember',
        'format': ['StrictPascalCase']
      },
      {
        'selector': 'variable',
        'types': ['boolean'],
        'format': ['StrictPascalCase'],
        'prefix': ['is', 'should', 'has', 'can', 'did', 'will']
      },
      {
        'selector': 'variable',
        'format': ['camelCase', 'UPPER_CASE', 'StrictPascalCase']
      },
      {
        'selector': 'memberLike',
        'format': ['camelCase', 'snake_case'],
      }
    ],
    'complexity': ['error', { max: 15 }],
    'max-depth': ['error', { max: 4 }],
    'prefer-const': [
      'error',
      {
        destructuring: 'all',
        ignoreReadBeforeAssign: true,
      },
    ],
    'prettier/prettier': ['error', { 'singleQuote': true, 'semi': true, 'trailingComma': 'all', 'printWidth': 120 }],
    'tsdoc/syntax': 'warn',
    'padding-line-between-statements': [
      'error',
      // Add a new line after multiple variable declaration
      { blankLine: 'always', prev: ['const', 'let', 'var'], next: '*'},
      { blankLine: 'any',    prev: ['const', 'let', 'var'], next: ['const', 'let', 'var']},
      // Add a new line after multiple import declaration
      { blankLine: 'always', prev: 'import', next: '*'},
      { blankLine: 'any',    prev: 'import', next: 'import'},
      // Add a new line after if
      { blankLine: 'always', prev: 'if', next: '*'},
      // Add a new line before return
      { blankLine: 'always', prev: '*', next: 'return'},
      { blankLine: 'always', prev: '*', next: 'block' },
      { blankLine: 'always', prev: 'block', next: '*' },
      { blankLine: 'always', prev: '*', next: 'block-like' },
      { blankLine: 'always', prev: 'block-like', next: '*' },
    ],
    'lines-between-class-members': ['error', 'always'],
    "no-console": 2, // Alow console.error only
    "tsdoc/syntax": "off"
  },
};
