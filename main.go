package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {

	cmd := exec.Command("cat")
	r := strings.NewReader("dog1\ndog2")

	cmd.Stdin = r
	data, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))

}
