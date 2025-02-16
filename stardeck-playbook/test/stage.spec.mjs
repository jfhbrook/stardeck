import { expect, test } from 'vitest';

import { stages } from '../stage.mjs';

const GRAPH = {
  a: {},
  d: { dependencies: ['a', 'b'] },
  c: { dependencies: ['a'] },
  b: { dependencies: ['a'] },
};

// TODO: This fails because toposort doesn't guarantee that b comes before
// d - after all, all its dependencies are done before it
test('stage', () => {
  expect(stages(GRAPH)).toBe([
    [{ name: 'a' }],
    [
      { name: 'c', dependencies: ['a'] },
      { name: 'b', dependencies: ['a'] },
    ],
    [{ name: 'd', dependencies: ['a', 'b'] }],
  ]);
});
