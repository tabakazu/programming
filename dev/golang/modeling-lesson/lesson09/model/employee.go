package model

// 社員
type Employee struct {
	ID    int    // ID
	JobID int    // 職種 ID
	Name  string // 名称
	Sex   int    // 性別
	// BirthedAt // 誕生日
	// HiredAt // 入社日
}

func NewEmployee(id int, job_id int, name string, sex int) *Employee {
	return &Employee{
		ID:    id,
		JobID: job_id,
		Name:  name,
		Sex:   sex,
	}
}
