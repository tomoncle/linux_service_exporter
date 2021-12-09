/*
Copyright 2019 tomoncle.

Licensed under the GNU General Public License, Version 3 (the "License")
*/

package collector

const (
	// ExporterName define this service name
	ExporterName = "service_exporter"
)

// ServiceClient 描述 service的客户端，提供服务
type ServiceClient struct {
	ServiceName string
}

// 提供数据接口
func (sc *ServiceClient) data(dataSource map[string]interface{}) map[string]interface{} {
	return dataSource
}

// Exporter 实现了prometheus.Collector interface, 指标出口
type Exporter struct {
	// 参数为服务客户端ServiceClient的结构体指针，来调用client的数据接口
	client *ServiceClient
}

// ServiceOpts 表示ServiceClient配置项，初始化Exporter时使用
type ServiceOpts struct {
	Name string
}

// NewExporter 返回一个初始化的 Exporter.
func NewExporter(opts ServiceOpts) (*Exporter, error) {
	client := &ServiceClient{ServiceName: opts.Name}
	//log.Info("注册服务名称为：", client.ServiceName)
	// Init our exporter.
	return &Exporter{
		client: client,
	}, nil
}
