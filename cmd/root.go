/*
Copyright © 2023 Mobile Technologies Inc. <connect-support@mtigs.com>
All Rights Reserved
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Unquabain/chatgpc/token"
	"github.com/spf13/cobra"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

var contextSize int
var outputLimit int
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chatgpc",
	Short: "An \"AI\" that will put all content creators out of work.",
	Long: `I read in text of any length and gather statistics about the
contents, and then I try to generate new text that matches the statistics.
I'm not aware of any content or language. It's all just numbers.

GPC is what my kids call my dad: Grandpa Charley.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		defer os.Stdin.Close()

		log.SetHandler(cli.New(os.Stderr))
		if debug {
			log.SetLevel(log.DebugLevel)
		} else {
			log.SetLevel(log.InfoLevel)
		}

		stats := token.NewStats()
		if err := stats.Read(contextSize, os.Stdin); err != nil {
			fmt.Fprintf(os.Stderr, `could not read text: %s`, err.Error())
			os.Exit(-1)
		}
		iter := stats.Iterator(contextSize)
		var output int
		for iter.Scan() {
			word := iter.Text()
			output += len(word)
			if outputLimit > 0 && output >= outputLimit {
				break
			}
			fmt.Print(word)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&contextSize, `context`, `c`, 8, `the number of tokens of context to record`)
	rootCmd.Flags().IntVarP(&outputLimit, `output`, `o`, 4096, `the maximum number of bytes to generate. If zero, spits out text until the algorithm fails`)
	rootCmd.Flags().BoolVarP(&debug, `debug`, `g`, false, `turn on debugging output`)

}
