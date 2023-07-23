VERSION=${1}
DIST_DIR="./dist"
FILE_DMG="${DIST_DIR}/cominnek-${VERSION}.dmg"

if [ -z "${VERSION}" ]; then
    echo "Usage: build.sh <version>"
    exit 1
fi

mkdir -p ${DIST_DIR}
mkdir -p ./build

cp ./assets/darwin/bin/* ./build
mkdir -p ./build/bin
rm -rf ./build/bin/*
go mod tidy;

GOOS=darwin GOARCH=amd64 go build -o ./build/bin/cominnek-amd64;
GOOS=darwin GOARCH=arm64 go build -o ./build/bin/cominnek-arm64;
GOOS=darwin GOARCH=arm go build -o ./build/bin/cominnek-arm;

rm -f ${DIST_DIR}/cominnek-${VERSION}.zip;
rm -f ${FILE_DMG};
hdiutil create -fs HFS+ -srcfolder "./build" -volname "cominnek-${VERSION}" "${FILE_DMG}"

rm -rf ./build/*