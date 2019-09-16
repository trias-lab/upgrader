package lib

import (
	"os"
	"os/exec"
	"strings"
)

func CmdExec(cstr string) string{
	cmd:=exec.Command(cstr)

	err :=cmd.Run()
	if err!=nil{
		LogHander(cstr+"cmd exec failed!",err)
		return "failed"
	}
	return "sucesss"
}


func AddSource(apts string){
	df:=os.Remove("/etc/apt/sources.list")
	if df != nil {
		InfoHander("del apt file faild")
	}
	//f,err := os.Create("/etc/apt/sources.list")
	f,err:=os.OpenFile("/etc/apt/sources.list",os.O_RDWR|os.O_CREATE,0666)
	defer f.Close()

	if err !=nil {
		//fmt.Println(err.Error())
		LogHander("Create apt file faild",err)
	}

	conts:="deb [trusted=yes] http://"+apts+"/ octa18 test "
	contb:=[]byte(conts)
	_,err=f.Write(contb)
	if err!=nil {
		LogHander("Write apt file faild",err)
	}
}

//func SetIma() string{
func SetIma(){
	//cmd:=exec.Command("grep 'ima_tcb'  /boot/grub/grub.cfg")

	if fileObj,err:=os.Open("/boot/grub/grub.cfg");err==nil {
		defer fileObj.Close()

		buf:=make([]byte,4096)
		if n,err:=fileObj.Read(buf);err==nil{
			res:=strings.Contains(string(n),"ima_tcb")
			if res!=true{
				sc:=CmdExec(`sed -i "/linux\t/s/$/& ima_tcb ima_template=\"ima\" ima_hash=\"sha1\"/g" /boot/grub/grub.cfg`)
				//return sc
				InfoHander(sc)
				}
			}
		}
	//return "failed"
	}


