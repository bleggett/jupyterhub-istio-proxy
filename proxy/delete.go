/*
Copyright 2020 Splunk Inc.

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

package proxy

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (i *IstioClient) deleteRoute(path string) error {
	name := i.virtualServiceNameWithPrefix(path)
	return i.NetworkingV1alpha3().VirtualServices(i.namespace).Delete(context.Background(), name, v1.DeleteOptions{})
}
