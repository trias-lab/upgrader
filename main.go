package main

import (
	"fmt"
	"os"
	"upgrader/lib"
)

const soip string ="192.168.1.125"
const sourl string ="http://192.168.1.125/"
const dlurl string ="http://192.168.1.125/packs/files/"
const rbn int  = 6

//help print
func helper() {
	//rollback
	fmt.Printf("|%-6s|%-6s|\n", "rollback", "--rollback trias data to the block height.")
	fmt.Printf("|%-6s|%-6s|\n", "upgrade", "--Update trias server to lastest version.")
	fmt.Printf("|%-6s|%-6s|\n", "genesis", "--Generate basic configuration.")
	fmt.Printf("|%-6s|%-6s|\n", "check", "--Check trias server version at local.")
	fmt.Printf("|%-6s|%-6s|\n", "ver", "--Show the current version .")
	fmt.Printf("|%-6s|%-6s|\n", "syncdata", "--Whether data is synchronized or not.")
	fmt.Printf("|%-6s|%-6s|\n", "new", "--Star the new nodes for trias.")
	fmt.Printf("|%-6s|%-6s|\n", "clean", "--Clear the all files of the local node.")
	fmt.Printf("|%-6s|%-6s|\n", "cdata", "--Clear the all data of the local node.")

}

//it does not affect local data files at this  Upgrade
func upgrade() {
	//clean old bin
	//syncdata()
	//add user
	lib.CmdBash("/etc/init.d/Trias stop")
	lib.CmdBash("/etc/init.d/BlackBoxClient stop")

	lib.Adduser()
	fmt.Println(".........................check and star user finished.")

	//add apt and pip source,install and setup packages
	lib.AddSource(soip)
	fmt.Println(".........................add source finished.")
	lib.GetBin("./requirements.txt",dlurl+"requirements.txt")
	//out:=lib.CmdExec("apt-get", "update")
	out:=lib.CmdBash("apt-get update")
	if out =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: apt-get update ")
		fmt.Println(".........................apt-get have exception.")
	}
	fmt.Println(".........................apt-get finished.")

	//lib.CmdExec("apt-get", "install", "-y", "openssl", "python3-pip", "8lab-zeromq4 ")
	installOut:=lib.CmdBash("apt-get install -y libgmp-dev openssl python3-pip 8lab-zeromq4")

	if installOut =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: apt-get install  ")
		fmt.Println(".........................install packs exception.")
	}
	fmt.Println(".........................install packs finished.")

	//pip3 install --no-index --trusted-host 192.168.1.125 --find-links=http://{{lab8_apt}}/packs/pypi -r /tmp/requirements.txt
	plink:=" --find-links="+sourl+"packs/pypi "
	//lib.CmdExec("pip3", "install", "--no-index", "--trusted-host", soip, plink, "-r", "requirements.txt ")
	pipyOut:=lib.CmdBash("pip3 install --no-index --trusted-host " + soip + plink + " -r requirements.txt ")

	if pipyOut =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: pip3 install ")
		fmt.Println(".........................install pip exception.")
	}
	fmt.Println(".........................install pip finish.")

	//download dir structure dlurl
	lib.GetBin("/txmodule.tar.gz",dlurl+"txmodule.tar.gz")
	lib.GetBin("/attestation.tar.gz",dlurl+"attestation.tar.gz")
	lib.GetBin("/trias.tar.gz",dlurl+"trias.tar.gz")
	fmt.Println(".........................get  dir structure finished.")

	//dtar:=lib.CmdExec("tar", "zxvf", "/trias.tar.gz","-C","/")

	//create dir and unzip
	dtar:=lib.TarZxvf("/trias.tar.gz")
	if dtar !="sucesss" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: tar zxvf trias faild ")
		fmt.Println(".........................unzip structure exception.")
	}

	txtar:=lib.TarZxvf("/txmodule.tar.gz")
	if txtar !="sucesss" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: tar zxvf txmodule faild ")
		fmt.Println(".........................unzip structure exception.")
	}

	attar:=lib.TarZxvf("/attestation.tar.gz")
	if attar !="sucesss" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: tar zxvf attestation faild ")
		fmt.Println(".........................unzip structure exception.")
	}
	fmt.Println(".........................unzip structure finished.")

	//download key bin and set configure
	lib.GetBin("/usr/local/bin/tendermint",dlurl+"tendermint10")
	lib.GetBin("/usr/local/bin/trias_accs",dlurl+"trias_accs")
	lib.GetBin("/usr/local/bin/triascode_app",dlurl+"triascode_app")
	lib.GetBin("/8lab/blackbox",dlurl+"blackbox")
	lib.GetBin("/8lab/blackbox_agent",dlurl+"blackbox_agent")
	lib.GetBin("/8lab/log/pk",dlurl+"pk")
	lib.GetBin("/8lab/log/vk",dlurl+"vk")
	lib.GetBin("/trias/.ethermint/tendermint/config/config.toml",dlurl+"config.toml")
	lib.SetTmHostname()
	lib.GetBin("/trias/p2p/p2p.json",dlurl+"p2p.json")

	fmt.Println(".........................key bin and set configure finished.")


	//change chmod and chown
	//lib.CmdExec("chown", "-R", "ubuntu:ubuntu", "/trias")
	lib.CmdBash("chown -R ubuntu:ubuntu /trias")
	//lib.CmdExec("chown", "-R", "verfiy:root", "/8lab")
	lib.CmdBash("chown -R verfiy:root /8lab")
	//lib.CmdExec("chmod", " +x ", "/usr/local/bin/tendermint")
	lib.CmdBash("chmod  +x  /usr/local/bin/tendermint")
	//lib.CmdExec("chmod", " +x ", "/usr/local/bin/trias_accs")
	lib.CmdBash("chmod  +x  /usr/local/bin/trias_accs")
	//lib.CmdExec("chmod", " +x ", "/usr/local/bin/triascode_app")
	lib.CmdBash("chmod  +x  /usr/local/bin/triascode_app")
	//lib.CmdExec("chmod", " +x ", "/8lab/blackbox")
	lib.CmdBash("chmod  +x  /8lab/blackbox")
	//lib.CmdExec("chmod", " +x ", "/8lab/blackbox_agent")
	lib.CmdBash("chmod  +x  /8lab/blackbox_agent")
	//lib.CmdExec("chown", " -R", " verfiy:root ", "/attestation")
	lib.CmdBash("chown  -R  verfiy:root  /attestation")
	//lib.CmdExec("chown ", "-R ", "ubuntu:ubuntu", " /txmodule")
	lib.CmdBash("chown -R ubuntu:ubuntu  /txmodule")
	fmt.Println(".........................change chmod finished.")

	//set start scripts and ima status
	lib.CmdBash("systemctl enable BlackBoxClientinit.service")
	lib.CmdBash("systemctl enable Triasinit.service")
	lib.SetIma()
	fmt.Println(".........................set start scripts and ima status finished.")

	//return ver status
	fmt.Println("upgrade trias node setup finish!")

}

