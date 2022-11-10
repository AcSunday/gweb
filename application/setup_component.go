package main

import (
	"fmt"

	"go_project_demo/component"
)

// setupComponent 配置组件
func setupComponent(conf string, port int) (err error) {

	// 配置配置组件
	if err = component.SetupConf(conf); err != nil {
		err = fmt.Errorf("component.SetupConf(%s): %w", conf, err)
		return
	}

	// 配置消息日志组件
	if err = component.SetupInfLogger(); err != nil {
		err = fmt.Errorf("component.SetupInfLogger: %w", err)
		return
	}

	// 配置错误日志组件
	if err = component.SetupErrLogger(); err != nil {
		err = fmt.Errorf("component.SetupErrLogger: %w", err)
		return
	}

	// 配置http客户端
	if err = component.SetupHttpClient(); err != nil {
		err = fmt.Errorf("component.SetupHttpClient: %w", err)
		return
	}

	// 配置缓存
	//if err = component.SetupCache(); err != nil {
	//	err = fmt.Errorf("component.SetupCache: %w", err)
	//	return
	//}

	// 配置DB
	//if err = component.SetupDB(); err != nil {
	//	err = fmt.Errorf("component.SetupDB: %w", err)
	//	return
	//}

	// 配置HTTP服务
	if err = component.SetupHttpServer(port); err != nil {
		err = fmt.Errorf("component.SetupHttpServer(%d): %v", port, err)
		return
	}

	return
}
