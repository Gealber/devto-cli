package display

import (
	"fmt"

	"github.com/Gealber/devto-cli/api"
)

func FollowTagsResponse(tags *api.FollowTagsResponse) {
	if len(*tags) == 0 {
		fmt.Println(toBoldRed("Sorry no tags"))
	}
	fmt.Println()
	for i, tag := range *tags {
		fmt.Println()
		id := fmt.Sprintf("%d. #%d:", i+1, tag.ID)
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow(tag.Name))
		fmt.Println()
	}
}

func TagsResponse(tags *api.TagsResponse) {
	if len(*tags) == 0 {
		fmt.Println(toBoldRed("Sorry no tags"))
	}
	fmt.Println()
	for i, tag := range *tags {
		fmt.Println()
		id := fmt.Sprintf("%d. #%d:", i+1, tag.ID)
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow(tag.Name))
		fmt.Println()
	}
}
