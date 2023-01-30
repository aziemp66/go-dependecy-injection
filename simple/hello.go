package simple

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	Hello SayHello
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name

}

func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}

}

func NewHelloService(h SayHello) *HelloService {
	return &HelloService{Hello: h}

}
