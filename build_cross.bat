:: Builds the entire source base for all platforms.
:: Make sure to bootstrap your go compiler before you cna run this.
@echo off
setlocal EnableDelayedExpansion

:: This is my local go cross compiler; change it to yours.
set GOROOT=C:\dev\go_cross\
:: TODO add existance check...

set GOPATH=%~d0%~p0
set GOBIN=%GOPATH%bin\cross\
cd %GOPATH%src

:: Compile!
for %%a in (amd64 386) do (
for %%b in (windows linux darwin) do (

set GOARCH=%%a
set GOOS=%%b
echo compiling !GOARCH!/!GOOS!...
%GOROOT%bin\go install ./...

)
)

:: On my machine, windows/386 doesn't get its own directory, so I need to
:: create it myself.
mkdir "%GOBIN%windows_386"
move "%GOBIN%*.exe" "%GOBIN%windows_386\"

endlocal
