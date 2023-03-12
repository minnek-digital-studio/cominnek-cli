VERSION=${1}
DIST_DIR="./dist"
CURERNT_VERSION_FOLDER="cominnek-${VERSION}"
BUILD_DIR="./build/${CURERNT_VERSION_FOLDER}"
FILE=${DIST_DIR}/cominnek-${VERSION}.tar.gz
CURRENT=$(echo "$PWD")
DEB_FILE="${DIST_DIR}/cominnek-${VERSION}.deb"
REPO_URL="https://github.com/Minnek-Digital-Studio/cominnek"

if [ -z "${VERSION}" ]; then
    echo "Usage: build.sh <version>"
    exit 1
fi

if [ ! -d "${DIST_DIR}" ]; then
    echo "making dir ${DIST_DIR}"
    mkdir -p ${DIST_DIR}
fi

rm -f ${FILE}
rm -f ${DEB_FILE}

mkdir -p ./build
rm -rf ./build/*

mkdir -p ${BUILD_DIR}

go mod tidy;
go build -o ${BUILD_DIR} -ldflags=-compressdwarf=false

cd ${BUILD_DIR}
dh_make --createorig --copyright gpl3 -e isaac@minnekdigital.com --single -y

echo "cominnek /usr/bin" > debian/cominnek.install

sed -i -e "s#<insert the upstream URL, if relevant>#${REPO_URL}#g" debian/control
sed -i -e 's/Build-Depends: debhelper-compat (= 13)/&,git,git-flow/g' debian/control
sed -i -e 's/<insert up to 60 chars description>/Create commits and pull requests in an easy way/g' debian/control
sed -i -e "s/ <insert long description, indented with spaces>/License: GPL-3.0/g" debian/control

dpkg-buildpackage -nc -i

cd "${CURRENT}"

mv "./build/cominnek_${VERSION}-1_amd64.deb" ${DEB_FILE}

rm -rf ./build/*