package goworker

//type workerFunc func(string, ...interface{}) error

type WorkerFunc func(*Worker, *Job) error