package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("Failed to create SimpleService")
	}

	return &SimpleService{
		SimpleRepository: repository,
	}, nil
}
