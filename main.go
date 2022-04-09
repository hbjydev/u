package main

import (
	"log"

	"github.com/hbjydev/zetman"
	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/config"
	"github.com/rwxrob/help"
	"github.com/rwxrob/uniq"
	"github.com/rwxrob/y2j"
	"github.com/rwxrob/yq"
)

func main() {
	log.SetFlags(0)

	Z.Aliases = map[string][]string{
		"isosec": {"uniq", "isosec"},
		"uuid":   {"uniq", "uuid"},
		"epoch":  {"uniq", "second"},
		"zet":    {"zetman"},
	}

	Cmd.Run()
}

var Cmd = &Z.Cmd{
	Name:      `u`,
	Summary:   `hbjydev's bonzai command tree`,
	Version:   `v0.0.3`,
	Copyright: `Copyright 2022 Hayden Young`,
	License:   `Apache-2.0`,
	Commands: []*Z.Cmd{
		help.Cmd, config.Cmd, y2j.Cmd, uniq.Cmd, yq.Cmd, zetman.Cmd,
	},
}
