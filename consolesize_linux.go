// Copyright (c) 2020 by Gilbert Ramirez <gram@alumni.rice.edu>

package consolesize

import (
	"fmt"
	"syscall"
	"unsafe"
)

/*
Idea from:
https://stackoverflow.com/questions/6812224/getting-terminal-size-in-c-for-window://stackoverflow.com/questions/16569433/get-terminal-size-in-go
*/

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func GetConsoleWidthHeight() (WidthHeight, error) {

	ws := &winsize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		return WidthHeight{}, fmt.Errorf("TIOCGWINSZ: errno=%d", errno)
	}

	return WidthHeight{
		Width:  int(ws.Col),
		Height: int(ws.Row),
	}, nil
}
