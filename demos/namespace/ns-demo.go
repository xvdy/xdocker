package namespace

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/*
# uts namespace
hostname domainname

pid=`echo $$`
readlink /proc/$pid/ns/uts

hostname -b xudy
hostname
*/

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
