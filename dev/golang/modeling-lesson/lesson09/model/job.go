package model

// 職種
type Job struct {
	ID   int    // ID
	Name string // 名称
}

func NewJob(id int, name string) *Job {
	return &Job{
		ID:   id,
		Name: name,
	}
}
