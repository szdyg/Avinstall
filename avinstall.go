package main

import (
	"avinstall/download"
	"avinstall/install"
	"flag"
	"fmt"
	"os"
	"path"
)

var InstallAV string



type AvInfo struct {
	avName        string
	fullName      string
	downUrl       string
	installAction func(string)
}

func init() {
	flag.StringVar(&InstallAV, "av", "", "需要安装的杀毒软件名称")
}

var allAV = []AvInfo{
	{"360safe", "360安全卫士", "https://down.360safe.com/setup.exe", install.Install_360safe},
	{"360sd", "360杀毒", "https://down.360safe.com/360sd/360sd_x64_std_7.0.0.1001C.exe", install.Install_360safe},
	{"duba", "金山毒霸", "https://cd001.www.duba.net/duba/install/packages/ever/duba20210603_100_100.exe", install.Install_360safe},
}

func PrintSupportAV() {
	fmt.Println("支持安装的杀毒软件如下")
	for _, av := range allAV {
		fmt.Printf("参数: %s  \t 全称: %s\n", av.avName, av.fullName)
	}
	fmt.Println("实例:  avinstall.exe -av 360safe")
}

func main() {

	flag.Parse()
	if len(InstallAV) == 0 {
		fmt.Println("参数列表:")
		flag.PrintDefaults()
		PrintSupportAV()
		return
	}
	for _, av := range allAV {
		if av.avName == InstallAV {
			tmpDir, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				return
			}

			localpath := path.Join(tmpDir, av.avName+".exe")
			os.Remove(localpath)
			err =download.DownloadFile(av.downUrl, localpath)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("下载完成，开始安装，过程是异步的，请耐心等待。")
			av.installAction(localpath)
		}
	}

}
