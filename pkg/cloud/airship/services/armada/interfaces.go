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

package armada

import "github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/services/armada"

// DeleteNetworkInterface deletes the NIC resource.
func (s *Service) DeleteNetworkInterface(resourceGroup string, networkInterfaceName string) (armada.InterfacesDeleteFuture, error) {
	//JEB return s.scope.AirshipClients.Interfaces.Delete(s.scope.Context, resourceGroup, networkInterfaceName)
	return armada.InterfacesDeleteFuture{}, nil
}

// WaitForNetworkInterfacesDeleteFuture waits for the DeleteNetworkInterface operation to complete.
func (s *Service) WaitForNetworkInterfacesDeleteFuture(future armada.InterfacesDeleteFuture) error {
	//JEB return future.Future.WaitForCompletionRef(s.scope.Context, s.scope.AirshipClients.Interfaces.Client)
	return nil
}
