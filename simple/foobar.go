package simple

type FoobarService struct {
	*FooService
	*BarService
}

func NewFoobarService(fooService *FooService, barService *BarService) *FoobarService {
	return &FoobarService{
		FooService: fooService,
		BarService: barService,
	}
}
