package cmd

import (
	"fmt"

	"github.com/mohsenMj/go-starter-kit/app/types/List"
)

func init() {
	Commands = append(Commands, Command{
		Run:         "test",
		Description: "a general purpose command to run and test part of codes ...",
		Example:     "go run . test",
		Author:      "Mohsen Majidi",
	})
}

func CmdTestHander() {
	fmt.Println("================= Test Start =================")
	fmt.Println("running test")

	ints := []int{5, 10, 150}
	floats := []float32{5.10, 3.10, 75.150, 10.5, 19.33, 58.19}

	fmt.Println("Sum : ", List.Sum(ints))
	fmt.Println("Sum : ", List.Sum(floats))
	fmt.Println("Min Ints : ", List.Min(ints))
	fmt.Println("Min Floats: ", List.Min(floats))
	fmt.Println("Max Ints : ", List.Max(ints))
	fmt.Println("Max Floats: ", List.Max(floats))
	fmt.Println("Before Removing  index 1: ", floats)
	fmt.Println("After  Removing index 1 : ", List.Remove(floats, 1))
	fmt.Println("Push: ", List.Push(floats, 90.3))
	fmt.Println("Pop: ", List.Pop(floats))
	fmt.Println("Sort: ", List.Sort(floats))

	fmt.Println("================= Test End =================")
}
