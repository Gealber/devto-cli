package display

import (
	"fmt"
	"strconv"

	"github.com/Gealber/devto-cli/api"
)

func ListingResponse(listings *api.ListingResponse) {
	fmt.Println()
	header()
	fmt.Println()
	for i, listing := range *listings {
		fmt.Println()
		id := fmt.Sprintf("%d. #%s:", i+1, strconv.FormatInt(listing.ID, 10))
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow(listing.Title))
		fmt.Println()
	}
}

func CreatedListing(listing *api.ListingType) {
	fmt.Println()
	id := fmt.Sprintf("#%s:", strconv.FormatInt(int64(listing.ID), 10))
	fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow(listing.Title))
	fmt.Printf("  %s\n", toItalicGreen("Tags: "+listing.TagList))
	fmt.Printf("  %s\n", toItalicYellow("Username: "+"@"+listing.User.Username))
	fmt.Println()
}
