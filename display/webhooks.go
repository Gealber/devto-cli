package display

import (
	"fmt"

	"github.com/Gealber/devto-cli/api"
)

func WebhooksResponse(webhooks *api.WebhooksResponse) {
	if len(*webhooks) == 0 {
		fmt.Println(toBoldRed("Sorry no webhooks"))
	}
	fmt.Println()
	for i, webhook := range *webhooks {
		fmt.Println()
		id := fmt.Sprintf("%d. #%d:", i+1, webhook.ID)
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow("Target URL: "+webhook.TargetURL))
		fmt.Println()
	}
}

func WebhooksCreated(webhook *api.WebhookCreatedResponse) {
	fmt.Println()
	id := fmt.Sprintf("  #%d:", webhook.ID)
	fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow("Username: @"+webhook.User.Username))
	fmt.Printf("  %s\n", toItalicYellow("Target URL: "+webhook.TargetURL))
	fmt.Println()
}

func WebhooksTypeBasic(webhook *api.WebhookTypeBasic) {
	fmt.Println()
	id := fmt.Sprintf("  #%d:", webhook.ID)
	fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow("Username: @"+webhook.User.Username))
	fmt.Printf("  %s\n", toItalicYellow("Target URL: "+webhook.TargetURL))
	fmt.Println()
}
