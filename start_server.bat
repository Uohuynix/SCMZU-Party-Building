@echo off
echo Starting Party Building Management System...
echo.
echo Make sure MySQL is running and the database 'party_school' exists.
echo.
echo Default configuration:
echo - Database: localhost:3306
echo - Username: root
echo - Password: Hjh994600!
echo - Database Name: party_school
echo - Server Port: 8080
echo.
echo Press any key to start the server...
pause > nul
echo.
echo Starting server...
main.exe