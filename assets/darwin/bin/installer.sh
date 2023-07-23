#!/bin/bash

DIRECTORY="/Applications/cominnek.app"

mkdir -p ${DIRECTORY}/Contents/MacOS

rm -f /usr/local/bin/cominnek

UNAME_M=$(uname -m)

if [ "${UNAME_M}" = "arm64" ]; then
    cp ./bin/cominnek-arm64 ${DIRECTORY}/Contents/MacOS/cominnek
fi

if [ "${UNAME_M}" = "x86_64" ]; then
    cp ./bin/cominnek-amd64 ${DIRECTORY}/Contents/MacOS/cominnek
fi

cp ./uninstaller.sh ${DIRECTORY}/uninstaller.sh
ln -s ${DIRECTORY}/Contents/MacOS/cominnek /usr/local/bin/cominnek

echo "cominnek has been installed to ${DIRECTORY}."
echo "To uninstall, run ${DIRECTORY}/uninstaller.sh"