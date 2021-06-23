package install

import (
	"fmt"
	"syscall"
)

var (
	kernel32, _        = syscall.LoadLibrary("kernel32.dll")
	pCreateProcessW, _ = syscall.GetProcAddress(kernel32, "CreateProcessW")
)

func Install_360safe(FilePath string) {

	var sI syscall.StartupInfo
	var pI syscall.ProcessInformation
	var err error

	err = syscall.CreateProcess(
		syscall.StringToUTF16Ptr(FilePath),
		syscall.StringToUTF16Ptr("/S"),
		nil,
		nil,
		true,
		0,
		nil,
		nil,
		&sI,
		&pI)
	if err != nil {
		fmt.Println("error:",err)
	}
}
