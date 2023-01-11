package internal

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/AYehia0/soundcloud-dl/pkg/client"
	"github.com/AYehia0/soundcloud-dl/pkg/soundcloud"
	"github.com/PuerkitoBio/goquery"
)

var defaultQuality string = "medium"

func Sc(args []string) {

	url := args[0]
	downloadPath := args[len(args)-1]

	statusCode, body, err := client.Get(url)

	if err != nil {
		log.Fatalf("An Error : %s happended while requesting : %s", err, url)
	}
	if statusCode != http.StatusOK {
		fmt.Println("URL doesn't exist : status not 200.")
		return
	}

	// conver the bytes array into something we can read, as goquery doesn't accept strings
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))

	// check if the url is valid
	if !soundcloud.IsValidUrl(url) {
		fmt.Printf("The Url : %s doesn't return a valid track are you sure the track is publicly accessed ?", url)
		return
	}

	soundData := soundcloud.IsValidTrack(doc)

	if soundData == nil {
		fmt.Printf("The Url : %s doesn't return a valid track are you sure the track is publicly accessed ?", url)
		return
	}
	clientId := soundcloud.GetClientId(doc)

	list := soundcloud.GetFormattedDL(soundData.Transcodings, clientId)

	soundcloud.Download(chooseTrackDownload(list, defaultQuality), downloadPath)
}

// TEMP: Just for now, return the quality
// the default quality is just mp3, highest is ogg
// if the quality doesn't exist return the first one
func chooseTrackDownload(tracks []soundcloud.DownloadTrack, target string) soundcloud.DownloadTrack {
	for _, track := range tracks {
		if track.Quality == target {
			return track
		}
	}
	return tracks[0]
}
