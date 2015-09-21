package goworker

var (
	workers map[string]WrapWorkerFunc
)

func init() {
	workers = make(map[string]WrapWorkerFunc)
}

// Register registers a goworker worker function. Class
// refers to the Ruby name of the class which enqueues the
// job. Worker is a function which accepts a queue and an
// arbitrary array of interfaces as arguments.
func Register(class string, worker WorkerFunc, opt WorkerOpt) {
	wrap := WrapWorkerFunc{Worker: worker, Opt: opt}
	workers[class] = wrap
}
