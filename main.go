package main

import (
	"syscall"
	"unsafe"
)

var (
	user32      = syscall.NewLazyDLL("user32.dll")
	messageBoxW = user32.NewProc("MessageBoxW")
)

func MessageBox(title, text string) {
	// Преобразование строк в UTF-16/Convert strings to UTF-16
	titlePtr, _ := syscall.UTF16PtrFromString(title)
	textPtr, _ := syscall.UTF16PtrFromString(text)

	//Вызов MessageBox/ Calling MessageBox
	messageBoxW.Call(0,
		uintptr(unsafe.Pointer(textPtr)),
		uintptr(unsafe.Pointer(titlePtr)),
		0)
}

func main() {
	MessageBox("Critical Error", "\nЗдравствуйте, я албанский компьютерный вирус.\n\nИз-за слабого развития технологий в моей стране, к сожалению, я не могу причинить вред вашему компьютеру. Пожалуйста, удалите важные файлы и папки вручную. И отправьте меня по почте друзьям.\n\nСпасибо за сотрудничество!\nС наилучшим пожеланиями, ваш албанский вирус.")
}
