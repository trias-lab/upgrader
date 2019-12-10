package lib

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CmdExec(cstr ...string) string{
	strs:=""
	for _, arge:=range cstr  {
		strs+=arge+" "
	}

	cmd:=exec.Command(strs)

	err :=cmd.Run()
	if err != nil {
		LogHander(strs+"cmd exec failed!", err)
		return "failed"
	}
	return "sucesss"
}

func CmdBash(cstr string) string{
	//strs:=""

	out,err:=exec.Command("/bin/bash", "-c", cstr).Output()

	if err != nil {
		LogHander(cstr+"cmd exec failed!", err)
		return "failed"
	}
	return string(out)
}

func CmdStr(cstr ...string) string{
	strs:=""
	for _, arge:=range cstr  {
		strs+=arge+" "
	}

	cmd:=exec.Command(strs)

	err :=cmd.Start()
	if err != nil {
		LogHander(strs+"cmd exec failed!", err)
		return "failed"
	}
	errc:=cmd.Wait()
	if errc!=nil{
		LogHander(strs+"cmd exec failed!", errc)
		return "failed"
	}
	return cmd.ProcessState.String()
}

func TarZxvf(str string) string {
	cmd:=exec.Command("tar","zxvf", str,"-C","/")
	err :=cmd.Run()
	if err != nil {
		LogHander("cmd exec failed!", err)
		return "failed"
	}
	return "sucesss"
}

func Chmod(str string) string {
	cmd:=exec.Command("tar","zxvf", str,"-C","/")
	err :=cmd.Run()
	if err != nil {
		LogHander("cmd exec failed!", err)
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
				//sc:=CmdExec(`sed -i "/linux\t/s/$/& ima_tcb ima_template=\"ima\" ima_hash=\"sha1\"/g" /boot/grub/grub.cfg`)
				sc:=CmdBash(`sed -i "/linux\t/s/$/& ima_tcb ima_template=\"ima\" ima_hash=\"sha1\"/g" /boot/grub/grub.cfg`)
				//return sc
				InfoHander(sc)
				}
			}
		}
	//return "failed"
	}

//changes tm config
func SetTmHostname(){
	//get the hostname
	//hn:=CmdStr("hostname")
	hn:= strings.Replace(CmdBash("hostname"),"\n","",-1)
	//if fileObj,err:=os.Open("/trias/.ethermint/tendermint/config/config.toml");err==nil
	input,err:=ioutil.ReadFile("/trias/.ethermint/tendermint/config/config.toml")
	if err!=nil{
		LogHander("read tm config err: ",err)
	}

	//content:=strings.Replace(string(input),"\n","",-1)
	content:=string(input)
	newcontent:=strings.Replace(content,"ubt18-trias-dag-141",hn,-1)

	errw:=ioutil.WriteFile("/trias/.ethermint/tendermint/config/config.toml",[]byte(newcontent),0)
	if errw!=nil{
		LogHander("wirte tm config err: ",errw)
	}
	//lines:=strings.Split(string(input),"\n")
	//for i,line:=range lines{
	//	if strings.Contains(line,"ubt18-trias-dag-141"){
	//		lines[i]=
	//	}
	//}
}


//changes blackbox config
func SetBlackbboxConf(){
	//get the conncet ip
	hn:=GetPulicIP()
	//hn:= strings.Replace(CmdBash("hostname"),"\n","",-1)
	input,err:=ioutil.ReadFile("/8lab/conf/configure.json")
	if err!=nil{
		LogHander("read blackbox config err: ",err)
	}

	//content:=strings.Replace(string(input),"\n","",-1)
	content:=string(input)
	newcontent:=strings.Replace(content,"192.168.1.141",hn,-1)

	errw:=ioutil.WriteFile("/8lab/conf/configure.json",[]byte(newcontent),0)
	if errw!=nil{
		LogHander("wirte blackbox config err: ",errw)
	}
}

type ReplaceHandle struct {
	Root    string //根目录
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHandle) DoWrok() error {
	return filepath.Walk(h.Root, h.walkCallback)
}

func (h ReplaceHandle) walkCallback(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		//fmt.Pringln("DIR:",path)
		return nil
	}
	//文件类型需要进行过滤
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)
	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)
	//重新写入
	ioutil.WriteFile(path, []byte(newContent), 0)
	return err
}

//test coding
//func main() {
//	flag.Parse()
//	hd := ReplaceHandle{
//		Root:    "/trias/.ethermint/tendermint/config/config.toml",
//		OldText: "ubt18-trias-dag-141",
//		NewText: "ubt18-trias-dag-142",
//	}
//	err := hd.DoWrok()
//	if err == nil {
//		fmt.Println("done!")
//	} else {
//		fmt.Println("error:", err.Error())
//	}
//}
