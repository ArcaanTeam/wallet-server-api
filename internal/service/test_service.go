package service

import "context"

type ITestRepo interface {
	Ping(ctx context.Context) string
}

type TestService struct {
	repo ITestRepo
}

func NewTestService(repo ITestRepo) *TestService {
	return &TestService{
		repo: repo,
	}
}

func (service *TestService) Ping(ctx context.Context) string {
	return service.repo.Ping(ctx)
}
