package apply

import (
	"errors"
	"github.com/antonmarin/skeleton/config"
	"github.com/antonmarin/skeleton/plugins/dummy"
	"github.com/antonmarin/skeleton/useCases/apply"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "apply",
	Short: "accept skeleton to project",
	Long: `accept skeleton to project files.
Be careful, it would replace previous skeleton blocks.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		factory := config.NewFactory()
		if factory == nil {
			return errors.New("error constructing config factory")
		}
		useCase := apply.NewApply(factory, dummy.NewPlugin())
		err := useCase.Execute()
		if err != nil {
			return err
		}
		return nil
	},
}

//goland:noinspection GoUnusedType
type applyUseCase interface {
	Execute() error
}
