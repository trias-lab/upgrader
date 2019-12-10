package lib

import (
	"fmt"
	"testing"
)



//test get ip and change configure
func TestGetExternal(t *testing.T)  {

	showip := GetExternal()
	//println(showip)
    fmt.Sprintf(showip)
}


//test get ip and change configure
func TestGetIntranetIp(t *testing.T)  {

	GetIntranetIp()
	//println(showip)
}

//test get ip and change configure
func TestGetPulicIP(t *testing.T)  {

	showip := GetPulicIP()
	println(showip)
}
