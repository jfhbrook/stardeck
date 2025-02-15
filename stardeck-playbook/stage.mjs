import { createRequire } from 'node:module';

const require = createRequire(import.meta.url);

const toposort = require('toposort');

// TODO: Write tests for this
export function stages(graph) {
  const sorted = toposort(topoGraph(graph));

  const stages = [[]];

  let i = 0;

  // Fill out the first stage
  while (i < sorted.length && !dependencies(graph[sorted[i]]).length) {
    stages[0].push(graph[sorted[i]]);
    i++;
  }

  // If we're done, we're done
  if (i >= sorted.length) {
    return stages;
  }

  // Upstream dependencies
  let upstream = new Set(stages[0]);
  // The currently processing layer of dependencies
  let current = new Set();
  // The current stage
  let n = 1;

  while (i < sorted.length) {
    const name = sorted[i];
    const node = graph[name];

    // Mutate the node, it's fine
    node.name = name;

    // If there are any dependencies not in an upstream stage...
    if (!dependencies(node).every((dep) => upstream.has(dep))) {
      // Upstream deps now include the current stage
      upstream = upstream.union(current);

      // Set up a new stage
      current = new Set();
      stages.push([]);
      n++;
    }
    // Push the node to the current stage
    stages[n].push(node);
    // Mark the node as being in the current stage
    current.add(name);
    i++;
  }

  return stages;
}

export function topoGraph(graph) {
  const topo = [];
  for (let [name, node] in Object.entries(graph)) {
    for (let dependency of node.dependencies || []) {
      topo.push([name, dependency]);
    }
  }
  return topo;
}

function dependencies(node) {
  return node.dependencies || [];
}
