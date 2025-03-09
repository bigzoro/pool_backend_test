package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"pool/global"
	"pool/initialize"
	"pool/log"
	"pool/routers"
	"pool/scheduler"
	"sync"
	"syscall"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// 初始化配置
	initialize.InitConfig()

	// 初始化日志
	err := initialize.InitLogger()
	if err != nil {
		panic(err)
	}

	// 初始化 mysql 数据库
	initialize.InitDB()

	// 初始化 redis
	initialize.InitRedisDB(ctx)

	// 初始化 routers
	router := routers.InitRouters()

	// 启动定时任务
	var wg sync.WaitGroup
	wg.Add(1)
	go scheduler.StartBlockScheduler(ctx, &wg)
	wg.Add(1)
	go scheduler.StartPoolScheduler(ctx, &wg)
	wg.Add(1)
	go scheduler.MonitorTransfer(ctx, &wg)

	// 启动 gin
	if global.Config.Port == 0 {
		global.Config.Port = 6678
	}
	addr := fmt.Sprintf(":%d", global.Config.Port)
	if err := router.Run(addr); err != nil {
		log.SystemLogger.Errorf("启动服务失败", err.Error())
	}

	// 接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞等待信号
	<-quit
	log.SystemLog().Info("🔴 收到退出信号，正在清理资源...")
	cancel()
	wg.Wait()

	log.SystemLog().Info("✅ 任务退出，服务已安全关闭")
}
