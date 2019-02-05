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

package resources

import (
	"github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/autorest"
	//JEB "github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/autorest/to"
	"github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/services/resources"
)

// CreateOrUpdateGroup creates or updates an airship resource group.
func (s *Service) CreateOrUpdateGroup(resourceGroupName string, location string) (resources.Group, error) {
	//JEB return s.scope.AirshipClients.Groups.CreateOrUpdate(s.scope.Context, resourceGroupName, resources.Group{Location: to.StringPtr(location)})
	return resources.Group{}, nil

}

// DeleteGroup deletes an airship resource group.
func (s *Service) DeleteGroup(resourceGroupName string) (resources.GroupsDeleteFuture, error) {
	//JEB return s.scope.AirshipClients.Groups.Delete(s.scope.Context, resourceGroupName)
	return resources.GroupsDeleteFuture{}, nil
}

// CheckGroupExistence checks oif the resource group exists or not.
func (s *Service) CheckGroupExistence(resourceGroupName string) (autorest.Response, error) {
	//JEB return s.scope.AirshipClients.Groups.CheckExistence(s.scope.Context, resourceGroupName)
	return autorest.Response{}, nil
}

// WaitForGroupsDeleteFuture returns when the DeleteGroup operation completes.
func (s *Service) WaitForGroupsDeleteFuture(future resources.GroupsDeleteFuture) error {
	//JEB return future.WaitForCompletionRef(s.scope.Context, s.scope.AirshipClients.Groups.Client)
	return nil
}
