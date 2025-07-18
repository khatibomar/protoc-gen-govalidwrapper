package main

import (
	"flag"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var flags flag.FlagSet
	var sourceRelative bool

	flags.BoolVar(&sourceRelative, "source_relative", false, "Generate relative to source files")

	protogen.Options{
		ParamFunc: func(name, value string) error {
			// Handle source_relative parameter
			if name == "source_relative" {
				sourceRelative = true
				return nil
			}
			// Handle other parameters by parsing them
			if strings.Contains(name, "=") {
				parts := strings.SplitN(name, "=", 2)
				return flags.Set(parts[0], parts[1])
			}
			return flags.Set(name, value)
		},
	}.Run(func(plugin *protogen.Plugin) error {
		for _, f := range plugin.Files {
			if !f.Generate {
				continue
			}
			if err := generate(plugin, f); err != nil {
				return err
			}
		}
		return nil
	})
}
