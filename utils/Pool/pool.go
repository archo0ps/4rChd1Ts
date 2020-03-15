package Pool

type Task struct {
	f    func(url string) error
	url  string
}

func NewTask(task func(string)  error, url string) *Task {
	return &Task{
		f:    task,
		url:  url,
	}
}

func (t *Task) Execute() {
	_ = t.f(t.url)
}

type Pool struct {
	EntryChannel chan *Task
	JobsChannel  chan *Task
	MaxWorker    int
}

func NewPool(cap int) *Pool {
	return &Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		MaxWorker:    cap,
	}
}

func (p *Pool) worker(WorkerId int) {
	for task := range p.JobsChannel {
		task.Execute()
	}
}

func (p *Pool) Run() {
	for i := 0; i < p.MaxWorker; i++ {
		go p.worker(i)
	}

	for task := range p.EntryChannel {
		p.JobsChannel <- task
	}
}
