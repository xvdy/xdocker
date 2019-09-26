package namespace

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/*
# ipc namespace
system v ipc
posix message queues

ipcs -q
ipcmk -Q
*/

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
