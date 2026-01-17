#!/usr/bin/env bash

DOCS_DIR="docs"

if [ ! -d "$DOCS_DIR" ]; then
  echo "Error: folder '$DOCS_DIR' does not exist"
  exit 1
fi

find "$DOCS_DIR" -name "*.md" -print \
| sed "s|^$DOCS_DIR/||" \
| sort \
| awk '
function prefix(d) {
  p = "|"
  for (j = 1; j < d; j++) {
    p = p "  |"
  }
  return p
}

{
  depth = split($0, parts, "/")
  path = ""

  for (i = 1; i < depth; i++) {
    path = path parts[i] "/"
    if (!seen[path]) {
      print prefix(i) "-- " parts[i]
      seen[path] = 1
    }
  }

  print prefix(depth) "-- " parts[depth]
}'
