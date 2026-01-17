#!/usr/bin/env bash

DOCS_DIR="docs"
RULES_DIR="rules"

if [ ! -d "$DOCS_DIR" ]; then
  echo "Error: folder '$DOCS_DIR' does not exist"
  exit 1
fi

mkdir -p "$RULES_DIR"

# ----------------------------
# 1) PRINT TREE (OLD SCRIPT)
# ----------------------------
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
}
'

echo "|"
echo "|-- extracting rules"
echo "|"

# ----------------------------
# 2) EXTRACT PERSISTENT CONTEXT
# ----------------------------
find "$DOCS_DIR" -name "*.md" -print \
| sed "s|^$DOCS_DIR/||" \
| sort \
| while IFS= read -r rel; do
  src="$DOCS_DIR/$rel"
  dest="$RULES_DIR/$rel"

  mkdir -p "$(dirname "$dest")"

  # Extract first ```md ... ``` block
  awk '
    BEGIN { in_block = 0 }
    /^```md[[:space:]]*$/ { in_block = 1; next }
    /^```[[:space:]]*$/ && in_block { exit }
    in_block { print }
  ' "$src" > "$dest.tmp"

  if [ -s "$dest.tmp" ]; then
    mv "$dest.tmp" "$dest"
    echo "|-- ✓ $rel"
  else
    rm "$dest.tmp"
    echo "|-- - $rel (no Persistent Context)"
  fi
done

echo "|"
echo "|-- done"
