package main

import (
	"flag"

	"github.com/ksysoev/grpc-redis/pkg/gen"
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

			err := gen.Generate(plugin, f)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
