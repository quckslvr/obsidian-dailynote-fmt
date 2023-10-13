package main

import "github.com/quckslvr/obsfmt/pkg/cmd"

func main() {
	cmd.LoadConfig("examples/tracking_example")
	cmd.Scanner()
	cmd.DryRun()
	cmd.Exec()
	cmd.Summary()
}
