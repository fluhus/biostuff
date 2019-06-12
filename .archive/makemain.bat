:: Switches 'goo' and 'go' files for main.
@echo off

:: Revert script
set _REVERT=makemain.r.bat

:: Check if there exists a parameter
if [%1]==[] goto :print

set _OLD=%1
set _NEW=%_OLD:~0,-1%

:: Check if target exists
if not exist %_OLD% goto :error

:: Revert previous target if exists
if not exist %_REVERT% goto :go
call %_REVERT%

:go
:: Turn goo into go
move %_OLD% %_NEW% >nul
echo NEW MAIN: %_NEW%

:: Create reverter
echo if [%%1]==[p] echo MAIN: %_NEW% ^&^& goto :end>%_REVERT%
echo move %_NEW% %_OLD% ^>nul ^&^& echo REVERTED: %_NEW% >>%_REVERT%
echo :end>>%_REVERT%

goto :end

:print
if exist %_REVERT% call %_REVERT% p && goto :end
echo No main file
goto :end

:error
echo File not found: %_OLD%

:end
