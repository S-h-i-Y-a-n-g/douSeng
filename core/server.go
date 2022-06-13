package core

import (
	"fmt"
	"project/global"
	"project/initialize"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

var LogoContent = `
	 _____   _____   _   _   _____   _   _   _____   __   _   _____  
	|  _  \ /  _  \ | | | | /  ___/ | | | | | ____| |  \ | | /  ___| 
	| | | | | | | | | | | | | |___  | |_| | | |__   |   \| | | |     
	| | | | | | | | | | | | \___  \ |  _  | |  __|  | |\   | | |  _  
	| |_| | | |_| | | |_| |  ___| | | | | | | |___  | | \  | | |_| | 
	|_____/ \_____/ \_____/ /_____/ |_| |_| |_____| |_|  \_| \_____/    
`


func RunWindowsServer() {
	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GSD_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GSD_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf("%s\n", LogoContent)
	fmt.Printf(`
	欢迎使用 抖声
	当前版本:V1.16
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
`, address)
	global.GSD_LOG.Error(s.ListenAndServe().Error())
}
