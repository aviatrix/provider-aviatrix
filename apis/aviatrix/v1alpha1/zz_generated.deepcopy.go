//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Account) DeepCopyInto(out *Account) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Account.
func (in *Account) DeepCopy() *Account {
	if in == nil {
		return nil
	}
	out := new(Account)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Account) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountList) DeepCopyInto(out *AccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Account, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountList.
func (in *AccountList) DeepCopy() *AccountList {
	if in == nil {
		return nil
	}
	out := new(AccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountObservation) DeepCopyInto(out *AccountObservation) {
	*out = *in
	if in.AwsCACertPath != nil {
		in, out := &in.AwsCACertPath, &out.AwsCACertPath
		*out = new(string)
		**out = **in
	}
	if in.AwssCapCertKeyPath != nil {
		in, out := &in.AwssCapCertKeyPath, &out.AwssCapCertKeyPath
		*out = new(string)
		**out = **in
	}
	if in.AwssCapCertPath != nil {
		in, out := &in.AwssCapCertPath, &out.AwssCapCertPath
		*out = new(string)
		**out = **in
	}
	if in.AwstsCapCertKeyPath != nil {
		in, out := &in.AwstsCapCertKeyPath, &out.AwstsCapCertKeyPath
		*out = new(string)
		**out = **in
	}
	if in.AwstsCapCertPath != nil {
		in, out := &in.AwstsCapCertPath, &out.AwstsCapCertPath
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountObservation.
func (in *AccountObservation) DeepCopy() *AccountObservation {
	if in == nil {
		return nil
	}
	out := new(AccountObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountParameters) DeepCopyInto(out *AccountParameters) {
	*out = *in
	if in.AlicloudAccessKeySecretRef != nil {
		in, out := &in.AlicloudAccessKeySecretRef, &out.AlicloudAccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AlicloudAccountID != nil {
		in, out := &in.AlicloudAccountID, &out.AlicloudAccountID
		*out = new(string)
		**out = **in
	}
	if in.AlicloudSecretKeySecretRef != nil {
		in, out := &in.AlicloudSecretKeySecretRef, &out.AlicloudSecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ArmApplicationIDSecretRef != nil {
		in, out := &in.ArmApplicationIDSecretRef, &out.ArmApplicationIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ArmApplicationKeySecretRef != nil {
		in, out := &in.ArmApplicationKeySecretRef, &out.ArmApplicationKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ArmDirectoryIDSecretRef != nil {
		in, out := &in.ArmDirectoryIDSecretRef, &out.ArmDirectoryIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.ArmSubscriptionID != nil {
		in, out := &in.ArmSubscriptionID, &out.ArmSubscriptionID
		*out = new(string)
		**out = **in
	}
	if in.AuditAccount != nil {
		in, out := &in.AuditAccount, &out.AuditAccount
		*out = new(bool)
		**out = **in
	}
	if in.AwsAccessKeySecretRef != nil {
		in, out := &in.AwsAccessKeySecretRef, &out.AwsAccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwsAccountNumber != nil {
		in, out := &in.AwsAccountNumber, &out.AwsAccountNumber
		*out = new(string)
		**out = **in
	}
	if in.AwsGatewayRoleApp != nil {
		in, out := &in.AwsGatewayRoleApp, &out.AwsGatewayRoleApp
		*out = new(string)
		**out = **in
	}
	if in.AwsGatewayRoleEC2 != nil {
		in, out := &in.AwsGatewayRoleEC2, &out.AwsGatewayRoleEC2
		*out = new(string)
		**out = **in
	}
	if in.AwsIAM != nil {
		in, out := &in.AwsIAM, &out.AwsIAM
		*out = new(bool)
		**out = **in
	}
	if in.AwsRoleApp != nil {
		in, out := &in.AwsRoleApp, &out.AwsRoleApp
		*out = new(string)
		**out = **in
	}
	if in.AwsRoleEC2 != nil {
		in, out := &in.AwsRoleEC2, &out.AwsRoleEC2
		*out = new(string)
		**out = **in
	}
	if in.AwsSecretKeySecretRef != nil {
		in, out := &in.AwsSecretKeySecretRef, &out.AwsSecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwschinaAccessKeySecretRef != nil {
		in, out := &in.AwschinaAccessKeySecretRef, &out.AwschinaAccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwschinaAccountNumber != nil {
		in, out := &in.AwschinaAccountNumber, &out.AwschinaAccountNumber
		*out = new(string)
		**out = **in
	}
	if in.AwschinaIAM != nil {
		in, out := &in.AwschinaIAM, &out.AwschinaIAM
		*out = new(bool)
		**out = **in
	}
	if in.AwschinaRoleApp != nil {
		in, out := &in.AwschinaRoleApp, &out.AwschinaRoleApp
		*out = new(string)
		**out = **in
	}
	if in.AwschinaRoleEC2 != nil {
		in, out := &in.AwschinaRoleEC2, &out.AwschinaRoleEC2
		*out = new(string)
		**out = **in
	}
	if in.AwschinaSecretKeySecretRef != nil {
		in, out := &in.AwschinaSecretKeySecretRef, &out.AwschinaSecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwsgovAccessKeySecretRef != nil {
		in, out := &in.AwsgovAccessKeySecretRef, &out.AwsgovAccessKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwsgovAccountNumber != nil {
		in, out := &in.AwsgovAccountNumber, &out.AwsgovAccountNumber
		*out = new(string)
		**out = **in
	}
	if in.AwsgovIAM != nil {
		in, out := &in.AwsgovIAM, &out.AwsgovIAM
		*out = new(bool)
		**out = **in
	}
	if in.AwsgovRoleApp != nil {
		in, out := &in.AwsgovRoleApp, &out.AwsgovRoleApp
		*out = new(string)
		**out = **in
	}
	if in.AwsgovRoleEC2 != nil {
		in, out := &in.AwsgovRoleEC2, &out.AwsgovRoleEC2
		*out = new(string)
		**out = **in
	}
	if in.AwsgovSecretKeySecretRef != nil {
		in, out := &in.AwsgovSecretKeySecretRef, &out.AwsgovSecretKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwssAccountNumber != nil {
		in, out := &in.AwssAccountNumber, &out.AwssAccountNumber
		*out = new(string)
		**out = **in
	}
	if in.AwssCAChainCertSecretRef != nil {
		in, out := &in.AwssCAChainCertSecretRef, &out.AwssCAChainCertSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwssCapAccountName != nil {
		in, out := &in.AwssCapAccountName, &out.AwssCapAccountName
		*out = new(string)
		**out = **in
	}
	if in.AwssCapAgency != nil {
		in, out := &in.AwssCapAgency, &out.AwssCapAgency
		*out = new(string)
		**out = **in
	}
	if in.AwssCapCertKeySecretRef != nil {
		in, out := &in.AwssCapCertKeySecretRef, &out.AwssCapCertKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwssCapCertSecretRef != nil {
		in, out := &in.AwssCapCertSecretRef, &out.AwssCapCertSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwssCapRoleName != nil {
		in, out := &in.AwssCapRoleName, &out.AwssCapRoleName
		*out = new(string)
		**out = **in
	}
	if in.AwssCapURL != nil {
		in, out := &in.AwssCapURL, &out.AwssCapURL
		*out = new(string)
		**out = **in
	}
	if in.AwstsAccountNumber != nil {
		in, out := &in.AwstsAccountNumber, &out.AwstsAccountNumber
		*out = new(string)
		**out = **in
	}
	if in.AwstsCAChainCertSecretRef != nil {
		in, out := &in.AwstsCAChainCertSecretRef, &out.AwstsCAChainCertSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwstsCapAgency != nil {
		in, out := &in.AwstsCapAgency, &out.AwstsCapAgency
		*out = new(string)
		**out = **in
	}
	if in.AwstsCapCertKeySecretRef != nil {
		in, out := &in.AwstsCapCertKeySecretRef, &out.AwstsCapCertKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwstsCapCertSecretRef != nil {
		in, out := &in.AwstsCapCertSecretRef, &out.AwstsCapCertSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AwstsCapMission != nil {
		in, out := &in.AwstsCapMission, &out.AwstsCapMission
		*out = new(string)
		**out = **in
	}
	if in.AwstsCapRoleName != nil {
		in, out := &in.AwstsCapRoleName, &out.AwstsCapRoleName
		*out = new(string)
		**out = **in
	}
	if in.AwstsCapURL != nil {
		in, out := &in.AwstsCapURL, &out.AwstsCapURL
		*out = new(string)
		**out = **in
	}
	if in.AzurechinaApplicationIDSecretRef != nil {
		in, out := &in.AzurechinaApplicationIDSecretRef, &out.AzurechinaApplicationIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AzurechinaApplicationKeySecretRef != nil {
		in, out := &in.AzurechinaApplicationKeySecretRef, &out.AzurechinaApplicationKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AzurechinaDirectoryIDSecretRef != nil {
		in, out := &in.AzurechinaDirectoryIDSecretRef, &out.AzurechinaDirectoryIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AzurechinaSubscriptionID != nil {
		in, out := &in.AzurechinaSubscriptionID, &out.AzurechinaSubscriptionID
		*out = new(string)
		**out = **in
	}
	if in.AzuregovApplicationIDSecretRef != nil {
		in, out := &in.AzuregovApplicationIDSecretRef, &out.AzuregovApplicationIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AzuregovApplicationKeySecretRef != nil {
		in, out := &in.AzuregovApplicationKeySecretRef, &out.AzuregovApplicationKeySecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AzuregovDirectoryIDSecretRef != nil {
		in, out := &in.AzuregovDirectoryIDSecretRef, &out.AzuregovDirectoryIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.AzuregovSubscriptionID != nil {
		in, out := &in.AzuregovSubscriptionID, &out.AzuregovSubscriptionID
		*out = new(string)
		**out = **in
	}
	if in.CloudType != nil {
		in, out := &in.CloudType, &out.CloudType
		*out = new(float64)
		**out = **in
	}
	if in.GcloudProjectCredentialsFilepath != nil {
		in, out := &in.GcloudProjectCredentialsFilepath, &out.GcloudProjectCredentialsFilepath
		*out = new(string)
		**out = **in
	}
	if in.GcloudProjectID != nil {
		in, out := &in.GcloudProjectID, &out.GcloudProjectID
		*out = new(string)
		**out = **in
	}
	if in.OciAPIPrivateKeyFilepathSecretRef != nil {
		in, out := &in.OciAPIPrivateKeyFilepathSecretRef, &out.OciAPIPrivateKeyFilepathSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.OciCompartmentIDSecretRef != nil {
		in, out := &in.OciCompartmentIDSecretRef, &out.OciCompartmentIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.OciTenancyIDSecretRef != nil {
		in, out := &in.OciTenancyIDSecretRef, &out.OciTenancyIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.OciUserIDSecretRef != nil {
		in, out := &in.OciUserIDSecretRef, &out.OciUserIDSecretRef
		*out = new(v1.SecretKeySelector)
		**out = **in
	}
	if in.RbacGroups != nil {
		in, out := &in.RbacGroups, &out.RbacGroups
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountParameters.
func (in *AccountParameters) DeepCopy() *AccountParameters {
	if in == nil {
		return nil
	}
	out := new(AccountParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpec) DeepCopyInto(out *AccountSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpec.
func (in *AccountSpec) DeepCopy() *AccountSpec {
	if in == nil {
		return nil
	}
	out := new(AccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountStatus) DeepCopyInto(out *AccountStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountStatus.
func (in *AccountStatus) DeepCopy() *AccountStatus {
	if in == nil {
		return nil
	}
	out := new(AccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateSubnetsObservation) DeepCopyInto(out *PrivateSubnetsObservation) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateSubnetsObservation.
func (in *PrivateSubnetsObservation) DeepCopy() *PrivateSubnetsObservation {
	if in == nil {
		return nil
	}
	out := new(PrivateSubnetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateSubnetsParameters) DeepCopyInto(out *PrivateSubnetsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateSubnetsParameters.
func (in *PrivateSubnetsParameters) DeepCopy() *PrivateSubnetsParameters {
	if in == nil {
		return nil
	}
	out := new(PrivateSubnetsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicSubnetsObservation) DeepCopyInto(out *PublicSubnetsObservation) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicSubnetsObservation.
func (in *PublicSubnetsObservation) DeepCopy() *PublicSubnetsObservation {
	if in == nil {
		return nil
	}
	out := new(PublicSubnetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublicSubnetsParameters) DeepCopyInto(out *PublicSubnetsParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublicSubnetsParameters.
func (in *PublicSubnetsParameters) DeepCopy() *PublicSubnetsParameters {
	if in == nil {
		return nil
	}
	out := new(PublicSubnetsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetsObservation) DeepCopyInto(out *SubnetsObservation) {
	*out = *in
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetsObservation.
func (in *SubnetsObservation) DeepCopy() *SubnetsObservation {
	if in == nil {
		return nil
	}
	out := new(SubnetsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubnetsParameters) DeepCopyInto(out *SubnetsParameters) {
	*out = *in
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubnetsParameters.
func (in *SubnetsParameters) DeepCopy() *SubnetsParameters {
	if in == nil {
		return nil
	}
	out := new(SubnetsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPC) DeepCopyInto(out *VPC) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPC.
func (in *VPC) DeepCopy() *VPC {
	if in == nil {
		return nil
	}
	out := new(VPC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPC) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCList) DeepCopyInto(out *VPCList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPC, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCList.
func (in *VPCList) DeepCopy() *VPCList {
	if in == nil {
		return nil
	}
	out := new(VPCList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCObservation) DeepCopyInto(out *VPCObservation) {
	*out = *in
	if in.AvailabilityDomains != nil {
		in, out := &in.AvailabilityDomains, &out.AvailabilityDomains
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AzureVnetResourceID != nil {
		in, out := &in.AzureVnetResourceID, &out.AzureVnetResourceID
		*out = new(string)
		**out = **in
	}
	if in.FaultDomains != nil {
		in, out := &in.FaultDomains, &out.FaultDomains
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.PrivateSubnets != nil {
		in, out := &in.PrivateSubnets, &out.PrivateSubnets
		*out = make([]PrivateSubnetsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PublicSubnets != nil {
		in, out := &in.PublicSubnets, &out.PublicSubnets
		*out = make([]PublicSubnetsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RouteTables != nil {
		in, out := &in.RouteTables, &out.RouteTables
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]SubnetsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCObservation.
func (in *VPCObservation) DeepCopy() *VPCObservation {
	if in == nil {
		return nil
	}
	out := new(VPCObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCParameters) DeepCopyInto(out *VPCParameters) {
	*out = *in
	if in.AccountName != nil {
		in, out := &in.AccountName, &out.AccountName
		*out = new(string)
		**out = **in
	}
	if in.AviatrixFirenetVPC != nil {
		in, out := &in.AviatrixFirenetVPC, &out.AviatrixFirenetVPC
		*out = new(bool)
		**out = **in
	}
	if in.AviatrixTransitVPC != nil {
		in, out := &in.AviatrixTransitVPC, &out.AviatrixTransitVPC
		*out = new(bool)
		**out = **in
	}
	if in.Cidr != nil {
		in, out := &in.Cidr, &out.Cidr
		*out = new(string)
		**out = **in
	}
	if in.CloudType != nil {
		in, out := &in.CloudType, &out.CloudType
		*out = new(float64)
		**out = **in
	}
	if in.EnableNativeGwlb != nil {
		in, out := &in.EnableNativeGwlb, &out.EnableNativeGwlb
		*out = new(bool)
		**out = **in
	}
	if in.EnablePrivateOobSubnet != nil {
		in, out := &in.EnablePrivateOobSubnet, &out.EnablePrivateOobSubnet
		*out = new(bool)
		**out = **in
	}
	if in.NumOfSubnetPairs != nil {
		in, out := &in.NumOfSubnetPairs, &out.NumOfSubnetPairs
		*out = new(float64)
		**out = **in
	}
	if in.PrivateModeSubnets != nil {
		in, out := &in.PrivateModeSubnets, &out.PrivateModeSubnets
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ResourceGroup != nil {
		in, out := &in.ResourceGroup, &out.ResourceGroup
		*out = new(string)
		**out = **in
	}
	if in.SubnetSize != nil {
		in, out := &in.SubnetSize, &out.SubnetSize
		*out = new(float64)
		**out = **in
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]SubnetsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCParameters.
func (in *VPCParameters) DeepCopy() *VPCParameters {
	if in == nil {
		return nil
	}
	out := new(VPCParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCSpec) DeepCopyInto(out *VPCSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCSpec.
func (in *VPCSpec) DeepCopy() *VPCSpec {
	if in == nil {
		return nil
	}
	out := new(VPCSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCStatus) DeepCopyInto(out *VPCStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCStatus.
func (in *VPCStatus) DeepCopy() *VPCStatus {
	if in == nil {
		return nil
	}
	out := new(VPCStatus)
	in.DeepCopyInto(out)
	return out
}
