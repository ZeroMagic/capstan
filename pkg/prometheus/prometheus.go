/*
Copyright (c) 2018 The ZJU-SEL Authors.

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

package prometheus

import (
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes"
)

// Config is the internal representation of prometheus configuration.
type Config struct {
	PushgatewayEndpoint string `json:"PushgatewayEndpoint"`
}

// Start is the entry to start prometheus.
func Start(kubeClient kubernetes.Interface, cfg Config) error {
	return errors.Errorf("Not Implemented")
}