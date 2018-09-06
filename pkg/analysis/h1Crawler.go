package analysis

import (
	"fmt"
	"net/http"
	"strings"

	log "github.com/llimllib/loglevel"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type H1tag struct {
	Text string `json:"text"`
}

type HttpError struct {
	original string
}

func (self HttpError) Error() string {
	return self.original
}

func tagReader(resp *http.Response) []H1tag {
	page := html.NewTokenizer(resp.Body)
	tags := []H1tag{}

	var start *html.Token
	var text string

	for {
		_ = page.Next()
		token := page.Token()
		if token.Type == html.ErrorToken {
			break
		}

		if start != nil && token.Type == html.TextToken {
			text = fmt.Sprintf("%s%s", text, token.Data)
		}

		if token.DataAtom == atom.H1 {
			switch token.Type {
			case html.StartTagToken:
				start = &token
			case html.EndTagToken:
				if start == nil {
					continue
				}
				tag := H1tag{Text: strings.TrimSpace(text)}
				tags = append(tags, tag)

				start = nil
				text = ""
			}
		}
	}

	return tags
}

func Downloader(url string) []H1tag {
	page, err := http.Get(url)
	if err != nil {
		log.Debugf("Error: %s", err)
	}

	if page.StatusCode > 299 {
		err = HttpError{fmt.Sprintf("Error (%d): %s", page.StatusCode, url)}
		log.Debug(err)
	}

	if err != nil {
		log.Error(err)
	}

	return tagReader(page)
}
