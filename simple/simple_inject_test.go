package simple

import (
	"fmt"
	"testing"

	"github.com/habibiiberahim/go-backend/simple"
)

func TestSimpleInject(t *testing.T) {
	simpleService := simple.InitializedService
	fmt.Println(simpleService)
}
