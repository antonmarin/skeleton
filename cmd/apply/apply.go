package apply

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply skeleton to project",
	Long: `Apply skeleton to project files.
Be careful, it would replace previous skeleton blocks.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		prompt := promptui.Select{
			Label: "Select",
			Items: []string{"Docker", "Vagrant"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return err
		}

		fmt.Printf("You choose %q\n", result)
		return nil
	},
}
