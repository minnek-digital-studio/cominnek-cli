#!/bin/bash

VERSION=${1}
DIST_DIR="./dist"
FILE_ARM86="${DIST_DIR}/cominnek-${VERSION}-arm86.dmg"
FILE_ARM64="${DIST_DIR}/cominnek-${VERSION}-arm64.dmg"
FILE_AMD86="${DIST_DIR}/cominnek-${VERSION}-amd86.dmg"
FILE_AMD64="${DIST_DIR}/cominnek-${VERSION}-amd64.dmg"

if [ -z "${VERSION}" ]; then
    echo "Usage: build.sh <version>"
    exit 1
fi

mkdir -p ${DIST_DIR}

# Build for ARM (32 bits)
GOARCH=arm GOOS=darwin GO386=softfloat go build -o ./build/bin-arm86
cp ./assets/darwin/bin/* ./build
cp ./build/bin-arm86 ./build/bin
rm -f ${FILE_ARM86}
hdiutil create -fs HFS+ -srcfolder "./build" -volname "cominnek-${VERSION}-arm86" "${FILE_ARM86}"

# Clean up for next build
rm -rf ./build/*

# Build for ARM (64 bits)
GOARCH=arm64 GOOS=darwin go build -o ./build/bin-arm64
cp ./assets/darwin/bin/* ./build
cp ./build/bin-arm64 ./build/bin
rm -f ${FILE_ARM64}
hdiutil create -fs HFS+ -srcfolder "./build" -volname "cominnek-${VERSION}-arm64" "${FILE_ARM64}"

# Clean up for next build
rm -rf ./build/*

# Build for AMD (64 bits)
GOARCH=amd64 GOOS=darwin go build -o ./build/bin-amd64
cp ./assets/darwin/bin/* ./build
cp ./build/bin-amd64 ./build/bin
rm -f ${FILE_AMD64}
hdiutil create -fs HFS+ -srcfolder "./build" -volname "cominnek-${VERSION}-amd64" "${FILE_AMD64}"

# Clean up
rm -rf ./build/*
    