/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package cmd

import (
	"github.com/palantir/godel/framework/pluginapi"

	"github.com/sniperkit/snk.fork.palantir-godel-amalgomate-plugin/amalgomateplugin/config"
)

var upgradeConfigCmd = pluginapi.CobraUpgradeConfigCmd(config.UpgradeConfig)

func init() {
	rootCmd.AddCommand(upgradeConfigCmd)
}
