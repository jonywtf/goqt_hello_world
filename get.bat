@echo off
call setvars.bat
go get -d -v ./%APP_PACKGE_PATH%
