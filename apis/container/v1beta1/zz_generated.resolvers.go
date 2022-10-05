/*
Copyright 2021 The Crossplane Authors.

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
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	v1beta11 "github.com/upbound/official-providers/provider-gcp/apis/cloudplatform/v1beta1"
	v1beta1 "github.com/upbound/official-providers/provider-gcp/apis/compute/v1beta1"
	common "github.com/upbound/official-providers/provider-gcp/config/common"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Cluster.
func (mg *Cluster) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Network),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.NetworkRef,
		Selector:     mg.Spec.ForProvider.NetworkSelector,
		To: reference.To{
			List:    &v1beta1.NetworkList{},
			Managed: &v1beta1.Network{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Network")
	}
	mg.Spec.ForProvider.Network = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NetworkRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NodeConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeConfig[i3].ServiceAccount),
			Extract:      resource.ExtractParamPath("email", true),
			Reference:    mg.Spec.ForProvider.NodeConfig[i3].ServiceAccountRef,
			Selector:     mg.Spec.ForProvider.NodeConfig[i3].ServiceAccountSelector,
			To: reference.To{
				List:    &v1beta11.ServiceAccountList{},
				Managed: &v1beta11.ServiceAccount{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NodeConfig[i3].ServiceAccount")
		}
		mg.Spec.ForProvider.NodeConfig[i3].ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NodeConfig[i3].ServiceAccountRef = rsp.ResolvedReference

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Subnetwork),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetworkRef,
		Selector:     mg.Spec.ForProvider.SubnetworkSelector,
		To: reference.To{
			List:    &v1beta1.SubnetworkList{},
			Managed: &v1beta1.Subnetwork{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Subnetwork")
	}
	mg.Spec.ForProvider.Subnetwork = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetworkRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this NodePool.
func (mg *NodePool) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Cluster),
		Extract:      common.ExtractResourceID(),
		Reference:    mg.Spec.ForProvider.ClusterRef,
		Selector:     mg.Spec.ForProvider.ClusterSelector,
		To: reference.To{
			List:    &ClusterList{},
			Managed: &Cluster{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Cluster")
	}
	mg.Spec.ForProvider.Cluster = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.ClusterRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.NodeConfig); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NodeConfig[i3].ServiceAccount),
			Extract:      resource.ExtractParamPath("email", true),
			Reference:    mg.Spec.ForProvider.NodeConfig[i3].ServiceAccountRef,
			Selector:     mg.Spec.ForProvider.NodeConfig[i3].ServiceAccountSelector,
			To: reference.To{
				List:    &v1beta11.ServiceAccountList{},
				Managed: &v1beta11.ServiceAccount{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NodeConfig[i3].ServiceAccount")
		}
		mg.Spec.ForProvider.NodeConfig[i3].ServiceAccount = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NodeConfig[i3].ServiceAccountRef = rsp.ResolvedReference

	}

	return nil
}