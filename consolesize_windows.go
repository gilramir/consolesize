// Copyright (c) 2020 by Gilbert Ramirez <gram@alumni.rice.edu>

package consolesize

import (
	"fmt"

	"golang.org/x/sys/windows"
)


/*
Idea from:
https://stackoverflow.com/questions/6812224/getting-terminal-size-in-c-for-windows
*/

func GetConsoleWidthHeight() (WidthHeight, error) {

	handle, err := windows.GetStdHandle( windows.STD_OUTPUT_HANDLE )
//	handle, err := windows.GetStdHandle( windows.STD_ERROR_HANDLE )
	if err != nil {
		return WidthHeight{}, fmt.Errorf("GetStdHandle: %q", err)
	}

	info := &windows.ConsoleScreenBufferInfo{}

	err = windows.GetConsoleScreenBufferInfo( handle, info )

	if err != nil {
		return WidthHeight{}, fmt.Errorf("GetConsoleScreenBufferInfo: %q", err)
	}

	return WidthHeight{
		Width: int(info.Window.Right - info.Window.Left + 1),
		Height: int(info.Window.Bottom - info.Window.Top + 1),
	}, nil
}
