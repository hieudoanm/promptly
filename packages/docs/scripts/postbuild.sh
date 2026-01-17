#!/usr/bin/env bash
set -euo pipefail

DOCS_DIR="docs"
RULES_DIR="rules"

START_TIME=$(date +"%Y-%m-%d %H:%M:%S")

echo "======================================="
echo "🚀 Starting rules extraction"
echo "🕒 Start time: $START_TIME"
echo "📂 Docs dir : $DOCS_DIR"
echo "📂 Rules dir: $RULES_DIR"
echo "======================================="

if [ ! -d "$DOCS_DIR" ]; then
  echo "❌ Error: folder '$DOCS_DIR' does not exist"
  exit 1
fi

# ----------------------------
# Counters
# ----------------------------
TOTAL_FILES=0
EXTRACTED=0
SKIPPED=0

# ----------------------------
# 🔥 Clean slate
# ----------------------------
if [ -d "$RULES_DIR" ]; then
  echo "🧹 Removing existing '$RULES_DIR' directory"
  rm -rf "$RULES_DIR"
fi

echo "📁 Creating '$RULES_DIR'"
mkdir -p "$RULES_DIR"

# ----------------------------
# 1) PRINT TREE
# ----------------------------
echo
echo "📜 Markdown tree:"
echo

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
  TOTAL_FILES=$((TOTAL_FILES + 1))

  src="$DOCS_DIR/$rel"
  dest="$RULES_DIR/$rel"

  mkdir -p "$(dirname "$dest")"

  awk '
    BEGIN { in_block = 0 }
    /^```md[[:space:]]*$/ { in_block = 1; next }
    /^```[[:space:]]*$/ && in_block { exit }
    in_block { print }
  ' "$src" > "$dest.tmp"

  if [ -s "$dest.tmp" ]; then
    mv "$dest.tmp" "$dest"
    EXTRACTED=$((EXTRACTED + 1))
    echo "|-- ✓ extracted: $rel"
  else
    rm "$dest.tmp"
    SKIPPED=$((SKIPPED + 1))
    echo "|-- - skipped  : $rel (no Persistent Context)"
  fi
done

END_TIME=$(date +"%Y-%m-%d %H:%M:%S")

echo "|"
echo "|-- done"
echo "|"

# ----------------------------
# SUMMARY
# ----------------------------
echo "======================================="
echo "📊 Extraction summary"
echo "======================================="
echo "🕒 Start time     : $START_TIME"
echo "🕒 End time       : $END_TIME"
echo "📄 Files scanned  : $TOTAL_FILES"
echo "✅ Rules extracted: $EXTRACTED"
echo "➖ Files skipped  : $SKIPPED"
echo "📁 Files written  : $(find "$RULES_DIR" -type f 2>/dev/null | wc -l | tr -d ' ')"
echo "======================================="
echo "✨ Done"
