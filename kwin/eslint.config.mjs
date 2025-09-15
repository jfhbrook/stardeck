import pluginJs from '@eslint/js';

export default [
  { files: ['**/*.js'] },
  pluginJs.configs.recommended,
  {
    rules: {
      'no-unused-vars': [
        'error',
        {
          argsIgnorePattern: '^_',
          caughtErrorsIgnorePattern: '^_',
        },
      ],
    },
  },
];
