package generics

import "fmt"

// Generics allow functions and types to be parameterized by type, enabling code reusability without sacrificing type safety.
// Introduced in Go 1.18, generics help avoid code duplication while keeping type safety intact.

type GenricInterface interface {
	int | float32 | string
}

func Helper() {
	fmt.Println("In Generics Package!!!")
	fmt.Println(Generic([]int{1, 2, 3, 4, 5}))
	fmt.Println(Generic([]float32{1.1, 2.2, 3.3, 4.4, 5.5}))
	fmt.Println(Generic([]string{"1", "2", "3", "4", "5"}))
}

func Generic[T GenricInterface](slice []T) T {
	var result T
	for _, v := range slice {
		result += v
	}
	return result
	// Generics use type parameters inside square brackets [T], where T represents a generic type.
}
