package somelib

import (
	"fmt"

	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

func DoIt() string {
	a := serviceadapter.Errand{}
	fmt.Println(a)
	return "Done it"
}

func AnotherFn() string {
	return "Bar"
}

func EvenMore() string {
	return "Sha"
}

func Another() string {
	return "??"
}
