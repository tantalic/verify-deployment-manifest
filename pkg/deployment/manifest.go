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

package deployment

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

// Manifest defines the JSON schema for a deployment manifest file
type Manifest struct {
	Commit string `json:"commit"`
	Ref    string `json:"ref"`
}

func (m Manifest) Verify(commit, ref string) error {
	var result *multierror.Error
	if commit != "" && !strings.HasPrefix(m.Commit, commit) {
		err := errors.Errorf("Commit %s does not match (expected value: %s)", m.Commit, commit)
		result = multierror.Append(result, err)
	}

	if ref != "" && ref != m.Ref {
		err := errors.Errorf("Ref %s does not match (expected value: %s)", m.Ref, ref)
		result = multierror.Append(result, err)
	}

	return result.ErrorOrNil()
}

func FetchManifest(url string) (Manifest, error) {
	var manifest Manifest

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return manifest, err
	}

	res, err := client.Do(req)
	if err != nil {
		return manifest, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return manifest, errors.Errorf("Received non-200 HTTP status code: %s", res.Status)
	}

	err = json.NewDecoder(res.Body).Decode(&manifest)
	res.Body.Close()
	return manifest, err
}
