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
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// The following environment variables are required for this script
const (
	URL    = "URL"
	COMMIT = "COMMIT"
	REF    = "REF"
)

// DeploymentManifest defines the JSON schema for a deployment manifest file
type DeploymentManifest struct {
	Commit string `json:"commit"`
	Ref    string `json:"ref"`
}

func main() {
	c, err := configFromEnv()
	if err != nil {
		os.Stderr.WriteString("Configuration issue: ")
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}

	err = verify(c.URL, c.Commit, c.Ref)
	if err != nil {
		os.Stderr.WriteString("Issue verifying deployment: ")
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}

	fmt.Println("Deployment verified")
	os.Exit(0)
}

func verify(url, commit, ref string) error {
	res, err := http.Get(url)
	if err != nil {
		return errors.Wrap(err, "Error fetching deployment manifest JSON")
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return errors.Errorf("Received non-200 HTTP status code: %s", res.Status)
	}

	defer res.Body.Close()
	var body DeploymentManifest
	json.NewDecoder(res.Body).Decode(&body)

	var result *multierror.Error
	if commit != "" && !strings.HasPrefix(body.Commit, commit) {
		err := errors.Errorf("Commit %s does not match (expected value: %s)", body.Commit, commit)
		result = multierror.Append(result, err)
	}

	if ref != "" && ref != body.Ref {
		err := errors.Errorf("Ref %s does not match (expected value: %s)", body.Ref, ref)
		result = multierror.Append(result, err)
	}

	return result.ErrorOrNil()
}

type config struct {
	URL    string
	Commit string
	Ref    string
}

func configFromEnv() (config, error) {
	var result *multierror.Error

	URL, ok := os.LookupEnv(URL)
	if !ok {
		err := errors.Errorf("The environment variable %s must be set to the URL of the deployment manifest.", URL)
		result = multierror.Append(result, err)
	}

	c := config{
		URL:    URL,
		Commit: os.Getenv(COMMIT),
		Ref:    os.Getenv(REF),
	}

	if c.Commit == "" && c.Ref == "" {
		err := errors.Errorf("The environment variables %s and %s are not set. At least one of these must be set.", COMMIT, REF)
		result = multierror.Append(result, err)
	}

	return c, result.ErrorOrNil()
}
