// Copyright 2021 The protobuf-tools Authors
// SPDX-License-Identifier: BSD-3-Clause

package version

import (
	"fmt"
	"runtime/debug"
	"strings"
	"text/tabwriter"
	"text/template"
)

// version is the protoc-gen-validate vesion.
var version = "v0.0.0"

var buildInfoTmpl = ` mod	{{ .Main.Path }}		{{ .Main.Version }}	{{ .Main.Sum }}
{{ range .Deps }} dep	{{ .Path }}		{{ .Version }}	{{ .Sum }}{{ if .Replace }}
	=> {{ .Replace.Path }}	{{ .Replace.Version }}	{{ .Replace.Sum }}{{ end }}
{{ end }}`

func moduleBuildInfo() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "not built in module mode"
	}

	var sb strings.Builder
	tw := tabwriter.NewWriter(&sb, 0, 8, 0, '\t', tabwriter.TabIndent)
	if err := template.Must(template.New("buildinfo").Parse(buildInfoTmpl)).Execute(tw, info); err != nil {
		panic(err)
	}
	if err := tw.Flush(); err != nil {
		panic(err)
	}

	return sb.String()
}

// String returns the protoc-gen-validate current version and build informations.
func String() string {
	return fmt.Sprintf("%s\n\nBuildInfo:\n%s", version, moduleBuildInfo())
}
