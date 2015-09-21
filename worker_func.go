package goworker

//type workerFunc func(string, ...interface{}) error

type WorkerOpt struct {
	IsRetry    bool
	FailRecord bool
}

type WorkerFunc func(*Worker, *Job) error

type WrapWorkerFunc struct {
	Worker WorkerFunc
	Opt    WorkerOpt
}
