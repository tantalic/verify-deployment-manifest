/*
Copyright 2017 Toolhouse, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	"github.com/toolhouse/deployment-manifest/pkg/deployment"
)

func main() {
	c, err := configFromEnv()
	if err != nil {
		exitWithError("Configuration issue", err)
	}

	manifest, err := deployment.FetchManifest(c.URL)
	if err != nil {
		exitWithError("Error fetching deployment manifest", err)
	}

	err = manifest.Verify(c.Commit, c.Ref)
	if err != nil {
		exitWithError("Issue verifying deployment", err)
	}

	fmt.Println("Deployment verified")
	os.Exit(0)
}

func exitWithError(desc string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s", desc, err.Error())
	os.Exit(1)
}
