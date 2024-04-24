package soundclouddl

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/AYehia0/soundcloud-dl/internal"
	"github.com/AYehia0/soundcloud-dl/pkg/theme"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "sc <url>",
	Short: "Sc is a simple CLI application to download soundcloud tracks",
	Long: `A blazingly fast go program to download tracks from soundcloud 
		using just the URL, with some cool features and beautiful UI.
	`,
	Args:    cobra.ArbitraryArgs,
	Version: "v1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		// get the URL
		if len(args) < 1 && !Search {
			if err := cmd.Usage(); err != nil {
				log.Fatal(err)
			}
			return
		}
		// run the core app
		// FIXME: Probably not the best thing to do lol, it's better to just pass it to the function, who cares.

		if len(args) == 0 {
			args = append(args, "")
		}

		wg := sync.WaitGroup{}
		for _, arg := range args {
			wg.Add(1)
			go internal.Sc(arg, DownloadPath, Quality, Search, &wg) //Currently unrestricted may need mex concurrent processes
		}
		wg.Wait()
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		cmd.Flags().Visit(func(f *pflag.Flag) {
			// check if <url> is passed with --search-and-download flag
			if len(args) != 0 {
				if strings.HasPrefix(args[0], "https") && Search {
					fmt.Printf("Can't use/pass a %s with --%s flag\n\n", theme.Green("<url>"), theme.Red(f.Name))
					cmd.Usage()
					os.Exit(0)
				}
			}
		})
	},
}

func Execute() {
	// initialize the arg parser variables
	InitConfigVars()

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Something went wrong : %s\n", err)
	}
}
