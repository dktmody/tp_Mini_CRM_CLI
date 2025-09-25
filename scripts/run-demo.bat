@echo off
REM Wrapper simple pour lancer la démo PowerShell sous Windows
powershell -NoProfile -ExecutionPolicy Bypass -Command "& '%~dp0run-demo.ps1'"
echo Demo script launched.
pause
