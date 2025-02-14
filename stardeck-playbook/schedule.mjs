export function schedule(graph) {
  return _schedule(_nodes(graph), _edges(graph));
}

export function _nodes(graph) {
  const nodes = new Set();
  for (let [name, spec] of Object.entries(graph)) {
    nodes.add(name);
    for (let dependency of spec.dependencies) {
      nodes.add(dependency);
    }
  }
  return Array.from(nodes);
}

export function _edges(graph) {
  const edges = [];

  for (let [name, spec] of Object.entries(graph)) {
    for (let dependency of spec.dependencies) {
      edges.push([name, dependency]);
    }
  }

  return edges;
}

export function _schedule(nodes, edges) {
  // TODO
  // Should be similar to a standard topo sort, but with "layers" that allow
  // for running code in parallel.
  let cursor = nodes.length;
}
