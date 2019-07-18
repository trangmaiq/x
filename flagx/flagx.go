package flagx

import (
	"github.com/spf13/cobra"
	"github.com/trangmaiq/x/cmdx"
)

// MustGetBool returns a bool flag or call cmdx.Fatalf if an error occurs
func MustGetBool(cmd *cobra.Command, name string) bool {
	val, err := cmd.Flags().GetBool(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return val
}

// MustGetString returns a string flag or call cmdx.Fatalf if an error occurs
func MustGetString(cmd *cobra.Command, name string) string {
	s, err := cmd.Flags().GetString(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return s
}

// MustGetStringSlice returns a []string flag or call cmdx.Fatalf id an error occurs
func MustGetStringSlice(cmd *cobra.Command, name string) []string {
	ss, err := cmd.Flags().GetStringSlice(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}

	return ss
}

func MustGetInt(cmd *cobra.Command, name string) int {
	i, err := cmd.Flags().GetInt(name)
	if err != nil {
		cmdx.Fatalf(err.Error())
	}
	return i
}
