package lib

import "os/user"

func Adduser(){

	if _, err := user.Lookup("ubuntu"); err != nil {
		//var cmd *exec.Cmd
		//CmdExec("/usr/sbin/groupadd","ubuntu")
		CmdBash("/usr/sbin/groupadd ubuntu")
		//CmdExec("/usr/sbin/useradd", "-m","-d", "/home/ubuntu", "ubuntu", "-g", "ubuntu")
		CmdBash("/usr/sbin/useradd -m -d /home/ubuntu ubuntu -g ubuntu")

	}

	if _, err := user.Lookup("verfiy"); err != nil {
		//var cmd *exec.Cmd
		//CmdExec("/usr/sbin/groupadd", "verfiy")
		CmdBash("/usr/sbin/groupadd verfiy")
		//CmdExec("/usr/sbin/useradd", "-m", "-d", "/home/verfiy", "verfiy", "-g", "verfiy")
		CmdBash("/usr/sbin/useradd -m -d /home/verfiy verfiy -g root")

	}

}

