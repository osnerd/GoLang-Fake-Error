// this example is basically a " retry, cancel error nothing much crazy "

package lib

import (
	"syscall"
	"unsafe"
)

const (
	MB_ICONERROR   = 0x10
	MB_RETRYCANCEL = 0x05
)

func ShowError(title, message string) {
	titlePtr, _ := syscall.UTF16PtrFromString(title)
	messagePtr, _ := syscall.UTF16PtrFromString(message)
	_, _, _ = syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		0,
		uintptr(unsafe.Pointer(messagePtr)),
		uintptr(unsafe.Pointer(titlePtr)),
		uintptr(MB_ICONERROR|MB_RETRYCANCEL),
	)
}
