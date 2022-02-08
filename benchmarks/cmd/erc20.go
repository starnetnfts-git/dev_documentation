package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var erc20 = &cobra.Command{
	Use:   "erc20",
	Short: "benchmark erc20",
	Long:  `This command benchmarks an erc20 type of token`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("erc20 benchmark called")
	},
}

func init() {
	RootCmd.AddCommand(erc20)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// byeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// byeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
