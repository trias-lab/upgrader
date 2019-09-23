package main

import (
	"fmt"
	"os"
	"trias-upgrade/lib"
)


func helper() {
	fmt.Printf("|%-6s|%-6s|\n", "upgrade", "--update trias server to lastest version.")
	fmt.Printf("|%-6s|%-6s|\n", "genesis", "--generate basic configuration.")
	fmt.Printf("|%-6s|%-6s|\n", "check", "--check trias server version at local.")
	fmt.Printf("|%-6s|%-6s|\n", "ver", "--show the current version .")
	fmt.Printf("|%-6s|%-6s|\n", "syncdata", "--whether data is synchronized or not.")
	fmt.Printf("|%-6s|%-6s|\n", "new", "--star the new nodes for trias.")

}

func upgrade() {
	lib.AddSource("192.168.1.125")
	lib.GetBin("./requirements.txt","http://192.168.1.125/requirements.txt")
	out:=lib.CmdExec("apt-get update ")
	if out !="sucesss" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: apt-get update ")
	}
	lib.CmdExec("apt-get install -y openssl python3-pip 8lab-zeromq4 ")
	if out !="sucesss" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: apt-get install  ")
	}
	lib.CmdExec("pip3 install -r requirements.txt ")
	if out !="sucesss" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: pip3 install ")
	}

	lib.GetBin("/trias.tar.gz","http://192.168.1.125/packs/files/trias.tar.gz")
	dtar:=lib.CmdExec(`tar zxvf /trias.tar.gz`)
	if dtar !="sucesss" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: get trias base config faild ")
	}
	txout := lib.CmdExec(`tar zxvf /trias.tar.gz`)
	if txout !="sucesss" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: tar zxvf  ")
	}

	lib.Adduser()

	lib.GetBin("/usr/local/bin/tendermint","http://192.168.1.125/packs/files/tendermint")
	lib.GetBin("/usr/local/bin/trias_accs","http://192.168.1.125/packs/files/trias_accs")
	lib.GetBin("/usr/local/bin/triascode_app","http://192.168.1.125/packs/files/triascode_app")
	lib.GetBin("/8lab/blackbox","http://192.168.1.125/packs/files/blackbox")
	lib.GetBin("/8lab/blackbox_agent","http://192.168.1.125/packs/files/blackbox_agent")
	lib.GetBin("/txmodule.tar.gz","http://192.168.1.125/packs/files/txmodule.tar.gz")
	lib.GetBin("/attestation.tar.gz","http://192.168.1.125/packs/files/attestation.tar.gz")

	lib.CmdExec(`tar zxvf /txmodule.tar.gz`)
	lib.CmdExec(`tar zxvf /attestation.tar.gz`)

	lib.CmdExec(`chown -R ubuntu:ubuntu /trias`)
	lib.CmdExec(`chown -R verfiy:root /8lab`)
	lib.CmdExec(`chmod +x /usr/local/bin/tendermint`)
	lib.CmdExec(`chmod +x /usr/local/bin/trias_accs`)
	lib.CmdExec(`chmod +x /usr/local/bin/triascode_app`)
	lib.CmdExec(`chmod +x /8lab/blackbox`)
	lib.CmdExec(`chmod +x /8lab/blackbox_agent`)
	lib.CmdExec(`chown -R verfiy:root /attestation`)
	lib.CmdExec(`chown -R ubuntu:ubuntu /txmodule`)

	lib.CmdExec(`systemctl enable BlackBoxClientinit.service`)
	lib.CmdExec(`systemctl enable Triasinit.service`)

	//lib.CmdExec(`reboot`)
}

func opts() {
	lib.GetBin("/deploy.tar.gz","http://192.168.1.125/packs/files/deploy.tar.gz")
	lib.CmdExec(`tar zxvf /deploy.tar.gz`)
	lib.CmdExec(``)

}

func genesis() {

}

func check() {

}

func ver() {

}

func new() {

}


func syncdata() {
	lib.CmdExec(`reboot`)

}

func main(){
	//fmt.Println(len(os.Args))
	//if len(os.Args)!=2{
	//	helper()
	//} else {
	//	for idx, args := range os.Args {
	//		fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
	//	}
	//}
	if len(os.Args)!=2{
		helper()
		//for idx, args := range os.Args {
		//	fmt.Println("参数" + strconv.Itoa(idx) + ":", args)
		//}
	}
	if len(os.Args)==2{
		if string(os.Args[1])=="upgrade"{
			upgrade()
			fmt.Println("参数:", string(os.Args[1]))
		}
		if string(os.Args[1])=="genesis"{
			genesis()
			fmt.Println("参数:", string(os.Args[1]))
		}
		if string(os.Args[1])=="opts"{
			genesis()
			fmt.Println("参数:", string(os.Args[1]))
		}
		if string(os.Args[1])=="new"{
			new()
			fmt.Println("参数:", string(os.Args[1]))
		}
	}

}

