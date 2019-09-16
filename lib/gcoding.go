package lib

import (
	"net/http"
	"os"
)

const ddir  = "/data/tmp"

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		//os.Chdir("/data/tmp")
		return false, nil
	}
	return false, err
}


func GetBin(wfile string,url string){
	exist, err := PathExists(ddir)
	if err!=nil{
		LogHander("get dir err!",err)
		return
	}
	if exist{
		return
	} else {
		InfoHander("create dir default is /data/tmp")
		err := os.Mkdir(ddir, os.ModePerm)
		if err !=nil{
			LogHander("create dir err!",err)
		}
	}

	fout,err:=os.Create(wfile)
	defer fout.Close()

	if err!=nil{
		LogHander("Create download file failed!",err)
	}

	res,err:=http.Get(url)
	if err!=nil{
		LogHander("http download file failed!",err)
	}

	buf:=make([]byte,1024)
	for{
		size,_:=res.Body.Read(buf)
		if size==0{
			break
		}else {
			fout.Write(buf[:size])
		}
	}
}