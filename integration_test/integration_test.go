/*
Sniperkit-Bot
- Status: analyzed
*/

// Copyright (c) 2018 Palantir Technologies Inc. All rights reserved.
// Use of this source code is governed by the Apache License, Version 2.0
// that can be found in the LICENSE file.

package integration_test

import (
	"testing"

	"github.com/palantir/godel/framework/pluginapitester"
	"github.com/palantir/godel/pkg/products/v2/products"
	"github.com/stretchr/testify/require"
)

func TestUpgradeConfig(t *testing.T) {
	pluginPath, err := products.Bin("amalgomate-plugin")
	require.NoError(t, err)
	pluginProvider := pluginapitester.NewPluginProvider(pluginPath)

	pluginapitester.RunUpgradeConfigTest(t,
		pluginProvider,
		nil,
		[]pluginapitester.UpgradeConfigTestCase{
			{
				Name: "legacy amalgomate config is upgraded",
				ConfigFiles: map[string]string{
					"godel/config/amalgomate.yml": `
amalgomators:
  test-product:
    config: test.yml
    output-dir: test-output
    pkg: test-pkg
  next-product:
    config: next.yml
    output-dir: next-output
    pkg: next-pkg
  other-product:
    config: other.yml
    output-dir: other-output
    pkg: other-pkg
`,
				},
				Legacy:     true,
				WantOutput: "Upgraded configuration for amalgomate-plugin.yml\n",
				WantFiles: map[string]string{
					"godel/config/amalgomate-plugin.yml": `amalgomators:
  next-product:
    order: 1
    config: next.yml
    output-dir: next-output
    pkg: next-pkg
  other-product:
    order: 2
    config: other.yml
    output-dir: other-output
    pkg: other-pkg
  test-product:
    config: test.yml
    output-dir: test-output
    pkg: test-pkg
`,
				},
			},
			{
				Name: "current config is unmodified",
				ConfigFiles: map[string]string{
					"godel/config/amalgomate-plugin.yml": `amalgomators:
  # comment
  test-product:
    order: 0
    config: test.yml
    output-dir: test-output
    pkg: test-pkg
  next-product:
    order: 1
    config: next.yml
    output-dir: next-output
    pkg: next-pkg
  other-product:
    order: 2
    config: other.yml
    output-dir: other-output
    pkg: other-pkg
`,
				},
				WantOutput: "",
				WantFiles: map[string]string{
					"godel/config/amalgomate-plugin.yml": `amalgomators:
  # comment
  test-product:
    order: 0
    config: test.yml
    output-dir: test-output
    pkg: test-pkg
  next-product:
    order: 1
    config: next.yml
    output-dir: next-output
    pkg: next-pkg
  other-product:
    order: 2
    config: other.yml
    output-dir: other-output
    pkg: other-pkg
`,
				},
			},
		},
	)
}
