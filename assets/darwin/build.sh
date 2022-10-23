VERSION=${1}
DIST_DIR="./dist"
FILE_DMG="${DIST_DIR}/cominnek-${VERSION}.dmg"

if [ -z "${VERSION}" ]; then
    echo "Usage: build.sh <version>"
    exit 1
fi

cp ./assets/darwin/bin/* ./build
mkdir -p ./build/bin
rm -rf ./build/bin/*
go mod tidy;
go build -o ./build/bin;
rm -f ${DIST_DIR}/cominnek-${VERSION}.zip;
rm -f ${FILE_DMG};
hdiutil create -fs HFS+ -srcfolder "./build" -volname "cominnek-${VERSION}" "${FILE_DMG}"
rm -rf ./build/*