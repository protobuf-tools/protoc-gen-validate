// Copyright 2021 The protobuf-tools Authors
// SPDX-License-Identifier: BSD-3-Clause

// Command protoc-gen-validate is the protoc plugin to generate Google AIP validators.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/compiler/protogen"

	"github.com/protobuf-tools/protoc-gen-validate/internal/version"
	"github.com/protobuf-tools/protoc-gen-validate/validate"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "-version" || os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%s: %s\n", filepath.Base(os.Args[0]), version.String())
		os.Exit(0)
	}

	var flags flag.FlagSet

	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		validate.Gen(gen)

		gen.SupportedFeatures = validate.SupportedFeatures

		return nil
	})
}
