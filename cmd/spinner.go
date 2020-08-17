package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
)

var (
	spinnerSet       = spinner.CharSets[26]
	spinnerEndSucces = "✓"
	spinnerEndFail   = "X"
)

func runSpinnerTask(name string, task func() error) error {
	s := spinner.New(spinnerSet, 600*time.Millisecond)
	s.Prefix = name + " "
	s.Start()
	err := task()
	spinnerEnd := spinnerEndSucces
	if err != nil {
		spinnerEnd = spinnerEndFail
	}
	s.FinalMSG = fmt.Sprintf("%s %s\n", name, spinnerEnd)
	time.Sleep(200 * time.Millisecond)
	s.Stop()
	return err
}
