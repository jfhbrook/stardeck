import swc from 'unplugin-swc';
import { defineConfig } from 'vite';

export default defineConfig({
  test: {
    exclude: ['node_modules'],
  },
  plugins: [
    swc.vite({
      jsc: {
        parser: {
          syntax: 'typescript',
          jsx: false,
          decorators: true,
          decoratorsBeforeExport: true,
          dynamicImport: true,
          privateMethod: true,
          functionBind: true,
          exportDefaultFrom: true,
          exportNamespaceFrom: true,
          topLevelAwait: false,
          importMeta: false,
        },
        transform: { legacyDecorator: false, decoratorMetadata: true },
        target: 'es2022',
        loose: false,
        externalHelpers: false,
        keepClassNames: true,
      },
      module: { type: 'nodenext' },
      sourceMaps: true,
      inlineSourcesContent: true,
      minify: false,
    }),
  ],
});
