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

package network

import (
	//JEB "github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/autorest"
	//JEB "github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/autorest/to"
	"github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/services/network"
)

const (
	// SecurityGroupDefaultName is the default name for the network security group of the cluster.
	SecurityGroupDefaultName = "ClusterAPINSG"
)

// NetworkSGIfExists returns the nsg reference if the nsg resource exists.
func (s *Service) NetworkSGIfExists(resourceGroupName string, networkSecurityGroupName string) (*network.SecurityGroup, error) {
	//JEB networkSG, err := s.scope.AirshipClients.SecurityGroups.Get(s.scope.Context, resourceGroupName, networkSecurityGroupName, "")
	//JEB if err != nil {
	//JEB 		if aerr, ok := err.(autorest.DetailedError); ok {
	//JEB 			if aerr.StatusCode.(int) == 404 {
	//JEB 				return nil, nil
	//JEB 			}
	//JEB 		}
	//JEB 		return nil, err
	//JEB 	}
	//JEB 	return &networkSG, nil
	return nil, nil
}

// CreateOrUpdateNetworkSecurityGroup creates or updates the nsg resource.
func (s *Service) CreateOrUpdateNetworkSecurityGroup(resourceGroupName string, networkSecurityGroupName string, location string) (*network.SecurityGroupsCreateOrUpdateFuture, error) {
	//JEB if networkSecurityGroupName == "" {
	//JEB networkSecurityGroupName = SecurityGroupDefaultName
	//JEB }
	//JEB sshInbound := network.SecurityRule{
	//JEB Name: to.StringPtr("ClusterAPISSH"),
	//JEB SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
	//JEB Protocol:                 network.SecurityRuleProtocolTCP,
	//JEB SourcePortRange:          to.StringPtr("*"),
	//JEB DestinationPortRange:     to.StringPtr("22"),
	//JEB SourceAddressPrefix:      to.StringPtr("*"),
	//JEB DestinationAddressPrefix: to.StringPtr("*"),
	//JEB Priority:                 to.Int32Ptr(1000),
	//JEB Direction:                network.SecurityRuleDirectionInbound,
	//JEB Access:                   network.SecurityRuleAccessAllow,
	//JEB },
	//JEB }
	//JEB
	//JEB kubernetesInbound := network.SecurityRule{
	//JEB Name: to.StringPtr("KubernetesAPI"),
	//JEB SecurityRulePropertiesFormat: &network.SecurityRulePropertiesFormat{
	//JEB Protocol:                 network.SecurityRuleProtocolTCP,
	//JEB SourcePortRange:          to.StringPtr("*"),
	//JEB DestinationPortRange:     to.StringPtr("6443"),
	//JEB SourceAddressPrefix:      to.StringPtr("*"),
	//JEB DestinationAddressPrefix: to.StringPtr("*"),
	//JEB Priority:                 to.Int32Ptr(1001),
	//JEB Direction:                network.SecurityRuleDirectionInbound,
	//JEB Access:                   network.SecurityRuleAccessAllow,
	//JEB },
	//JEB }
	//JEB
	//JEB securityGroupProperties := network.SecurityGroupPropertiesFormat{
	//JEB SecurityRules: &[]network.SecurityRule{sshInbound, kubernetesInbound},
	//JEB }
	//JEB securityGroup := network.SecurityGroup{
	//JEB Location:                      to.StringPtr(location),
	//JEB SecurityGroupPropertiesFormat: &securityGroupProperties,
	//JEB }
	//JEB sgFuture, err := s.scope.AirshipClients.SecurityGroups.CreateOrUpdate(s.scope.Context, resourceGroupName, networkSecurityGroupName, securityGroup)
	//JEB if err != nil {
	//JEB return nil, err
	//JEB }
	//JEB return &sgFuture, nil
	return nil, nil
}

// DeleteNetworkSecurityGroup deletes the nsg resource.
func (s *Service) DeleteNetworkSecurityGroup(resourceGroupName string, networkSecurityGroupName string) (network.SecurityGroupsDeleteFuture, error) {
	//JEB return s.scope.AirshipClients.SecurityGroups.Delete(s.scope.Context, resourceGroupName, networkSecurityGroupName)
	return network.SecurityGroupsDeleteFuture{}, nil
}

// WaitForNetworkSGsCreateOrUpdateFuture returns when the CreateOrUpdateNetworkSecurityGroup operation completes.
func (s *Service) WaitForNetworkSGsCreateOrUpdateFuture(future network.SecurityGroupsCreateOrUpdateFuture) error {
	//JEB return future.Future.WaitForCompletionRef(s.scope.Context, s.scope.AirshipClients.SecurityGroups.Client)
	return nil
}
