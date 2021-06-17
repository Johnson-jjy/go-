package main

import "fmt"

type Filter interface {
	About() string
	Process([]int) []int
}

// A UniqueFilter will remove duplicated numbers
type UniqueFilter struct {}

func (UniqueFilter) About() string {
	return "remove duplicated numbers"
}

func (UniqueFilter) Process(inputs []int) []int {
	var outs = make([]int, 0, len(inputs))
	var pushed = make(map[int]bool)
	for _, n := range inputs {
		if !pushed[n] {
			pushed[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}


// A MultipleFilter will only keep numbers which are
// multiples of the MultipleFilter as an integer
type MultipleFilter int

func (mf MultipleFilter) About() string {
	return fmt.Sprintf("keep multiples of %v", mf)
}

func (mf MultipleFilter) Process(inputs []int) []int {
	var outs = make([]int, 0 ,len(inputs))
	for _, n := range inputs {
		if n % int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

// With the help of polymorphism, only one "filterAndPrint"
// function is needed.
func filterAndPrint(fltr Filter, unfiltered []int) []int {
	// Call the methods of "fltr" will call the methods
	// of the value boxed in "fltr".
	filtered := fltr.Process(unfiltered)
	fmt.Println(fltr.About() + ":\n\t", filtered)
	return filtered
}

func testPolymorphism()  {
	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("before filtering:\n\t", numbers)

	// Three non-interface values are boxed into three Filter
	// interface slice element values.
	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(3),
	}

	// Each slice element will be assigned to the local variable
	// "fltr" (of interface type Filter) one by one. The value
	// boxed in each element will also be copied to "fltr".
	for _, fltr := range filters {
		numbers = filterAndPrint(fltr, numbers)
	}
}

func testReflection1()  {
	// Compiler will deduce the type of 123 as int.
	var x interface{} = 123

	// Case 1:
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true
	n = x.(int)
	fmt.Println(n) // 123

	// Case 2:
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false

	// Case 3:
	a = x.(float64) // Will panic.
}

func testReflection2()  {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		// Here, v is declared once, but it denotes
		// different varialbes in different branches.
		switch v := x.(type) {
		case []int: // type literal
			// The type of v is []int.
			fmt.Println("int slice:", v)
		case string: // one type name
			// The type of v is string.
			fmt.Println("string:", v)
		case int, float64, int32: // multiple type names
			// The type of v is always same as x.
			// In this example, it is interface{}.
			fmt.Println("number:", v)
		case nil:
			// The type of v is always same as x.
			// In this example, it is interface{}.
			fmt.Println(v)
		default:
			// The type of v is always same as x.
			// In this example, it is interface{}.
			fmt.Println("others:", v)
		}
		// Note, each variable denoted by v in the
		// last three branches is a copy of x.
	}
}


func main()  {
	//接口的定义

	//接口的初始化（赋值）

	//接口的使用（即为什么需要类型）

	//接口的指针类型和值类型

	//空接口的定义即相关应用

	//类型断言

	testPolymorphism()
}
