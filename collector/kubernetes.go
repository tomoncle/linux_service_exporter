package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	namespace    = "service"
	subsystem    = "kube"
	//labels       =  []string{"host", "service"}
)

var (
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
		prometheus.BuildFQName(namespace, subsystem, "let"),
		"kubernetes service of kubelet.",
		[]string{"host", "service"}, nil,
	)

)

// 构造数据源
func dataSource() map[string]interface{} {
	log.Info("调用数据源...")
	servicesA := []map[string]string{
		{"service_id": "a1", "service_name": "serviceA1", "value": fmt.Sprint(rand.Float64())},
		{"service_id": "a2", "service_name": "serviceA2", "value": fmt.Sprint(rand.Float64())},
	}
	servicesB := []map[string]string{
		{"service_id": "b1", "service_name": "serviceB1", "value": fmt.Sprint(rand.Float64())},
		{"service_id": "b2", "service_name": "serviceB2", "value": fmt.Sprint(rand.Float64())},
	}
	servicesC := []map[string]string{
		{"service_id": "c1", "service_name": "serviceC1", "value": fmt.Sprint(rand.Float64())},
		{"service_id": "c2", "service_name": "serviceC2", "value": fmt.Sprint(rand.Float64())},
	}
	servicesD := []map[string]string{
		{"service_id": "d1", "service_name": "serviceD1", "value": fmt.Sprint(rand.Float64())},
		{"service_id": "d2", "service_name": "serviceD2", "value": fmt.Sprint(rand.Float64())},
	}

	return map[string]interface{}{"a": servicesA, "b": servicesB, "c": servicesC, "d": servicesD,}
}

// 描述服务导出的所有度量指标
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	log.Info("call Exporter.Describe function.")
	ch <- kubeApiserver
	ch <- kubeControllerManager
	ch <- kubeScheduler
	ch <- kubeProxy
	ch <- kubelet
}

// 根据服务类型收集从配置的服务统计信息
func (e *Exporter) CollectServiceType(ch chan<- prometheus.Metric, desc *prometheus.Desc, serviceMap map[string]string) {
	f, err := strconv.ParseFloat(serviceMap["value"], 64)
	if err != nil {
		log.Error("int 转 float64 异常: ", e, serviceMap["value"])
		return
	}
	// 更新ch对象中的prometheus.Desc类型数据
	ch <- prometheus.MustNewConstMetric(
		desc, prometheus.GaugeValue, f, serviceMap["service_id"], serviceMap["service_name"],
	)
}

// 实现了prometheus.collector
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	log.Info("call Exporter.Collect function.")
	// 调用数据接口，返回动态数据
	services := e.client.data(dataSource())
	for k := range services {
		servicesList := services[k].([]map[string]string)
		if k == "a" {
			for _, serviceMap := range servicesList {
				e.CollectServiceType(ch, kubeApiserver, serviceMap)
			}

		} else if k == "b" {
			for _, serviceMap := range servicesList {
				e.CollectServiceType(ch, kubeControllerManager, serviceMap)
			}

		} else if k == "c" {
			for _, serviceMap := range servicesList {
				e.CollectServiceType(ch, kubeScheduler, serviceMap)
			}

		} else if k == "d" {
			for _, serviceMap := range servicesList {
				e.CollectServiceType(ch, kubeProxy, serviceMap)
			}
		} else {
			log.Error("未知的服务.", k)
		}
	}
}
