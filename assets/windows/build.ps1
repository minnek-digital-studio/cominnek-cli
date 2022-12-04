$setupFile="$pwd\assets\windows\bin\setup.iss"

go mod tidy; go build -o .\Build

echo "Reading $setupFile"
Start-Process -Wait -FilePath "C:\Program Files (x86)\Inno Setup 6.\ISCC.exe" -Argument "/q $setupFile"

echo "Build complete"