package sleep

import (
	"syscall"
)

var (
	powrProfDll         = syscall.NewLazyDLL("PowrProf.dll")
	user32Dll           = syscall.NewLazyDLL("user32.dll")
	setSuspendStateProc = powrProfDll.NewProc("SetSuspendState")
	lockWorkStationProc = user32Dll.NewProc("LockWorkStation")
)

// Sleep will suspend the system
func Sleep() {
	setSuspendStateProc.Call(0, 0, 0)
}

// Hibernate will Hibernate the system
func Hibernate() {
	setSuspendStateProc.Call(1, 0, 0)
}

// ForceSleep will susped the system without waiting for processes
func ForceSleep() {
	setSuspendStateProc.Call(0, 1, 0)
}

// ForceHibernate will hibernate the ssytem witout waiting for processes
func ForceHibernate() {
	setSuspendStateProc.Call(1, 1, 0)
}

// LockScreen will lock the screen
func LockScreen() {
	lockWorkStationProc.Call()
}
