package stack

func GetStacks() []Stack {

	stacks := make([]Stack, 0)

	for _, s := range Registry {
		stacks = append(stacks, s)
	}

	return stacks
}