func opts() {
	lib.GetBin("/deploy.tar.gz","http://192.168.1.125/packs/files/deploy.tar.gz")
	lib.CmdExec(`tar zxvf /deploy.tar.gz`)
	lib.CmdExec(``)

}

//sync data form zero or genesis status
func genesis() {
	lib.GetBin("/trias/.ethermint/tendermint/config/genesis.json",dlurl+"genesis.json")
	lib.SetTmHostname()

	fmt.Println(".........................set genesis finished.")

}

func check() {

}

func ver() {

}

func clean() {
	//rmout:=lib.CmdExec("rm", "-rf", "/8lab", "/trias*", "/var/log/8lab/", "/txmodule*","/attestation*")
	rmout:=lib.CmdBash("rm -rf /8lab /trias* /var/log/8lab/ /txmodule* /attestation*")
	if rmout =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: rm ")
	}
	fmt.Println(".........................clean finished.")
}

func new() {
	//add user
	lib.Adduser()
	fmt.Println(".........................check and star user finished.")

	//add apt and pip source,install and setup packages
	lib.AddSource(soip)
	fmt.Println(".........................add source finished.")
	lib.GetBin("./requirements.txt",dlurl+"requirements.txt")
	//out:=lib.CmdExec("apt-get", "update")
	out:=lib.CmdBash("apt-get update")
	if out =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: apt-get update ")
		fmt.Println(".........................apt-get have exception.")
	}
	fmt.Println(".........................apt-get finished.")

	//lib.CmdExec("apt-get", "install", "-y", "openssl", "python3-pip", "8lab-zeromq4 ")
	installOut:=lib.CmdBash("apt-get install -y libgmp-dev openssl python3-pip 8lab-zeromq4")

	if installOut =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: apt-get install  ")
		fmt.Println(".........................install packs exception.")
	}
	fmt.Println(".........................install packs finished.")

	//pip3 install --no-index --trusted-host 192.168.1.125 --find-links=http://{{lab8_apt}}/packs/pypi -r /tmp/requirements.txt
	plink:=" --find-links="+sourl+"/packs/pypi "
	//lib.CmdExec("pip3", "install", "--no-index", "--trusted-host", soip, plink, "-r", "requirements.txt ")
	pipyOut:=lib.CmdBash("pip3 install --no-index --trusted-host " + soip + plink + " -r requirements.txt ")

	if pipyOut =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: pip3 install ")
		fmt.Println(".........................install pip exception.")
	}
	fmt.Println(".........................install pip finish.")

	//download dir structure dlurl
	lib.GetBin("/txmodule.tar.gz",dlurl+"txmodule.tar.gz")
	lib.GetBin("/attestation.tar.gz",dlurl+"attestation.tar.gz")
	lib.GetBin("/trias.tar.gz",dlurl+"trias.tar.gz")
	fmt.Println(".........................get  dir structure finished.")

	//dtar:=lib.CmdExec("tar", "zxvf", "/trias.tar.gz","-C","/")

	//create dir and unzip
	dtar:=lib.TarZxvf("/trias.tar.gz")
	if dtar !="sucesss" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: tar zxvf trias faild ")
		fmt.Println(".........................unzip structure exception.")
	}

	txtar:=lib.TarZxvf("/txmodule.tar.gz")
	if txtar !="sucesss" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: tar zxvf txmodule faild ")
		fmt.Println(".........................unzip structure exception.")
	}

	attar:=lib.TarZxvf("/attestation.tar.gz")
	if attar !="sucesss" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: tar zxvf attestation faild ")
		fmt.Println(".........................unzip structure exception.")
	}
	fmt.Println(".........................unzip structure finished.")

	//download key bin and set configure
	lib.GetBin("/usr/local/bin/tendermint",dlurl+"tendermint10")
	lib.GetBin("/usr/local/bin/trias_accs",dlurl+"trias_accs")
	lib.GetBin("/usr/local/bin/triascode_app",dlurl+"triascode_app")
	lib.GetBin("/8lab/blackbox",dlurl+"blackbox")
	lib.GetBin("/8lab/blackbox_agent",dlurl+"blackbox_agent")
	lib.GetBin("/8lab/log/pk",dlurl+"pk")
	lib.GetBin("/8lab/log/vk",dlurl+"vk")
	lib.GetBin("/trias/.ethermint/tendermint/config/config.toml",dlurl+"config.toml")
	lib.SetTmHostname()
	genesis()
	lib.GetBin("/trias/p2p/p2p.json",dlurl+"p2p.json")

	fmt.Println(".........................key bin and set configure finished.")


	//change chmod and chown
	//lib.CmdExec("chown", "-R", "ubuntu:ubuntu", "/trias")
	lib.CmdBash("chown -R ubuntu:ubuntu /trias")
	//lib.CmdExec("chown", "-R", "verfiy:root", "/8lab")
	lib.CmdBash("chown -R verfiy:root /8lab")
	//lib.CmdExec("chmod", " +x ", "/usr/local/bin/tendermint")
	lib.CmdBash("chmod  +x  /usr/local/bin/tendermint")
	//lib.CmdExec("chmod", " +x ", "/usr/local/bin/trias_accs")
	lib.CmdBash("chmod  +x  /usr/local/bin/trias_accs")
	//lib.CmdExec("chmod", " +x ", "/usr/local/bin/triascode_app")
	lib.CmdBash("chmod  +x  /usr/local/bin/triascode_app")
	//lib.CmdExec("chmod", " +x ", "/8lab/blackbox")
	lib.CmdBash("chmod  +x  /8lab/blackbox")
	//lib.CmdExec("chmod", " +x ", "/8lab/blackbox_agent")
	lib.CmdBash("chmod  +x  /8lab/blackbox_agent")
	//lib.CmdExec("chown", " -R", " verfiy:root ", "/attestation")
	lib.CmdBash("chown  -R  verfiy:root  /attestation")
	//lib.CmdExec("chown ", "-R ", "ubuntu:ubuntu", " /txmodule")
	lib.CmdBash("chown -R ubuntu:ubuntu  /txmodule")
	fmt.Println(".........................change chmod finished.")

	//set start scripts and ima status
	lib.CmdBash("systemctl enable BlackBoxClientinit.service")
	lib.CmdBash("systemctl enable Triasinit.service")
	lib.SetIma()
	fmt.Println(".........................set start scripts and ima status finished.")

	//return ver status
	fmt.Println("new trias node setup finish!")
}


