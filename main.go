package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Select",
		Items: []string{"Docker", "Vagrant"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}
