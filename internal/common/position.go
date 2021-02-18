package common

import (
	"github.com/spf13/cobra"
	"os"
)

func GetCurrentPosition() string{
	position, err := os.Getwd()
	cobra.CheckErr(err)
	return position
}
