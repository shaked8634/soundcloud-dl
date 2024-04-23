package soundclouddl

import "os"

var Search bool
var DownloadPath string
var Quality string

// define flags and handle configuration
func InitConfigVars() {
	tmpDLdir, _ := os.Getwd()
	rootCmd.PersistentFlags().BoolVarP(&Search, "search-and-download", "s", false, "Search for tracks by title and prompt one for download ")
	rootCmd.PersistentFlags().StringVarP(&DownloadPath, "download-path", "p", tmpDLdir, "The download path where tracks are stored.")
	rootCmd.PersistentFlags().StringVarP(&Quality, "quality", "q", "", "Track audio quality to to download (low, medium or high).")
}
