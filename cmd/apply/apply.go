package apply

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply skeleton to project",
	Long: `Apply skeleton to project files.
Be careful, it would replace previous skeleton blocks.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

type applyUseCase interface {
	Execute() error
}
