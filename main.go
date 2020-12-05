package main

import (
	"DataCertPlatform/db_mysql"
	_ "DataCertPlatform/routers"
	"__MACOSX/DataCertPlatform/blockchain"
	"github.com/astaxie/beego"
)

func main() {
	//先准备一条区块链‘
	blockchain.NewBlockChain()

	//连接数据
	db_mysql.Connect()

	//设置静态资源文件映射
	beego.SetStaticPath("/js","./static/js")
	beego.Run()
}

