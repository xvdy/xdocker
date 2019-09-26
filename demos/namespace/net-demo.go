package namespace

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/*
# net namespace

in process:
ip a
ifconfig

n host:
cd /proc/3034/ns
readlink net
*/

func main() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNET,
		//Credential: &syscall.Credential{Uid: uint32(1), Gid: uint32(1)},
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
