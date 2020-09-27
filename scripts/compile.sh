#!/bin/bash

# Error function
error() {
    echo "[-] ${1}" ; exit 1
}

# SIGTERM

_term() { 
  echo "Caught SIGTERM signal!" 
  kill -TERM "${$}" 2>/dev/null
}

trap _term SIGTERM

# Get root path
root=$(dirname "${0}")/..
# Compile frontend
sh "${root}"/scripts/front2back.sh 2>/dev/null || error "Can't compile frontend"
# Compile the application
rm -rf "${root}"/dist
compile() {
    test "${1}/${2}" = "darwin/386" && return
    e=
    test "${1}" = "windows" && e=.exe
    n="${root}/dist/masterchef_${1}-${2}${e}"
    GOOS="${1}" GOARCH="${2}" go build -ldflags "-s -w" -i -o "${n}" "${root}"/main.go
    echo "[+] ${n} compiled ($(du -h "${n}" | awk '{ print $1 }'))"
}
arch="386 amd64"
os="darwin linux windows"
for o in ${os}; do
    for a in ${arch}; do
        compile "${o}" "${a}" &
    done
done
wait
# Compress the binary
which upx >/dev/null || error "upx needed to compress"
compress() {
    test -z "${1##*.exe}" && win=.exe
    u="${1//.exe/}"_upx
    cp "${1}" "${u}"
    upx -9 "${2}" "${u}" > /dev/null
    test "${u}" = "${u}${win}" || mv "${u}" "${u}${win}"
    echo "[+] ${1} compressed ($(du -h "${1}" | awk '{ print $1 }') -> $(du -h "${u}""${win}" | awk '{ print $1 }'))"
}
binaries=$(find "${root}/dist" -type f)
for b in ${binaries}; do
    if test "${1}" = "-9" ; then t="--ultra-brute" ; else t= ; fi
    compress "${b}" "${t}" &
done
wait