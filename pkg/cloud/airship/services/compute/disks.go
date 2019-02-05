/*
Copyright 2018 The Kubernetes Authors.

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

package compute

import (
	"github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/services/compute"
)

// DeleteManagedDisk deletes a managed disk resource.
func (s *Service) DeleteManagedDisk(resourceGroup string, name string) (compute.DisksDeleteFuture, error) {
	return s.scope.AirshipClients.Disks.Delete(s.scope.Context, resourceGroup, name)
}

// WaitForDisksDeleteFuture returns when the DeleteManagedDisk operation completes.
func (s *Service) WaitForDisksDeleteFuture(future compute.DisksDeleteFuture) error {
	return future.Future.WaitForCompletionRef(s.scope.Context, s.scope.AirshipClients.Disks.Client)
}
