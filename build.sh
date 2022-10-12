VERSION=${1}
DIST_DIR="./dist"

if [ -z "${VERSION}" ]; then
    echo "Usage: build.sh <version>"
    exit 1
fi

cp ./build-MacOS/* ./build
mkdir -p ./build/bin
rm -rf ./build/bin/*
go mod tidy;
go build -o ./build/bin;
rm -f ${DIST_DIR}/cominnek-${VERSION}.zip;
hdiutil create -fs HFS+ -srcfolder "./build" -volname "cominnek-${VERSION}" "${DIST_DIR}/cominnek-${VERSION}.dmg"
rm -rf ./build/*