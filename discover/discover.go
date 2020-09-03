package discover

import (
	"context"
	stars_registry_center "github.com/stars-palace/stars-registry-center"
)

/**
 * Copyright (C) @2020 hugo network Co. Ltd
 *
 * @author: hugo
 * @version: 1.0
 * @date: 2020/8/27
 * @time: 21:37
 * @description:
 */

//获取实例的查询参数
type ServerInstancesParam struct {
	ServiceName string
	GroupName   string // 默认值DEFAULT_GROUP
	Clusters    []string
}

// Discover register/deregister service
// Discover impl should control rpc timeout
type Discover interface {
	GetServerInstance(con context.Context, param *ServerInstancesParam) ([]*stars_registry_center.ServiceInfo, error)
}

// Nop Discover, used for local development/debugging
// 用于本地开发 不进行注册
type Nop struct{}

// RegisterService ...
func (n Nop) GetServerInstance(con context.Context, param *ServerInstancesParam) ([]*stars_registry_center.ServiceInfo, error) {
	return nil, nil
}
