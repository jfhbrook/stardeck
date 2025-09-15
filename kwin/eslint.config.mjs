import pluginJs from '@eslint/js';

export default [
  { files: ['**/*.js'] },
  pluginJs.configs.recommended,
  {
    languageOptions: {
      globals: {
        callDBus: 'readable',
        workspace: 'readable',
        setInterval: 'readable',
      },
    },
  },
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
