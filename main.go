package main

import (
	cmd "dalayqueue/cmd"
)

func main() {
	cmd := new(cmd.Cmd)
	cmd.Run()
}
