package dalayqueue

import (
	"dalayqueue/config"
	"dalayqueue/delayqueue"
	"dalayqueue/routers"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Cmd 应用入口Command
type Cmd struct{}

var (
	version    bool
	configFile string
)

const (
	// AppVersion 应用版本号
	AppVersion = "0.4"
)

// Run 运行应用
func (cmd *Cmd) Run() {

	// 解析命令行参数
	cmd.parseCommandArgs()
	if version {
		fmt.Println(AppVersion)
		welcome()
		os.Exit(0)
	}
	// 初始化配置
	config.Init(configFile)
	// 初始化队列
	delayqueue.Init()

	// 运行web server
	cmd.runWeb()
}

// 解析命令行参数
func (cmd *Cmd) parseCommandArgs() {
	// 配置文件
	flag.StringVar(&configFile, "c", "", "./delay-queue -c /path/to/delay-queue.conf")
	// 版本
	flag.BoolVar(&version, "v", false, "./delay-queue -v")
	flag.Parse()
}

// 欢迎信息
func welcome() {
	logo := `
_____   _____   _           ___  __    __  _____    
|  _  \ | ____| | |         /   | \ \  / / /  _  \   
| | | | | |__   | |        / /| |  \ \/ /  | | | |   
| | | | |  __|  | |       / / | |   \  /   | | | |   
| |_| | | |___  | |___   / /  | |   / /    | |_| |_  
|_____/ |_____| |_____| /_/   |_|  /_/     \_______| 
	`
	fmt.Println(logo)
	fmt.Println("Service:		delayq")
	fmt.Println("Version:		" + AppVersion)
}

// 运行Web Server
func (cmd *Cmd) runWeb() {
	http.HandleFunc("/push", routers.Push)
	http.HandleFunc("/pop", routers.Pop)
	http.HandleFunc("/finish", routers.Delete)
	http.HandleFunc("/delete", routers.Delete)
	http.HandleFunc("/get", routers.Get)

	log.Printf("listen %s\n", config.Setting.BindAddress)
	err := http.ListenAndServe(config.Setting.BindAddress, nil)
	log.Fatalln(err)
}
