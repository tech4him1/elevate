package elevate

import (
    "log"
    "golang.org/x/sys/windows"
    "unsafe"
)

func Elevate(exePath, params, workDir string) uint {
    shell32 := windows.NewLazyDLL("shell32.dll")
    shellExecuteW := shell32.NewProc("ShellExecuteW")

    returnCode, _, _ := shellExecuteW.Call(0,
        uintptr(unsafe.Pointer(strToNullStr("runas"))),
        uintptr(unsafe.Pointer(strToNullStr(exePath))),
        uintptr(unsafe.Pointer(strToNullStr(params))),
        uintptr(unsafe.Pointer(strToNullStr(workDir))),
        1)
    return uint(returnCode)
}

func strToNullStr(s string) *uint16 {
    ptr, err := windows.UTF16PtrFromString(s)
    if err != nil {
        log.Print(err)
    }
    return ptr
}
