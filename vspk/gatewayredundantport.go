/*
  Copyright (c) 2015, Alcatel-Lucent Inc
  All rights reserved.

  Redistribution and use in source and binary forms, with or without
  modification, are permitted provided that the following conditions are met:
      * Redistributions of source code must retain the above copyright
        notice, this list of conditions and the following disclaimer.
      * Redistributions in binary form must reproduce the above copyright
        notice, this list of conditions and the following disclaimer in the
        documentation and/or other materials provided with the distribution.
      * Neither the name of the copyright holder nor the names of its contributors
        may be used to endorse or promote products derived from this software without
        specific prior written permission.

  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
  ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
  WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY
  DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
  (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
  LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
  ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
  (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
  SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package vspk

import "github.com/nuagenetworks/go-bambou/bambou"

// GatewayRedundantPortIdentity represents the Identity of the object
var GatewayRedundantPortIdentity = bambou.Identity{
	Name:     "gatewayredundantport",
	Category: "gatewayredundantports",
}

// GatewayRedundantPortsList represents a list of GatewayRedundantPorts
type GatewayRedundantPortsList []*GatewayRedundantPort

// GatewayRedundantPortsAncestor is the interface that an ancestor of a GatewayRedundantPort must implement.
// An Ancestor is defined as an entity that has GatewayRedundantPort as a descendant.
// An Ancestor can get a list of its child GatewayRedundantPorts, but not necessarily create one.
type GatewayRedundantPortsAncestor interface {
	GatewayRedundantPorts(*bambou.FetchingInfo) (GatewayRedundantPortsList, *bambou.Error)
}

// GatewayRedundantPortsParent is the interface that a parent of a GatewayRedundantPort must implement.
// A Parent is defined as an entity that has GatewayRedundantPort as a child.
// A Parent is an Ancestor which can create a GatewayRedundantPort.
type GatewayRedundantPortsParent interface {
	GatewayRedundantPortsAncestor
	CreateGatewayRedundantPort(*GatewayRedundantPort) *bambou.Error
}

// GatewayRedundantPort represents the model of a gatewayredundantport
type GatewayRedundantPort struct {
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	VLANRange                   string `json:"VLANRange,omitempty"`
	Name                        string `json:"name,omitempty"`
	PermittedAction             string `json:"permittedAction,omitempty"`
	Description                 string `json:"description,omitempty"`
	PhysicalName                string `json:"physicalName,omitempty"`
	PortPeer1ID                 string `json:"portPeer1ID,omitempty"`
	PortPeer2ID                 string `json:"portPeer2ID,omitempty"`
	PortType                    string `json:"portType,omitempty"`
	UseUserMnemonic             bool   `json:"useUserMnemonic"`
	UserMnemonic                string `json:"userMnemonic,omitempty"`
	AssociatedEgressQOSPolicyID string `json:"associatedEgressQOSPolicyID,omitempty"`
	Status                      string `json:"status,omitempty"`
}

// NewGatewayRedundantPort returns a new *GatewayRedundantPort
func NewGatewayRedundantPort() *GatewayRedundantPort {

	return &GatewayRedundantPort{}
}

// Identity returns the Identity of the object.
func (o *GatewayRedundantPort) Identity() bambou.Identity {

	return GatewayRedundantPortIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *GatewayRedundantPort) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *GatewayRedundantPort) SetIdentifier(ID string) {

	o.ID = ID
}

// Fetch retrieves the GatewayRedundantPort from the server
func (o *GatewayRedundantPort) Fetch() *bambou.Error {

	return bambou.CurrentSession().FetchEntity(o)
}

// Save saves the GatewayRedundantPort into the server
func (o *GatewayRedundantPort) Save() *bambou.Error {

	return bambou.CurrentSession().SaveEntity(o)
}

// Delete deletes the GatewayRedundantPort from the server
func (o *GatewayRedundantPort) Delete() *bambou.Error {

	return bambou.CurrentSession().DeleteEntity(o)
}

// VLANs retrieves the list of child VLANs of the GatewayRedundantPort
func (o *GatewayRedundantPort) VLANs(info *bambou.FetchingInfo) (VLANsList, *bambou.Error) {

	var list VLANsList
	err := bambou.CurrentSession().FetchChildren(o, VLANIdentity, &list, info)
	return list, err
}

// CreateVLAN creates a new child VLAN under the GatewayRedundantPort
func (o *GatewayRedundantPort) CreateVLAN(child *VLAN) *bambou.Error {

	return bambou.CurrentSession().CreateChild(o, child)
}
