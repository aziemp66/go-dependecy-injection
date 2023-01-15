package simple

type SimpleRepository struct {
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

type SimpleService struct {
	SimpleRepository SimpleRepository
}

func NewSimpleService(simpleRepository SimpleRepository) *SimpleService {
	return &SimpleService{SimpleRepository: simpleRepository}
}
