import { expect, test } from 'vitest';

import { stages } from '../stage.mjs';

const GRAPH = {
  a: {},
  d: { dependencies: ['a', 'b'] },
  c: { dependencies: ['a'] },
  b: { dependencies: ['a'] },
};

test('stage', () => {
  expect(stages(GRAPH)).toEqual([
    [{ name: 'a' }],
    [{ name: 'b', dependencies: ['a'] }],
    [
      { name: 'd', dependencies: ['a', 'b'] },
      // NOTE: c could be sheduled alongside b. However, because it's listed
      // later in the graph, it shows up here becuase the toposort keeps it
      // *after* d.
      { name: 'c', dependencies: ['a'] },
    ],
  ]);
});
