@echo off
set GOARCH=386
set GOOS=windows
set CGO_ENABLED=1
call setvars.bat

set TARGET_PATH=%ROOTDIR%bin\hello_world.exe
cd %ROOTDIR%%APP_PACKGE_PATH%
go build -i -o %TARGET_PATH% -pkgdir %PKGDIR%
cd %ROOTDIR%
