# Русский

Программа: Albanian Virus (MessageBox Demo)
--------------------------------------------

Описание:
----------
Это демонстрационная программа на Go, которая вызывает
окно Windows MessageBox с шутливым сообщением.

Программа использует WinAPI (функцию MessageBoxW из user32.dll)
через пакет syscall.

При запуске показывает окно с сообщением, без запуска терминала
(консольное окно скрыто).

Инструкция по сборке:
-----------------------

1. Убедитесь, что установлен Go:
   https://go.dev/dl/

2. Сохраните исходный код в файл, например:
   main.go

3. Откройте командную строку (CMD или PowerShell) и выполните:

`**go build -ldflags="-H windowsgui" -o AlbanianVirus.exe main.go**`

Это создаст исполняемый файл AlbanianVirus.exe **без появления терминала**.

Примечания:
-----------
- Работает только на Windows.
- Программа не делает ничего вредоносного, это просто шутливое сообщение.
- Не отправляйте исполняемый файл другим людям без их согласия.

# English
Program: Albanian Virus (MessageBox Demo)
--------------------------------------------

Description:
----------
This is a Go demo program that calls a
Windows MessageBox window with a funny message.

The program uses WinAPI (the MessageBoxW function from user32.dll)
via the syscall package.

When launched, it shows a message box without starting a terminal
(the console window is hidden).

Build instructions:
-----------------------

1. Make sure Go is installed:
   https://go.dev/dl/

2. Save the source code to a file, for example:
   main.go

3. Open a command prompt (CMD or PowerShell) and run:

`**go build -ldflags="-H windowsgui" -o AlbanianVirus.exe main.go**
`

This will create an executable file AlbanianVirus.exe **without showing the terminal**.

Notes:
-----------
- Only works on Windows.
- The program does not do anything malicious, it is just a joke message.
- Do not send the executable to other people without their consent.