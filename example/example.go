package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/narqo/go-badge"
)

var (
	subject    = flag.String("subject", "", "Badge subject")
	status     = flag.String("status", "", "Badge status")
	color      = flag.String("color", "blue", "Badge color")
	badgeColor = flag.String("bcolor", "#bbb", "Badge color")
	labelColor = flag.String("lcolor", "#fff", "Badge color")
)

func main() {
	flag.Parse()
	err := badge.Render(*subject, *status, badge.Color(*badgeColor), badge.Color(*labelColor), badge.Color(*color), os.Stdout)
	if err != nil {
		log.Fatal(err)
	}

	badge, err := badge.RenderBytes(*subject, *status, badge.Color(*badgeColor), badge.Color(*labelColor), badge.Color(*color))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", badge)
}
