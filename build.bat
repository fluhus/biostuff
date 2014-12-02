:: Builds the entire source base for the current platform.
@echo off
setlocal EnableDelayedExpansion

set GOPATH=%~d0%~p0
cd %GOPATH%src
go install ./...

endlocal
