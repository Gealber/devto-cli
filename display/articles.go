package display

import (
	"fmt"
	"strconv"

	"github.com/Gealber/devto-cli/api"
)

func RetrievedArticles(data *api.GetArticlesResponse) {
	fmt.Println()
	header()
	fmt.Println()
	for i, article := range *data {
		fmt.Println()
		id := fmt.Sprintf("%d. #%s:", i+1, strconv.FormatInt(int64(article.ID), 10))
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow(article.Title))
		fmt.Printf("  %s\n", toItalicCyan("URL: "+article.URL))
		fmt.Println()
	}
}

func ModifiedArticle(article *api.ModifiedArticle) {
	fmt.Println()
	id := fmt.Sprintf("#%s:", strconv.FormatInt(int64(article.ID), 10))
	fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow(article.Title))
	fmt.Printf("  %s\n", toItalicCyan("URL: "+article.URL))
	fmt.Printf("  %s\n", toItalicGreen("Tags: "+article.TagList))
	fmt.Printf("  %s\n", toItalicYellow("Username: "+article.User.Username))
	fmt.Println()
}

func ModifiedArticleBody(article *api.ModifiedArticle) {
	fmt.Println()
	fmt.Println(article.BodyMarkdown)
	fmt.Println()
}

func RetrievedArticlesVideos(data *api.ArticlesVideoResponse) {
	fmt.Println()
	header()
	fmt.Println()
	for i, article := range *data {
		fmt.Println()
		id := fmt.Sprintf("%d. #%s:", i+1, strconv.FormatInt(int64(article.ID), 10))
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow(article.Title))
		fmt.Printf("  %s\n", toItalicGreen("Duration: "+article.VideoDurationInMinutes))
		fmt.Printf("  %s\n", toItalicCyan("Source video URL: "+article.VideoSourceURL))
		fmt.Println()
	}
}

func RetrievedMyArticles(data *api.GetArticlesMeResponse) {
	fmt.Println()
	header()
	fmt.Println()
	for i, article := range *data {
		fmt.Println()
		id := fmt.Sprintf("%d. #%s:", i+1, strconv.FormatInt(int64(article.ID), 10))
		fmt.Printf("  %s %s\n", toBoldGreen(id), toItalicYellow(article.Title))
		fmt.Printf("  %s\n", toItalicCyan("URL: "+article.URL))
		fmt.Println()
	}
}
