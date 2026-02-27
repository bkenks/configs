package main

import "github.com/charmbracelet/huh"

type OS int

const (
	Mac OS = iota
	Ubuntu
)

func main() {

}

type Options struct {
	Action string
	IsDryRun bool
}

func form () *huh.Form {
	options := 

	huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
			Title("Would you like to push or pull files?").
			Options(
				huh.NewOption("Push", "push"),
				huh.NewOption("Pull", "pull"),
			).
			Value()
		),
	)
}

func rsync(os OS, isDryRun bool) {
	
}