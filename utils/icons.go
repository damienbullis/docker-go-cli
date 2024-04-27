package utils

import "github.com/charmbracelet/log"

type IconMap struct {
	Check  string
	Cross  string
	Rocket string
	Fire   string
	Mag    string
	Tada   string
	Time   string
}

var Icon = &IconMap{
	Check:  "✔",
	Cross:  "✖",
	Rocket: "🚀",
	Fire:   "🔥",
	Mag:    "🔍",
	Tada:   "🎉",
	Time:   "⏰",
}

func supportsEmoji() bool {
	log.Error("Emoji support not implemented")
	return true
}

func init() {
	// check for emoji support
	if !supportsEmoji() {
		log.Debug("Disabling emoji icons")
		Icon.Check = ""
		Icon.Cross = ""
		Icon.Rocket = ""
		Icon.Fire = ""
		Icon.Mag = ""
		Icon.Tada = ""
		Icon.Time = ""
	}
}
