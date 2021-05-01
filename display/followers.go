package display

import (
	"fmt"
	"strconv"

	"github.com/Gealber/devto-cli/api"
)

func FollowersResponse(followers *api.FollowersResponse) {
	fmt.Println()
	if len(*followers) > 0 {
		header()
	} else {
		fmt.Println(toBoldRed("Sorry no followes"))
		return
	}
	fmt.Println()
	for i, follower := range *followers {
		fmt.Println()
		id := fmt.Sprintf("%d. #%s:", i+1, strconv.FormatInt(int64(follower.ID), 10))
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow("@"+follower.Username))
		fmt.Printf("  %s\n", toBoldGreen("Name: "+follower.Name))
		fmt.Println()
	}
}
