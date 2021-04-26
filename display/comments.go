package display

import (
	"fmt"

	"github.com/Gealber/devto-cli/api"
)

func CommentResponse(comments *api.CommentsResponse) {
	if len(*comments) == 0 {
		fmt.Println(toBoldRed("Sorry no comments"))
	}
	fmt.Println()
	for i, comment := range *comments {
		fmt.Println()
		id := fmt.Sprintf("%d. #%s:", i+1, comment.IDCode)
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow("@"+comment.User.Username))
		fmt.Printf("  %s\n", toItalicCyan("BodyHTML: "+comment.BodyHtml))
		fmt.Println()
	}
}

func CommentType(comment *api.CommentType) {
	fmt.Println()
	id := fmt.Sprintf(" #%s:", comment.IDCode)
	fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow("@"+comment.User.Username))
	fmt.Printf("  %s\n", toItalicCyan("BodyHTML: "+comment.BodyHtml))
	fmt.Println()
}
