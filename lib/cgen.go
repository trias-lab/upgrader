package lib

import "os/user"

func Adduser(){

	if _, err := user.Lookup("ubuntu"); err != nil {
		//var cmd *exec.Cmd
		CmdExec("/usr/sbin/groupadd","ubuntu")
		CmdExec("/usr/sbin/useradd", "ubuntu", "-g", "ubuntu","-m")

	}

	if _, err := user.Lookup("verfiy"); err != nil {
		//var cmd *exec.Cmd
		CmdExec("/usr/sbin/groupadd", "verfiy")
		CmdExec("/usr/sbin/useradd","verfiy", "-g", "verfiy", "-m")

	}

}

