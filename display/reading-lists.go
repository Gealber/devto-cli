package display

import (
	"fmt"

	"github.com/Gealber/devto-cli/api"
)

func ReadingListResponse(readingLists *api.ReadingListResponse) {
	if len(*readingLists) == 0 {
		fmt.Println(toBoldRed("Sorry no reading lists"))
	}
	fmt.Println()
	for i, rdList := range *readingLists {
		fmt.Println()
		id := fmt.Sprintf("%d. #%d:", i+1, rdList.ID)
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow("Article Title: "+rdList.Article.Title))
		fmt.Println()
	}
}
