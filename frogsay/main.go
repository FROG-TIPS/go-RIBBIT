package main

import (
	"fmt"

	"github.com/FROG-TIPS/go-RIBBIT"
)

func main() {
	client := ribbit.NewClient()

	tips, _ := client.Croak()
	for _, tip := range tips {
		fmt.Println(tip.Tip)
	}
}
