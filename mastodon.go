package main

import (
	"log"
	"context"
	"github.com/mattn/go-mastodon"
	"fmt"
)

func main() {
	c := mastodon.NewClient(&mastodon.Config{
		Server:       "https://mstdn.jp",
		ClientID:     "7060beee4e0d040a93c794d038a89349b45a42f9890b26e6585893248a620b82",
		ClientSecret: "10534aa04a831dc6e9bb0eddb38ecc30e1e0c01d8b68b4c2de047b62ddec2216",
	})

	err := c.Authenticate(context.Background(), "jun.hiroe@gmail.com", "rah7AhFahX$4uxae")
	if err != nil {
		log.Fatal(err)
	}
	timeline, err := c.GetTimelineHome(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	for i := len(timeline) - 1; i >= 0; i-- {
		fmt.Println(timeline[i])
	}
}
