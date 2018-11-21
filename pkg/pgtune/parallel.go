package pgtune

import (
	"fmt"
	"math"
)

// Keys in the conf file that are tuned related to parallelism
const (
	MaxWorkerProcessesKey       = "max_worker_processes"
	MaxParallelWorkersGatherKey = "max_parallel_workers_per_gather"
	MaxParallelWorkers          = "max_parallel_workers"

	errOneCPU = "cannot make recommendations with just 1 CPU"
)

// ParallelKeys is an array of keys that are tunable for parallelism
var ParallelKeys = []string{
	MaxWorkerProcessesKey,
	MaxParallelWorkersGatherKey,
	MaxParallelWorkers,
}

// ParallelRecommender gives recommendations for ParallelKeys based on system resources
type ParallelRecommender struct {
	cpus int
}

// NewParallelRecommender returns a ParallelRecommender that recommends based on the given
// number of cpus.
func NewParallelRecommender(cpus int) *ParallelRecommender {
	return &ParallelRecommender{cpus}
}

// Recommend returns the recommended PostgreSQL formatted value for the conf
// file for a given key.
func (r *ParallelRecommender) Recommend(key string) string {
	var val string
	if r.cpus <= 1 {
		panic(errOneCPU)
	}
	if key == MaxWorkerProcessesKey || key == MaxParallelWorkers {
		val = fmt.Sprintf("%d", r.cpus)
	} else if key == MaxParallelWorkersGatherKey {
		val = fmt.Sprintf("%d", int(math.Round(float64(r.cpus)/2.0)))
	} else {
		panic(fmt.Sprintf("unknown key: %s", key))
	}
	return val
}
