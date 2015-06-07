:: Runs `go test` on all packages.
@echo off
setlocal EnableDelayedExpansion

set GOPATH=%~d0%~p0
cd %GOPATH%src
go test ./... && echo PASS || echo FAIL

endlocal
