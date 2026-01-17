#!/usr/bin/env bash
set -euo pipefail

# ----------------------------
# Configuration
# ----------------------------
DOCS_DIR="docs"
RULES_DIR="rules"
CURSOR_DIR="../../.cursor"
CURSOR_RULES_DIR="$CURSOR_DIR/rules"

START_TIME=$(date +"%Y-%m-%d %H:%M:%S")

echo "======================================="
echo "🚀 Starting rules extraction & sync"
echo "🕒 Start time : $START_TIME"
echo "📂 Docs dir  : $DOCS_DIR"
echo "📂 Rules dir : $RULES_DIR"
echo "📂 Cursor dir: $CURSOR_DIR"
echo "======================================="

# ----------------------------
# Pre-flight checks
# ----------------------------
if [ ! -d "$DOCS_DIR" ]; then
  echo "❌ Error: '$DOCS_DIR' does not exist"
  exit 1
fi

if [ ! -d "$CURSOR_DIR" ]; then
  echo "❌ Error: Cursor directory '$CURSOR_DIR' does not exist"
  exit 1
fi

# ----------------------------
# Counters
# ----------------------------
TOTAL_FILES=0
EXTRACTED=0
SKIPPED=0

# ----------------------------
# 🔥 Clean slate (local rules)
# ----------------------------
if [ -d "$RULES_DIR" ]; then
  echo "🧹 Removing existing local '$RULES_DIR'"
  rm -rf "$RULES_DIR"
fi

echo "📁 Creating local '$RULES_DIR'"
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
  for (j = 1; j < d; j++) p = p "  |"
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

echo "|"
echo "|-- local extraction complete"
echo "|"

# ----------------------------
# 3) SYNC TO CURSOR
# ----------------------------
echo "🔄 Syncing rules to Cursor"

if [ -d "$CURSOR_RULES_DIR" ]; then
  echo "🧹 Removing old Cursor rules: $CURSOR_RULES_DIR"
  rm -rf "$CURSOR_RULES_DIR"
fi

echo "📁 Copying new rules → $CURSOR_RULES_DIR"
cp -R "$RULES_DIR" "$CURSOR_RULES_DIR"

# ----------------------------
# SUMMARY
# ----------------------------
END_TIME=$(date +"%Y-%m-%d %H:%M:%S")

echo
echo "======================================="
echo "📊 Summary"
echo "======================================="
echo "🕒 Start time      : $START_TIME"
echo "🕒 End time        : $END_TIME"
echo "📄 Files scanned   : $TOTAL_FILES"
echo "✅ Rules extracted : $EXTRACTED"
echo "➖ Files skipped   : $SKIPPED"
echo "📁 Local rules     : $(find "$RULES_DIR" -type f | wc -l | tr -d ' ') files"
echo "📁 Cursor rules    : $(find "$CURSOR_RULES_DIR" -type f | wc -l | tr -d ' ') files"
echo "======================================="
echo "✨ Cursor rules successfully updated"
