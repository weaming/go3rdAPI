package wikipedia

import (
	"fmt"

	libhttp "github.com/weaming/golib/http"
)

func GetAPI(language, action, text string, limit int) string {
	//Create search url
	search, _ := libhttp.URLEncoded(text)
	return fmt.Sprintf("https://%v.wikipedia.org/w/api.php?action=%v&search=%v&limit=%v&origin=*&format=json", language, action, search, limit)
}

func Search(text, language string) []string {
	request := GetAPI(language, "opensearch", text, 3)
	return WikipediaAPI(request)
}
