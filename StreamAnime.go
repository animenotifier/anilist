package anilist

import (
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/parnurzeal/gorequest"
)

// StreamAnime returns a stream of all anime on anilist.co.
func StreamAnime() chan *Anime {
	channel := make(chan *Anime)
	page := 1
	ticker := time.NewTicker(1100 * time.Millisecond)
	rateLimit := ticker.C

	go func() {
		defer close(channel)
		defer ticker.Stop()

		for {
			animePage := []*Anime{}
			request := gorequest.New().Get("https://anilist.co/api/browse/anime?page=" + strconv.Itoa(page) + "&access_token=" + AccessToken)
			_, _, errs := request.EndStruct(&animePage)

			if len(errs) > 0 {
				color.Red(errs[0].Error())
				page++
				<-rateLimit
				continue
			}

			// We have reached the end
			if len(animePage) == 0 {
				break
			}

			for _, anime := range animePage {
				channel <- anime
			}

			page++
			<-rateLimit
		}
	}()

	return channel
}
