// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package plugin

import (
	"net/http"

	"github.com/apache/apisix-go-plugin-runner/internal/plugin"
	pkgHTTP "github.com/apache/apisix-go-plugin-runner/pkg/http"
)

// PluginOpts represents the attributes of the Plugin
type PluginOpts struct {
	// Name (required) is the plguin name
	Name string
	// ParseConf (required) is the method to parse given plugin configuration. When the
	// configuration can't be parsed, it will be skipped.
	ParseConf func(in []byte) (conf interface{}, err error)
	// Filter (required) is the method to handle request.
	// It is like the `http.ServeHTTP`, plus the ctx and the configuration created by
	// ParseConf.
	//
	// When the `w` is written, the execution of plugin chain will be stopped.
	// We don't use onion model like Gin/Caddy because we don't serve the whole request lifecycle
	// inside the runner. The plugin is only a filter running at one stage.
	Filter func(conf interface{}, w http.ResponseWriter, r pkgHTTP.Request)
}

// RegisterPlugin register a plugin. Plugin which has the same name can't be registered twice.
func RegisterPlugin(opt *PluginOpts) error {
	return plugin.RegisterPlugin(opt.Name, opt.ParseConf, opt.Filter)
}