package services

import (
	"educative/adapters/queue"
	"educative/domain"
	"time"
)

type JobWorkerResult struct {
	Job     domain.Job
	Message *queue.Message
	Error   error
}

func Work(messages chan queue.Message, returnChan chan JobWorkerResult, job domain.Job, JobService) {
	for message := range messages {

		// instance a new job
		job := &domain.Job{ID: "1", Status: "STARTING", CreatedAt: time.Now()}

		// persist job on database
		err := jobService.JobRepository.Insert(job)
		if err != nil {
			returnChan <- JobWorkerResult{domain.Job{}, &message, err}
		}

		if err = jobService.Start(); err != nil {
			returnChan <- JobWorkerResult{domain.Job{}, &message, err}
		}
	}
}
