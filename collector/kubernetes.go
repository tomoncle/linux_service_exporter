/*
Copyright 2019 tomoncle.

Licensed under the GNU General Public License, Version 3 (the "License")
*/

package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"github.com/tomoncle/linux_service_exporter/tools"
	"strconv"
)

const (
	namespace = "service"
	subsystem = "kube"
	//labels       =  []string{"host", "service"}
)

var (
	ipAddress, _ = tools.GetInterface()

	// 构造 kube-apiserver 服务
	kubeApiserver = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, "apiserver"),
		"kubernetes service of kube-apiserver.",
		[]string{"host", "service"}, nil,
	)

	// 构造 kube-controller-manager 服务
	kubeControllerManager = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, "controller_manager"),
		"kubernetes service of kube-controller-manager.",
		[]string{"host", "service"}, nil,
	)

	// 构造 kube-scheduler 服务
	kubeScheduler = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, "scheduler"),
		"kubernetes service of kube-scheduler.",
		[]string{"host", "service"}, nil,
	)

	// 构造 kube-proxy 服务
	kubeProxy = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, "proxy"),
		"kubernetes service of kube-proxy.",
		[]string{"host", "service"}, nil,
	)

	// 构造 kubelet 服务
	kubelet = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, subsystem, "kubelet"),
		"kubernetes service of kubelet.",
		[]string{"host", "service"}, nil,
	)
)

// serviceInfo returns "map" format data
func serviceInfo(serviceName string) map[string]string {
	status, _ := tools.CentOSServiceActive(serviceName)
	return map[string]string{
		"host": ipAddress, "service": serviceName, "value": fmt.Sprint(status)}
}

// dataSource returns kubernetes service status.
func dataSource() map[string]interface{} {
	return map[string]interface{}{
		"kube-apiserver":          serviceInfo("kube-apiserver"),
		"kube-controller-manager": serviceInfo("kube-controller-manager"),
		"kube-scheduler":          serviceInfo("kube-scheduler"),
		"kube-proxy":              serviceInfo("kube-proxy"),
		"kubelet":                 serviceInfo("kubelet"),
	}
}

// Describe returns kubernetes services of the collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	log.Info("call Exporter.Describe function.")
	ch <- kubeApiserver
	ch <- kubeControllerManager
	ch <- kubeScheduler
	ch <- kubeProxy
	ch <- kubelet
}

// CollectServiceType 根据服务类型收集从配置的服务统计信息
func (e *Exporter) CollectServiceType(ch chan<- prometheus.Metric, desc *prometheus.Desc, serviceMap map[string]string) {
	f, err := strconv.ParseFloat(serviceMap["value"], 64)
	if err != nil {
		log.Error("int 转 float64 异常: ", e, serviceMap["value"])
		return
	}
	// 更新ch对象中的prometheus.Desc类型数据
	ch <- prometheus.MustNewConstMetric(
		desc, prometheus.GaugeValue, f, serviceMap["host"], serviceMap["service"],
	)
}

// Collect 实现了prometheus.collector
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	// 调用数据接口，返回动态数据
	services := e.client.data(dataSource())
	for k := range services {
		serviceMap := services[k].(map[string]string)
		if k == "kube-apiserver" {
			e.CollectServiceType(ch, kubeApiserver, serviceMap)
		} else if k == "kube-controller-manager" {
			e.CollectServiceType(ch, kubeControllerManager, serviceMap)
		} else if k == "kube-scheduler" {
			e.CollectServiceType(ch, kubeScheduler, serviceMap)
		} else if k == "kube-proxy" {
			e.CollectServiceType(ch, kubeProxy, serviceMap)
		} else if k == "kubelet" {
			e.CollectServiceType(ch, kubelet, serviceMap)
		} else {
			log.Error("未知的服务.", k)
		}
	}
}
