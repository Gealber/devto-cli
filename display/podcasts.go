package display

import (
	"fmt"

	"github.com/Gealber/devto-cli/api"
)

func PodcastResponse(podcasts *api.PodcastResponse) {
	if len(*podcasts) == 0 {
		fmt.Println(toBoldRed("Sorry no podcasts"))
	}
	fmt.Println()
	header()
	fmt.Println()
	for i, podcast := range *podcasts {
		fmt.Println()
		id := fmt.Sprintf("%d. #%d:", i+1, podcast.ID)
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow(podcast.Title))
		fmt.Println()
	}
}
