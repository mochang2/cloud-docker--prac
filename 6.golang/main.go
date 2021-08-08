package main

// practice separating space

import "fmt"
import "os"
import "os/exec"
import "syscall" // need super user privileges to use

func myrun() {
	// fork
	cmd := exec.Command("/proc/self/exe", append([]string{"mychild"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// One way to distinguish the distinct(namespace) is to use hostname(or UTS in Unix)
	// new namespace for hostname execution -> sudo go run main.go run hostname container
	// create ps -> change hostname -> exit -> Ubuntu is not affected
	cmd.SysProcAttr = &syscall.SysProcAttr{ // require systemcall to Kernel for some reasons, in this case: give me a new UTS
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	// syscall.Sethostname([]bytes("bob")) // If I use this operation with 'sudo', my local hostname is also changed
	// To avoid this, I used fork(12th line)

	cmd.Run()
}

func mychild() {
	// variable number of the arguments
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// syscall.Sethostname([]bytes("bob")) // This does not affect the 'real'(not in process) hostname, compared to myRun()

	// to use bash in container. have to copy actual Ubuntu fs dirs in MY_ROOT // skip(no memory space in my VM)
	//syscall.Chroot("./MY_ROOT")
	//syscall.Chdir("/")

	cmd.Run()
}

func myexec() {
	fmt.Println("exec")
}


// such as : docker <cmd>run <container-name>
// go run main.go <cmd> <arguments>
func main() { // ignore the error
	fmt.Printf("Running %v as pid %d\n", os.Args[1:], os.Getpid())

	switch os.Args[1]{
	case "run":
		myrun()
	case "exec":
		myexec()
	case "mychild":
		mychild()
	default:
		panic("help")
	}
}