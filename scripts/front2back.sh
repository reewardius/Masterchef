#!/bin/sh

# Get root path
root=$(dirname "${0}")/..
# Compile frontend
if ! test -f "${root}/frontend/dist/index.html"
then
    npm run --prefix "${root}/frontend" compile 1>/dev/null || exit 1
fi
# Copy to backend
cat > "${root}/pkg/public/index.go" << EOF 
package public

// ====================
//  GLOBALS
// ====================

// Source code (Frontend: HTML + CSS + JS) in Go Template format to allow the use of variables
// Current variables:
// - .Addr :: Web socket server address
const Source = \`$(cat "${root}/frontend/dist/index.html")\`
EOF
# Output
echo "[+] Frontend generated ($(du -h "${root}/frontend/dist/index.html" | awk '{ print $1 }'))"