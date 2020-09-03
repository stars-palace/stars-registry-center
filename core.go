package stars_registry_center

import (
	"context"
	"fmt"
)

/**
 * Copyright (C) @2020 hugo network Co. Ltd
 *
 * @author: hugo
 * @version: 1.0
 * @date: 2020/9/3
 * @time: 23:38
 * @description:
 */

// ServiceInfo ...
type ServiceInfo struct {
	Name        string
	Scheme      string
	IP          string
	Port        int
	Weight      float64
	Enable      bool
	Healthy     bool
	Ephemeral   bool
	Metadata    map[string]string
	Region      string
	Zone        string
	GroupName   string
	ClusterName string
}

// Label ...
func (si ServiceInfo) Label() string {
	return fmt.Sprintf("%s://%s:%d", si.Scheme, si.IP, si.Port)
}

// Server ...
type Server interface {
	Serve() error
	Stop() error
	GracefulStop(ctx context.Context) error
	Info(group, cluster string) *ServiceInfo
}
