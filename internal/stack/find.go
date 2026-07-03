package stack

import "fmt"

func Find(name string) (Stack, error) {

	stack, ok := Registry[name]

	if !ok {
		return Stack{}, fmt.Errorf("stack not found")
	}

	return stack, nil
}