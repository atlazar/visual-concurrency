@echo off

set CGO_ENABLED=1

if "%1" == "cli" (
    goto cli
)

if "%1" == "gui" (
    goto gui
)

echo "unknown command: %1"
goto end

:cli
go build -o .\bin\ .\cmd\cli\
.\bin\cli.exe
goto end

:gui
go build -o .\bin\ .\cmd\gui\
.\bin\gui.exe
goto end

:end