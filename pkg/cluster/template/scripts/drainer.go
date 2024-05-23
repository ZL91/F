// Copyright 2020 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package scripts

import (
	"bytes"
	"path"
	"text/template"

	"github.com/pingcap/tiup/embed"
	"github.com/pingcap/tiup/pkg/utils"
)

// DrainerScript represent the data to generate drainer config
type DrainerScript struct {
	NodeID string
	Addr   string
	PD     string

	DeployDir string
	DataDir   string
	LogDir    string

	NumaNode  string
	IP        string
	Port      int
	Endpoints []*PDScript
}

// ConfigToFile write config content to specific file.
func (c *DrainerScript) ConfigToFile(file string) error {
	fp := path.Join("templates", "scripts", "run_drainer.sh.tpl")
	tpl, err := embed.ReadTemplate(fp)
	if err != nil {
		return err
	}
	tmpl, err := template.New("Drainer").Parse(string(tpl))
	if err != nil {
		return err
	}

	content := bytes.NewBufferString("")
	if err := tmpl.Execute(content, c); err != nil {
		return err
	}

	return utils.WriteFile(file, content.Bytes(), 0755)
}