func syncdata() {
	rmtx:=lib.CmdBash("rm -rf /txmodule/data/* ")
	if rmtx =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: rm tx data ")
	}
	rmtm:=lib.CmdBash("rm -rf /trias/.ethermint/tendermint/data/*.db ")
	if rmtm =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: rm tm data ")
	}
	rmwal:=lib.CmdBash("rm -rf /trias/.ethermint/tendermint/data/*.wal ")
	if rmwal =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: rm wal data ")
	}
	fmt.Println(".........................data clean finished.")

	sedstate:=lib.CmdBash(`sed -i 's#\"height\": \"[0-9]*\"#\"height\": \"0\"#' /trias/.ethermint/tendermint/data/priv_validator_state.json  `)
	if sedstate =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: sed state data ")
	}
	fmt.Println(".........................block state  rsync finished.")

}

func rollback() {
	rmtx:=lib.CmdBash("rm -rf /txmodule/data/* ")
	if rmtx =="failed" {
		//fmt.Println(err.Error())
		lib.InfoHander("exec faild: rm tx data ")
	}

	fmt.Println(".........................block state  rsync finished.")

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
			fmt.Println("upgrade down, please restart services.")
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
			fmt.Println("exec:", string(os.Args[1]))
			fmt.Println("create new node down, please reboot.")
		}
		if string(os.Args[1])=="clean"{
			clean()
			fmt.Println("exec:", string(os.Args[1]))
			fmt.Println("clean down,please reboot.")
		}
		if string(os.Args[1])=="cdata"{
			syncdata()
			fmt.Println("exec:", string(os.Args[1]))
			fmt.Println("clean data down,please restart services.")
		}
		if string(os.Args[1])=="rollback"{
			rollback()
			fmt.Println("exec:", string(os.Args[1]))
			fmt.Println("data rollback finished, please restart services.")
		}
	}

}

