package console

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cobra-example",
	Short: "An example of cobra",
	Long: `This application shows how to create modern CLI
			applications in go using Cobra CLI library`,
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
