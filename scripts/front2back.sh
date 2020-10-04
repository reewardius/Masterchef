#!/bin/sh

# Get root path
root=$(dirname "${0}")/..
# Compile frontend
npm run --prefix "${root}/frontend" compile 1>/dev/null || exit 1
# Copy to backend
cat > "${root}/pkg/index.go" << EOF 
package pkg

// ====================
//  GLOBALS
// ====================

const source = \`$(cat "${root}/frontend/dist/index.html")\`
EOF
# Output
echo "[+] Frontend generated ($(du -h "${root}/frontend/dist/index.html" | awk '{ print $1 }'))"