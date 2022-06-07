package repositories

import "educative/domain"

type JobRepository interface {
	Insert(job *domain.Job) error
	Update(job *domain.Job) error
}

type InMemoryJobRepository struct {
	Jobs []domain.Job
}

func (repo *InMemoryJobRepository) Insert(job *domain.Job) error {
	repo.Jobs = append(repo.Jobs, *job)
	return nil
}

func (repo *InMemoryJobRepository) Update(job *domain.Job) error {
	repo.Jobs = append(repo.Jobs, *job)
	
	return nil
}
