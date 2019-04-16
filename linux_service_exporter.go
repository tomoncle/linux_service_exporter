package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	collect "github.com/tomoncle/linux_service_exporter/collector"
)

func init() {
	log.Info("call init function.")
	prometheus.MustRegister(version.NewCollector(collect.ExporterName))
}

func main() {
	var (
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(":9110").String()
		metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").String()
		serviceName   = kingpin.Flag("service.name", "fake service name params.").Default("linux_service_exporter").String()
		// 初始化数据源参数
		opts = collect.ServiceOpts{}
	)

	log.AddFlags(kingpin.CommandLine)
	kingpin.Version(version.Print(collect.ExporterName))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	log.Infoln("Starting service_exporter", version.Info())
	log.Infoln("Build context", version.BuildContext())
	// 传入配置项参数
	opts.Name = *serviceName
	// 创建 Exporter
	exporter, err := collect.NewExporter(opts)
	if err != nil {
		log.Fatalln(err)
	}
	prometheus.MustRegister(exporter)

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Service Exporter</title></head>
             <body>
             <h1>Service Exporter</h1>
             <p><a href='` + *metricsPath + `'>Metrics</a></p>
             <h2>Build</h2>
             <pre>` + version.Info() + ` ` + version.BuildContext() + `</pre>
             </body>
             </html>`))
	})

	log.Infoln("Listening on", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
