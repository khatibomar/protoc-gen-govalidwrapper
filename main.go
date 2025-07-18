package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
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
