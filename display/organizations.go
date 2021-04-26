package display

import (
	"fmt"

	"github.com/Gealber/devto-cli/api"
)

func OrganizationResponse(organization *api.OrganizationResponse) {
	fmt.Println()
	fmt.Println()
	fmt.Printf("  %s %s\n", toBoldGreen("Name: "+organization.Name), toItalicYellow("@"+organization.Username))
	fmt.Printf("  %s\n", toItalicGreen("Summary: "+organization.Summarry))
	fmt.Printf("  %s\n", toItalicCyan("URL: "+organization.URL))
	fmt.Println()
}

func UserOnOrganizationResponse(users *api.UserOnOrganizationResponse) {
	if len(*users) == 0 {
		fmt.Println(toBoldRed("Sorry no users"))
	}
	fmt.Println()
	for i, user := range *users {
		fmt.Println()
		id := fmt.Sprintf("%d. #%d:", i+1, user.ID)
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow("@"+user.Username))
		fmt.Printf("  %s\n", toItalicGreen("Summary: "+user.Summary))
		fmt.Println()
	}
}
