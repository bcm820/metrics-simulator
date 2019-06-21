package metrics

import (
	"time"
)

type Metrics struct {
	StartTime int64 `json:"system/start_time"`
}

func New() *Metrics {
	const nsPerMs = 1000000
	return &Metrics{
		StartTime: time.Duration(time.Now().UnixNano()).Nanoseconds() / nsPerMs,
	}
}

// {
//   "grey-matter-metrics-version": "1.0.0",
//   "Total/requests": 0,
//   "HTTP/requests": 0,
//   "HTTPS/requests": 0,
//   "RPC/requests": 0,
//   "RPC_TLS/requests": 0,
//   "all/requests": 0,
//   "all/routes": "",
//   "all/latency_ms.avg": 0,
//   "all/latency_ms.count": 0,
//   "all/latency_ms.max": 0,
//   "all/latency_ms.min": 0,
//   "all/latency_ms.sum": 0,
//   "all/latency_ms.p50": 0,
//   "all/latency_ms.p90": 0,
//   "all/latency_ms.p95": 0,
//   "all/latency_ms.p99": 0,
//   "all/latency_ms.p9990": 0,
//   "all/latency_ms.p9999": 0,
//   "all/errors.count": 0,
//   "all/in_throughput": 0,
//   "all/out_throughput": 0,
//   "go_metrics/runtime/num_goroutines": 14,
//   "system/start_time": 1561028736513,
//   "system/cpu.pct": 1.551724,
//   "system/cpu_cores": 6,
//   "os": "linux",
//   "os_arch": "amd64",
//   "system/memory/available": 7132164096,
//   "system/memory/used": 883466240,
//   "system/memory/used_percent": 10.567284,
//   "process/memory/used": 72546552
// }
