package display

import "fmt"

func toItalicYellow(text string) string {
	return fmt.Sprintf("\033[3;33m%s\033[0m", text)
}

func toBoldGreen(text string) string {
	return fmt.Sprintf("\033[1;32m%s\033[0m", text)
}

func toBoldRed(text string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", text)
}

func toItalicGreen(text string) string {
	return fmt.Sprintf("\033[3;32m%s\033[0m", text)
}

func toItalicCyan(text string) string {
	return fmt.Sprintf("\033[3;36m%s\033[0m", text)
}

func header() {
	fmt.Printf("%s %s\n  %s", toBoldGreen("  #<id>  "), toItalicYellow("<title>"), toItalicCyan("URL: <url>"))
}
