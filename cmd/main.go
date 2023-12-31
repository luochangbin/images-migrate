package main

import (
	"flag"
	"fmt"
	"github.com/AliyunContainerService/image-syncer/pkg/client"
	"github.com/AliyunContainerService/image-syncer/pkg/utils"
	"images-migrate/pkg"
	"os"
)

var (
	logPath, configFile, namespaces string

	authFile = "auth.json"

	imageFile = "images.json"

	procNum, retries int

	forceUpdate bool
)

func init() {
	flag.StringVar(&configFile, "config", "config.yaml", "config file path")
	flag.StringVar(&logPath, "log", "", "log file path (default in os.Stderr)")
	flag.StringVar(&namespaces, "namespaces", "", "the namespaces that you want to sync,multiple values use ','. by default, all are sync")
	flag.IntVar(&procNum, "proc", 5, "numbers of working goroutines")
	flag.IntVar(&retries, "retries", 2, "times to retry failed task")
	flag.BoolVar(&forceUpdate, "force", false, "force update manifest whether the destination manifest exists")
}

func main() {
	flag.Parse()
	config, err := pkg.ReadConfigFromFile(configFile)
	if err != nil {
		fmt.Println("从配置文件读取失败: ", err)
		os.Exit(1)
	}

	if err := pkg.GenAuthFile(authFile, config); err != nil {
		fmt.Println("生成认证文件失败: ", err)
		os.Exit(1)
	}
	if err := pkg.GenImagesFile(imageFile, namespaces, config); err != nil {
		fmt.Println("获取镜像列表失败: ", err)
		os.Exit(1)
	}
	client, err := client.NewSyncClient("", authFile, imageFile, logPath, procNum, retries,
		utils.RemoveEmptyItems([]string{}), utils.RemoveEmptyItems([]string{}), forceUpdate)
	if err != nil {
		fmt.Println("init sync client error: ", err)
		os.Exit(1)
	}
	client.Run()
}
