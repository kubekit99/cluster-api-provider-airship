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

import (
	//JEB "github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/autorest/to"
	"github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/services/armada"
)

const (
	// VnetDefaultName is the default name for the cluster's virtual armada.
	VnetDefaultName = "ClusterAPIVnet"
	// SubnetDefaultName is the default name for the cluster's subnet.
	SubnetDefaultName        = "ClusterAPISubnet"
	defaultPrivateSubnetCIDR = "10.0.0.0/24"
)

// CreateOrUpdateVnet creates or updates a virtual network resource.
func (s *Service) CreateOrUpdateVnet(resourceGroupName string, virtualNetworkName string, location string) (*armada.VirtualNetworksCreateOrUpdateFuture, error) {
	//JEB	if virtualNetworkName == "" {
	//JEB		virtualNetworkName = VnetDefaultName
	//JEB	}
	//JEB
	//JEB	subnets := []armada.Subnet{
	//JEB		{
	//JEB			Name: to.StringPtr(SubnetDefaultName),
	//JEB			SubnetPropertiesFormat: &armada.SubnetPropertiesFormat{
	//JEB				AddressPrefix: to.StringPtr(defaultPrivateSubnetCIDR),
	//JEB			},
	//JEB		},
	//JEB	}
	//JEB	virtualNetworkProperties := armada.VirtualNetworkPropertiesFormat{
	//JEB		AddressSpace: &armada.AddressSpace{
	//JEB			AddressPrefixes: &[]string{defaultPrivateSubnetCIDR},
	//JEB		},
	//JEB		Subnets: &subnets,
	//JEB	}
	//JEB	virtualNetwork := armada.VirtualNetwork{
	//JEB		Location:                       to.StringPtr(location),
	//JEB		VirtualNetworkPropertiesFormat: &virtualNetworkProperties,
	//JEB	}
	//JEB	sgFuture, err := s.scope.AirshipClients.VirtualNetworks.CreateOrUpdate(s.scope.Context, resourceGroupName, virtualNetworkName, virtualNetwork)
	//JEB	if err != nil {
	//JEB		return nil, err
	//JEB	}
	//JEB	return &sgFuture, nil
	return nil, nil
}

// WaitForVnetCreateOrUpdateFuture returns when the CreateOrUpdateVnet operation completes.
func (s *Service) WaitForVnetCreateOrUpdateFuture(future armada.VirtualNetworksCreateOrUpdateFuture) error {
	//JEB return future.Future.WaitForCompletionRef(s.scope.Context, s.scope.AirshipClients.VirtualNetworks.Client)
	return nil
}
