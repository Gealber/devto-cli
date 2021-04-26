package display

import (
	"fmt"

	"github.com/Gealber/devto-cli/api"
)

func ProfileImageResponse(profileImg *api.ProfileImageResponse) {
	fmt.Println()
	imageOf := fmt.Sprintf("ImageOf %s:", profileImg.ImageOf)
	fmt.Printf("  %s\n", toBoldGreen(imageOf))
	fmt.Println()
}
