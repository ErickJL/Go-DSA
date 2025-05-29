@echo off
setlocal

REM output and directory bin folder
set OUTPUT_NAME=app.exe
set BIN_DIR=bin

REM create bin folder if not exist
if not exist %BIN_DIR% (
    mkdir %BIN_DIR%
)

REM build app
echo Building Go application...
go build -o %BIN_DIR%\%OUTPUT_NAME% main.go

REM check if build is error
if %errorlevel% neq 0 (
    echo Build gagal.
    exit /b %errorlevel%
)

echo Build success. Run the application...
%BIN_DIR%\%OUTPUT_NAME%

endlocal
REM pause
