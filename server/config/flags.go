package config

import (
	"flag"
)

var configDir string

func BindFlags(fs *flag.FlagSet) {
	fs.StringVar(&configDir, "config-dir", "./config/fs", "server configuration directory")
}
