package millionRequestfunc

type Dispatcher struct {
	MaxWorkers int
	WorkerPool chan chan Job
}
func NewDispatcher(maxWorker int) *Dispatcher{
	pool:=make(chan chan Job,maxWorker)
	return &Dispatcher{WorkerPool:pool}
}
func (d *Dispatcher)Run()  {
	for i:=0;i<d.MaxWorkers;i++{
		worker:=NewWorker(d.WorkerPool)
		worker.Start()
	}
}
func (d *Dispatcher)dispatch(){
	for{
		select {
		case job := <-JobQueue:
			go func(job Job) {
				jobChannel:= <-d.WorkerPool
				jobChannel <-job
			}(job)
		}
	}
}