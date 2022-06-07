package services

import (
	"educative/domain"
	"educative/repositories"
	"time"
)

type JobService struct {
	Job           *domain.Job
	JobRepository repositories.JobRepository
}

func (svc *JobService) Start() error {
	svc.Job.Status = "PROGRESS"
	if err := svc.JobRepository.Update(svc.Job); err != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)

	svc.Job.Status = "COMPLETED"
	if err := svc.JobRepository.Update(svc.Job); err != nil {
		return err
	}

	time.Sleep(500 * time.Millisecond)

	return nil
}
