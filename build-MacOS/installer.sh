#!/bin/bash

DIRECTORY="/Applications/cominnek.app"

mkdir -p ${DIRECTORY}/Contents/MacOS

rm -f /usr/local/bin/cominnek
cp ./bin/cominnek ${DIRECTORY}/Contents/MacOS/cominnek
cp ./uninstaller.sh ${DIRECTORY}/uninstaller.sh
ln -s ${DIRECTORY}/Contents/MacOS/cominnek /usr/local/bin/cominnek

echo "cominnek has been installed to ${DIRECTORY}."
echo "To uninstall, run ${DIRECTORY}/uninstaller.sh"