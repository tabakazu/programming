package main_test

import (
	"testing"

	"github.com/tabakazu/modeling-lesson/lesson09/model"
)

type jobTest struct {
	ID   int
	Name string
}

type employeeTest struct {
	ID    int
	JobID int
	Name  string
	Sex   int
}

var jobTests = []jobTest{
	{ID: 1, Name: "技術職"},
	{ID: 2, Name: "営業職"},
}

var employeeTests = []employeeTest{
	{ID: 1, JobID: 1, Name: "田中花子", Sex: 1},
	{ID: 2, JobID: 1, Name: "山田太郎", Sex: 2},
}

func TestMain(t *testing.T) {
	for i := range jobTests {
		test := &jobTests[i]
		job := model.NewJob(test.ID, test.Name)

		if job.ID != test.ID || job.Name != test.Name {
			t.Errorf("job.ID, job.Name = %v, %v want %v, %v", job.ID, job.Name, test.ID, test.Name)
		}
	}

	for i := range employeeTests {
		test := &employeeTests[i]
		employee := model.NewEmployee(test.ID, test.JobID, test.Name, test.Sex)

		if employee.ID != test.ID || employee.JobID != test.JobID ||
			employee.Name != test.Name || employee.Sex != test.Sex {
			t.Errorf("employee.ID, employee.JobID, employee.Name, employee.Sex = %v, %v, %v, %v want %v, %v, %v, %v",
				employee.ID, employee.JobID, employee.Name, employee.Sex, test.ID, test.JobID, test.Name, test.Sex)
		}
	}
}
