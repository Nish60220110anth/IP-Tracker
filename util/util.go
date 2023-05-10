package util

import (
	"github.com/fatih/color"
)

func GenHeader(head string) string {
	return color.New(color.FgHiBlue).Add(color.Underline).Add(color.Bold).Sprintf("%-15s", head)
}
