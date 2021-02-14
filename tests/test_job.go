package tests

//NewTestJob ..
func NewTestJob() *TestJob {
	return &TestJob{
		Number: 0,
	}
}

//TestJob ...
type TestJob struct {
	Number int
}

//Execute ...
func (job *TestJob) Execute() {
	job.Number++
}
