import { createRequire } from 'node:module';

const require = createRequire(import.meta.url);

const toposort = require('toposort');

export function stages(graph) {
  const [topo, noDeps] = topoGraph(graph);
  const sorted = toposort(topo);

  const stages = [
    noDeps.map((name) => {
      const node = graph[name];
      node.name = name;
      return node;
    }),
  ];

  if (!sorted.length) {
    if (!noDeps.length) {
      return [];
    }
    return stages;
  }

  let i = 0;

  // Fill out the first stage
  while (i < sorted.length && !dependencies(graph[sorted[i]]).length) {
    const name = sorted[i];
    if (!noDeps.includes(name)) {
      const node = graph[name];
      node.name = name;
      stages[0].push(node);
    }
    i++;
  }

  // If we're done, we're done
  if (i >= sorted.length) {
    return stages;
  }

  // Upstream dependencies
  let upstream = new Set(stages[0].map((node) => node.name));
  // The currently processing layer of dependencies
  let current = new Set();
  // The current stage
  let n = 0;

  if (stages[0].length) {
    stages.push([]);
    n++;
  }

  while (i < sorted.length) {
    const name = sorted[i];
    const node = graph[name];

    // Mutate the node, it's fine
    node.name = name;

    // If there are any dependencies not in an upstream stage...
    if (!dependencies(node).every((dep) => upstream.has(dep))) {
      // Upstream deps now include the current stage
      upstream = upstream.union(current);

      if (current.size) {
        // Set up a new stage
        current = new Set();
        stages.push([]);
        n++;
      }
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
  const noDeps = [];
  for (let [name, node] of Object.entries(graph)) {
    const deps = dependencies(node);
    if (!deps.length) {
      noDeps.push(name);
      continue;
    }

    for (let dependency of deps) {
      topo.push([dependency, name]);
    }
  }
  return [topo, noDeps];
}

function dependencies(node) {
  const deps = node.dependencies || [];

  if (!Array.isArray(deps)) {
    throw new Error(`Unexpected value for dependencies: ${node.dependencies}`);
  }

  return deps;
}
