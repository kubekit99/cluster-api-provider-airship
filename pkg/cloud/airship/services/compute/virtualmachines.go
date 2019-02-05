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
	//JEB "github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/autorest"
	//JEB "github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/autorest/to"
	"github.com/kubekit99/cluster-api-provider-airship/pkg/airship-go-api/services/compute"
)

// RunCommand executes a command on the VM.
func (s *Service) RunCommand(resoureGroup string, name string, cmd string) (compute.VirtualMachinesRunCommandFuture, error) {
	//JEB cmdInput := compute.RunCommandInput{
	//JEB		CommandID: to.StringPtr("RunShellScript"),
	//JEB		Script:    to.StringSlicePtr([]string{cmd}),
	//JEB }
	//JEB return s.scope.AirshipClients.VM.RunCommand(s.scope.Context, resoureGroup, name, cmdInput)
	return compute.VirtualMachinesRunCommandFuture{}, nil
}

// VMIfExists returns the reference to the VM object if it exists.
func (s *Service) VMIfExists(resourceGroup string, name string) (*compute.VirtualMachine, error) {
	//JEB vm, err := s.scope.AirshipClients.VM.Get(s.scope.Context, resourceGroup, name, "")
	//JEB if err != nil {
	//JEB 		if aerr, ok := err.(autorest.DetailedError); ok {
	//JEB 			if aerr.StatusCode.(int) == 404 {
	//JEB 				return nil, nil
	//JEB 			}
	//JEB 		}
	//JEB 		return nil, err
	//JEB 	}
	//JEB return &vm, nil
	return nil, nil
}

// DeleteVM deletes the virtual machine.
func (s *Service) DeleteVM(resourceGroup string, name string) (compute.VirtualMachinesDeleteFuture, error) {
	//JEB return s.scope.AirshipClients.VM.Delete(s.scope.Context, resourceGroup, name)
	return compute.VirtualMachinesDeleteFuture{}, nil
}

// WaitForVMRunCommandFuture returns when the RunCommand operation completes.
func (s *Service) WaitForVMRunCommandFuture(future compute.VirtualMachinesRunCommandFuture) error {
	//JEB return future.Future.WaitForCompletionRef(s.scope.Context, s.scope.AirshipClients.VM.Client)
	return nil
}

// WaitForVMDeletionFuture returns when the DeleteVM operation completes.
func (s *Service) WaitForVMDeletionFuture(future compute.VirtualMachinesDeleteFuture) error {
	//JEB return future.Future.WaitForCompletionRef(s.scope.Context, s.scope.AirshipClients.VM.Client)
	return nil
}
