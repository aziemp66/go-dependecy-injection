package test

import (
	"fmt"
	"testing"

	"github.com/aziemp66/go-dependecy-injection/simple"
)

func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializeService()
	fmt.Println(err)
	fmt.Println(simpleService)
}
