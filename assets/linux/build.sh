VERSION=${1}
DIST_DIR="./dist"
FILE=${DIST_DIR}/cominnek-${VERSION}.tar.gz

if [ -z "${VERSION}" ]; then
    echo "Usage: build.sh <version>"
    exit 1
fi

if [ ! -d "${DIST_DIR}" ]; then
    echo "making dir ${DIST_DIR}"
    mkdir -p ${DIST_DIR}
fi

rm -f ${FILE}

cp ./assets/linux/bin/* ./build
mkdir -p ./build/bin
rm -rf ./build/bin/*
go mod tidy;
go build -o ./build/bin

tar -czvf ${FILE} -C ./build .

# tar -czvf ${FILE} ./build/*

rm -rf ./build/*