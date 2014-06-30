if [%1]==[p] echo MAIN: try.go && goto :end
move try.go try.goo >nul && echo REVERTED: try.go 
:end
