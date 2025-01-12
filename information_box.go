// 0x40 Explanation:
// 0x40 corresponds to the MB_ICONINFORMATION flag.
// It adds an information icon (ℹ️) to the message box.

package session

import (
	"syscall"
	"unsafe"
)

func DisplaySuccess() {
	var title, text *uint16
	title, _ = syscall.UTF16PtrFromString("INFORMATION!")
	text, _ = syscall.UTF16PtrFromString("github.com/osnerd.")
	syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(0,
		uintptr(unsafe.Pointer(text)),
		uintptr(unsafe.Pointer(title)),
		0x40)
}
