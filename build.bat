@echo off
set GOPATH=%CD%
set CGO_ENABLED=1
set PATH=%PATH%;%GOPATH%\bin
set GOOS=windows
set GOARCH=386
go build -v -o output\%GOOS%-%GOARCH%\MRS-Launcher.exe
fyne package -os windows -icon icon.png -executable output\%GOOS%-%GOARCH% -name MRS-launcher.exe
go build -v -o output\%GOOS%-%GOARCH%\MRS-Launcher.exe
set GOARCH=amd64
go build -v -o output\%GOOS%-%GOARCH%\MRS-Launcher.exe
fyne package -os windows -icon icon.png -executable output\%GOOS%-%GOARCH% -name MRS-launcher.exe
go build -v -o output\%GOOS%-%GOARCH%\MRS-Launcher.exe