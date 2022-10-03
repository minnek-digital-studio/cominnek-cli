VERSION=${1}

mkdir ./build-MacOS/application;
mkdir ./build-MacOS/target;
go mod tidy;
go build -o ./build-MacOS/application;
rm build-MacOS/darwin/Resources/LICENSE.txt;
cp ./LICENSE build-MacOS/darwin/Resources/LICENSE.txt;
bash ./build-MacOS/build-macos-x64.sh cominnek ${VERSION};
