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
	"os"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// The following environment variables are required
const (
	URL    = "URL"
	COMMIT = "COMMIT"
	REF    = "REF"
)

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
