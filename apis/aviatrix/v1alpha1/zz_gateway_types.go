/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GatewayObservation struct {

	// Cloud instance ID of the gateway.
	// Instance ID of the gateway.
	CloudInstanceID *string `json:"cloudInstanceId,omitempty" tf:"cloud_instance_id,omitempty"`

	// ELB DNS name.
	// ELB DNS Name.
	ELBDNSName *string `json:"elbDnsName,omitempty" tf:"elb_dns_name,omitempty"`

	// The lan interface id of the of FQDN gateway with additional LAN interface. This attribute will be exported when enabling FQDN gateway firenet in Azure. Available in provider version R2.17.1+.
	// FQDN gateway lan interface id.
	FqdnLanInterface *string `json:"fqdnLanInterface,omitempty" tf:"fqdn_lan_interface,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The image version of the gateway. Use aviatrix_gateway_image data source to programmatically retrieve this value for the desired software_version. If set, we will attempt to update the gateway to the specified version if current version is different. If left blank, the gateway upgrades can be managed with the aviatrix_controller_config resource. Type: String. Example: "hvm-cloudx-aws-022021". Available as of provider version R2.20.0.
	// image_version can be used to set the desired image version of the gateway. If set, we will attempt to update the gateway to the specified version.
	ImageVersion *string `json:"imageVersion,omitempty" tf:"image_version,omitempty"`

	// Cloud instance ID of the HA gateway.
	// Instance ID of the peering HA gateway.
	PeeringHaCloudInstanceID *string `json:"peeringHaCloudInstanceId,omitempty" tf:"peering_ha_cloud_instance_id,omitempty"`

	// Aviatrix gateway unique name of HA gateway.
	// Aviatrix gateway unique name of HA gateway.
	PeeringHaGwName *string `json:"peeringHaGwName,omitempty" tf:"peering_ha_gw_name,omitempty"`

	// Private IP address of HA gateway.
	// Private IP address of HA gateway.
	PeeringHaPrivateIP *string `json:"peeringHaPrivateIp,omitempty" tf:"peering_ha_private_ip,omitempty"`

	// HA security group used for the gateway.
	// Peering HA security group used for the gateway.
	PeeringHaSecurityGroupID *string `json:"peeringHaSecurityGroupId,omitempty" tf:"peering_ha_security_group_id,omitempty"`

	// Private IP address of the gateway created.
	// Private IP address of the Gateway created.
	PrivateIP *string `json:"privateIp,omitempty" tf:"private_ip,omitempty"`

	// DNS server used by the gateway. Default is "8.8.8.8", can be overridden with the VPC's setting.
	// NS server used by the gateway.
	PublicDNSServer *string `json:"publicDnsServer,omitempty" tf:"public_dns_server,omitempty"`

	// Security group used for the gateway.
	// Security group used for the gateway.
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// The software version of the gateway. If set, we will attempt to update the gateway to the specified version if current version is different. If left blank, the gateway upgrade can be managed with the aviatrix_controller_config resource. Type: String. Example: "6.5.821". Available as of provider version R2.20.0.
	// software_version can be used to set the desired software version of the gateway. If set, we will attempt to update the gateway to the specified version. If left blank, the gateway software version will continue to be managed through the aviatrix_controller_config resource.
	SoftwareVersion *string `json:"softwareVersion,omitempty" tf:"software_version,omitempty"`
}

type GatewayParameters struct {

	// Account name. This account will be used to launch Aviatrix gateway.
	// Account name. This account will be used to launch Aviatrix gateway.
	// +kubebuilder:validation:Required
	AccountName *string `json:"accountName" tf:"account_name,omitempty"`

	// A list of destination CIDR ranges that will also go through the VPN tunnel when Split Tunnel Mode is enabled.
	// A list of destination CIDR ranges that will also go through the VPN tunnel when Split Tunnel Mode is enabled.
	// +kubebuilder:validation:Optional
	AdditionalCidrs *string `json:"additionalCidrs,omitempty" tf:"additional_cidrs,omitempty"`

	// A list of CIDR ranges separated by comma to configure when "Designated Gateway" feature is enabled. Example: "10.8.0.0/16,10.9.0.0/16,10.10.0.0/16".
	// A list of CIDR ranges separated by comma to configure when 'designated_gateway' feature is enabled.
	// +kubebuilder:validation:Optional
	AdditionalCidrsDesignatedGateway *string `json:"additionalCidrsDesignatedGateway,omitempty" tf:"additional_cidrs_designated_gateway,omitempty"`

	// If set to false, use an available address in Elastic IP pool for this gateway. Otherwise, allocate a new Elastic IP and use it for this gateway. Available in Controller 2.7+. Valid values: true, false. Default: true.
	// When value is false, reuse an idle address in Elastic IP pool for this gateway. Otherwise, allocate a new Elastic IP and use it for this gateway.
	// +kubebuilder:validation:Optional
	AllocateNewEIP *bool `json:"allocateNewEip,omitempty" tf:"allocate_new_eip,omitempty"`

	// Availability domain. Required and valid only for OCI. Available as of provider version R2.19.3.
	// Availability domain for OCI.
	// +kubebuilder:validation:Optional
	AvailabilityDomain *string `json:"availabilityDomain,omitempty" tf:"availability_domain,omitempty"`

	// Name of public IP Address resource and its resource group in Azure to be assigned to the gateway instance. Example: "IP_Name:Resource_Group_Name". Required when allocate_new_eip is false and cloud_type is Azure, AzureGov or AzureChina. Available as of provider version 2.20+.
	// The name of the public IP address and its resource group in Azure to assign to this Gateway.
	// +kubebuilder:validation:Optional
	AzureEIPNameResourceGroup *string `json:"azureEipNameResourceGroup,omitempty" tf:"azure_eip_name_resource_group,omitempty"`

	// Cloud service provider to use to launch the gateway. Requires an integer value. Currently supports AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud (8192), AWS Top Secret (16384) and AWS Secret (32768).
	// Type of cloud service provider.
	// +kubebuilder:validation:Required
	CloudType *float64 `json:"cloudType" tf:"cloud_type,omitempty"`

	// Customer-managed key ID.
	// Customer managed key ID.
	// +kubebuilder:validation:Optional
	CustomerManagedKeysSecretRef *v1.SecretKeySelector `json:"customerManagedKeysSecretRef,omitempty" tf:"-"`

	// API hostname for DUO auth mode. Required: Yes if otp_mode is "2".
	// API hostname for DUO auth mode.
	// +kubebuilder:validation:Optional
	DuoAPIHostname *string `json:"duoApiHostname,omitempty" tf:"duo_api_hostname,omitempty"`

	// Integration key for DUO auth mode. Required if otp_mode is "2".
	// Integration key for DUO auth mode.
	// +kubebuilder:validation:Optional
	DuoIntegrationKey *string `json:"duoIntegrationKey,omitempty" tf:"duo_integration_key,omitempty"`

	// Push mode for DUO auth. Required if otp_mode is "2". Valid values: "auto", "selective" and "token".
	// Push mode for DUO auth.
	// +kubebuilder:validation:Optional
	DuoPushMode *string `json:"duoPushMode,omitempty" tf:"duo_push_mode,omitempty"`

	// Secret key for DUO auth mode. Required if otp_mode is "2".
	// Secret key for DUO auth mode.
	// +kubebuilder:validation:Optional
	DuoSecretKeySecretRef *v1.SecretKeySelector `json:"duoSecretKeySecretRef,omitempty" tf:"-"`

	// Specified EIP to use for gateway creation. Required when allocate_new_eip is false.  Available in Controller version 3.5+. Only available for AWS, GCP, Azure, OCI, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret.
	// Required when allocate_new_eip is false. It uses specified EIP for this gateway.
	// +kubebuilder:validation:Optional
	EIP *string `json:"eip,omitempty" tf:"eip,omitempty"`

	// A name for the ELB that is created. If it is not specified, a name is generated automatically.
	// A name for the ELB that is created.
	// +kubebuilder:validation:Optional
	ELBName *string `json:"elbName,omitempty" tf:"elb_name,omitempty"`

	// Enable Designated Gateway feature for Gateway. Only supported for AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret gateways. Valid values: true, false. Default value: false. Please view documentation here for more information on this feature.
	// Enable 'designated_gateway' feature for Gateway. Valid values: true, false.
	// +kubebuilder:validation:Optional
	EnableDesignatedGateway *bool `json:"enableDesignatedGateway,omitempty" tf:"enable_designated_gateway,omitempty"`

	// Specify whether to enable ELB or not. Not supported for OCI gateways. Valid values: true, false.
	// Specify whether to enable ELB or not.
	// +kubebuilder:validation:Optional
	EnableELB *bool `json:"enableElb,omitempty" tf:"enable_elb,omitempty"`

	// Enable EBS volume encryption for the gateway. Only supported for AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret gateways. Valid values: true, false. Default value: false.
	// Enable encrypt gateway EBS volume. Only supported for AWS provider. Valid values: true, false. Default value: false.
	// +kubebuilder:validation:Optional
	EnableEncryptVolume *bool `json:"enableEncryptVolume,omitempty" tf:"enable_encrypt_volume,omitempty"`

	// Enable jumbo frames for this gateway. Default value is true.
	// Enable jumbo frame support for Gateway. Valid values: true or false. Default value: true.
	// +kubebuilder:validation:Optional
	EnableJumboFrame *bool `json:"enableJumboFrame,omitempty" tf:"enable_jumbo_frame,omitempty"`

	// Enable/disable LDAP. Valid values: true, false. Default value: false.
	// Specify whether to enable LDAP or not. Supported values: 'yes' and 'no'.
	// +kubebuilder:validation:Optional
	EnableLdap *bool `json:"enableLdap,omitempty" tf:"enable_ldap,omitempty"`

	// If set to true, the Monitor Gateway Subnets feature is enabled. Default value is false. Available in provider version R2.17.1+.
	// Enable monitor gateway subnets. Valid values: true, false. Default value: false.
	// +kubebuilder:validation:Optional
	EnableMonitorGatewaySubnets *bool `json:"enableMonitorGatewaySubnets,omitempty" tf:"enable_monitor_gateway_subnets,omitempty"`

	// Create a Public Subnet Filtering gateway. Valid values: true or false. Default value: false. Available as of provider version R2.18+.
	// Create a [Public Subnet Filtering gateway](https://docs.aviatrix.com/HowTos/public_subnet_filtering_faq.html).
	// +kubebuilder:validation:Optional
	EnablePublicSubnetFiltering *bool `json:"enablePublicSubnetFiltering,omitempty" tf:"enable_public_subnet_filtering,omitempty"`

	// Enable spot instance. NOT supported for production deployment.
	// Enable spot instance. NOT supported for production deployment.
	// +kubebuilder:validation:Optional
	EnableSpotInstance *bool `json:"enableSpotInstance,omitempty" tf:"enable_spot_instance,omitempty"`

	// Enable VPC DNS Server for gateway. Currently only supported for AWS, Azure, AzureGov, AWSGov, AWSChina, AzureChina, Alibaba Cloud, AWS Top Secret and AWS Secret gateways. Valid values: true, false. Default value: false.
	// Enable vpc_dns_server for Gateway. Valid values: true, false.
	// +kubebuilder:validation:Optional
	EnableVPCDNSServer *bool `json:"enableVpcDnsServer,omitempty" tf:"enable_vpc_dns_server,omitempty"`

	// Enable/disable VPN NAT. Only supported for VPN gateway. Valid values: true, false. Default value: true.
	// This field indicates whether to enable VPN NAT or not. Only supported for VPN gateway. Valid values: true, false. Default value: true.
	// +kubebuilder:validation:Optional
	EnableVPNNAT *bool `json:"enableVpnNat,omitempty" tf:"enable_vpn_nat,omitempty"`

	// Fault domain. Required and valid only for OCI. Available as of provider version R2.19.3.
	// Fault domain for OCI.
	// +kubebuilder:validation:Optional
	FaultDomain *string `json:"faultDomain,omitempty" tf:"fault_domain,omitempty"`

	// If fqdn_lan_cidr is set, the FQDN gateway will be created with an additional LAN interface using the provided CIDR. This attribute is required when enabling FQDN gateway FireNet in Azure or GCP. Available in provider version R2.17.1+.
	// FQDN gateway lan interface cidr.
	// +kubebuilder:validation:Optional
	FqdnLanCidr *string `json:"fqdnLanCidr,omitempty" tf:"fqdn_lan_cidr,omitempty"`

	// FQDN LAN VPC ID. This attribute is required when enabling FQDN gateway FireNet in GCP. Available as of provider version R2.18.1+.
	// LAN VPC ID. Only used for GCP FQDN Gateway.
	// +kubebuilder:validation:Optional
	FqdnLanVPCID *string `json:"fqdnLanVpcId,omitempty" tf:"fqdn_lan_vpc_id,omitempty"`

	// Size of the gateway instance. Example: AWS/AWSGov/AWSChina: "t2.large", GCP: "n1-standard-1", Azure/AzureGov/AzureChina: "Standard_B1s", OCI: "VM.Standard2.2".
	// Size of Gateway Instance.
	// +kubebuilder:validation:Required
	GwSize *string `json:"gwSize" tf:"gw_size,omitempty"`

	// It sets the value (seconds) of the idle timeout. This idle timeout feature is enable only if this attribute is set, otherwise it is disabled. The entered value must be an integer number greater than 300.  Available in provider version R2.17.1+.
	// Typed value when modifying idle_timeout. If it's -1, this feature is disabled.
	// +kubebuilder:validation:Optional
	IdleTimeout *float64 `json:"idleTimeout,omitempty" tf:"idle_timeout,omitempty"`

	// , please see notes here.
	// Enable Insane Mode for Gateway. Valid values: true, false.
	// +kubebuilder:validation:Optional
	InsaneMode *bool `json:"insaneMode,omitempty" tf:"insane_mode,omitempty"`

	// Region + Availability Zone of subnet being created for Insane Mode gateway. Required for AWS, AWSGov, AWS China, AWS Top Secret or AWS Secret if insane_mode is set. Example: AWS: "us-west-1a".
	// AZ of subnet being created for Insane Mode Gateway. Required if insane_mode is set.
	// +kubebuilder:validation:Optional
	InsaneModeAz *string `json:"insaneModeAz,omitempty" tf:"insane_mode_az,omitempty"`

	// LDAP base DN. Required if enable_ldap is true.
	// LDAP base DN. Required: Yes if enable_ldap is 'yes'.
	// +kubebuilder:validation:Optional
	LdapBaseDn *string `json:"ldapBaseDn,omitempty" tf:"ldap_base_dn,omitempty"`

	// LDAP bind DN. Required if enable_ldap is true.
	// LDAP bind DN. Required: Yes if enable_ldap is 'yes'.
	// +kubebuilder:validation:Optional
	LdapBindDn *string `json:"ldapBindDn,omitempty" tf:"ldap_bind_dn,omitempty"`

	// LDAP password. Required if enable_ldap is true.
	// LDAP password. Required: Yes if enable_ldap is 'yes'.
	// +kubebuilder:validation:Optional
	LdapPasswordSecretRef *v1.SecretKeySelector `json:"ldapPasswordSecretRef,omitempty" tf:"-"`

	// LDAP server address. Required if enable_ldap is true.
	// LDAP server address. Required: Yes if enable_ldap is 'yes'.
	// +kubebuilder:validation:Optional
	LdapServer *string `json:"ldapServer,omitempty" tf:"ldap_server,omitempty"`

	// LDAP user attribute. Required if enable_ldap is true.
	// LDAP user attribute. Required: Yes if enable_ldap is 'yes'.
	// +kubebuilder:validation:Optional
	LdapUsernameAttribute *string `json:"ldapUsernameAttribute,omitempty" tf:"ldap_username_attribute,omitempty"`

	// Maximum number of active VPN users allowed to be connected to this gateway. Required if vpn_access is true. Make sure the number is smaller than the VPN CIDR block. Example: 100. NOTE: Please see notes
	// Maximum connection of VPN access.
	// +kubebuilder:validation:Optional
	MaxVPNConn *string `json:"maxVpnConn,omitempty" tf:"max_vpn_conn,omitempty"`

	// Set of monitored instance ids. Only valid when 'enable_monitor_gateway_subnets' = true. Available in provider version R2.17.1+.
	// A set of monitored instance ids. Only valid when 'enable_monitor_gateway_subnets' = true.
	// +kubebuilder:validation:Optional
	MonitorExcludeList []*string `json:"monitorExcludeList,omitempty" tf:"monitor_exclude_list,omitempty"`

	// A list of DNS servers used to resolve domain names by a connected VPN user when Split Tunnel Mode is enabled.
	// A list of DNS servers used to resolve domain names by a connected VPN user when Split Tunnel Mode is enabled.
	// +kubebuilder:validation:Optional
	NameServers *string `json:"nameServers,omitempty" tf:"name_servers,omitempty"`

	// Token for Okta auth mode. Required if otp_mode is "3".
	// Token for Okta auth mode.
	// +kubebuilder:validation:Optional
	OktaTokenSecretRef *v1.SecretKeySelector `json:"oktaTokenSecretRef,omitempty" tf:"-"`

	// URL for Okta auth mode. Required if otp_mode is "3".
	// URL for Okta auth mode.
	// +kubebuilder:validation:Optional
	OktaURL *string `json:"oktaUrl,omitempty" tf:"okta_url,omitempty"`

	// Username suffix for Okta auth mode. Example: "aviatrix.com".
	// Username suffix for Okta auth mode.
	// +kubebuilder:validation:Optional
	OktaUsernameSuffix *string `json:"oktaUsernameSuffix,omitempty" tf:"okta_username_suffix,omitempty"`

	// Two step authentication mode. Valid values: "2" for DUO, "3" for Okta.
	// Two step authentication mode.
	// +kubebuilder:validation:Optional
	OtpMode *string `json:"otpMode,omitempty" tf:"otp_mode,omitempty"`

	// Peering HA gateway availability domain. Required and valid only for OCI. Available as of provider version R2.19.3.
	// Peering HA availability domain for OCI.
	// +kubebuilder:validation:Optional
	PeeringHaAvailabilityDomain *string `json:"peeringHaAvailabilityDomain,omitempty" tf:"peering_ha_availability_domain,omitempty"`

	// Name of public IP address resource and its resource group in Azure to be assigned to the HA peering instance. Example: "IP_Name:Resource_Group_Name". Required if peering_ha_eip is set and cloud_type is Azure, AzureGov or AzureChina. Available as of provider version 2.20+.
	// The name of the public IP address and its resource group in Azure to assign to the Peering HA Gateway.
	// +kubebuilder:validation:Optional
	PeeringHaAzureEIPNameResourceGroup *string `json:"peeringHaAzureEipNameResourceGroup,omitempty" tf:"peering_ha_azure_eip_name_resource_group,omitempty"`

	// Public IP address to be assigned to the HA peering instance. Only available for AWS, GCP, Azure, OCI, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret.
	// Public IP address that you want assigned to the HA peering instance.
	// +kubebuilder:validation:Optional
	PeeringHaEIP *string `json:"peeringHaEip,omitempty" tf:"peering_ha_eip,omitempty"`

	// Peering HA gateway fault domain. Required and valid only for OCI. Available as of provider version R2.19.3.
	// Peering HA fault domain for OCI.
	// +kubebuilder:validation:Optional
	PeeringHaFaultDomain *string `json:"peeringHaFaultDomain,omitempty" tf:"peering_ha_fault_domain,omitempty"`

	// Size of the Peering HA Gateway to be created. Required if enabling Peering HA. NOTE: Please see notes
	// Peering HA Gateway Size.
	// +kubebuilder:validation:Optional
	PeeringHaGwSize *string `json:"peeringHaGwSize,omitempty" tf:"peering_ha_gw_size,omitempty"`

	// The image version of the HA gateway. Use aviatrix_gateway_image data source to programmatically retrieve this value for the desired ha_software_version. If set, we will attempt to update the HA gateway to the specified version if current version is different. If left blank, the gateway upgrades can be managed with the aviatrix_controller_config resource. Type: String. Example: "hvm-cloudx-aws-022021". Available as of provider version R2.20.0.
	// peering_ha_image_version can be used to set the desired image version of the HA gateway. If set, we will attempt to update the gateway to the specified version.
	// +kubebuilder:validation:Optional
	PeeringHaImageVersion *string `json:"peeringHaImageVersion,omitempty" tf:"peering_ha_image_version,omitempty"`

	// Region + Availability Zone of subnet being created for Insane Mode-enabled Peering HA Gateway. Required for AWS only if insane_mode is set and peering_ha_subnet is set. Example: AWS: "us-west-1a".
	// AZ of subnet being created for Insane Mode Peering HA Gateway. Required if insane_mode is set.
	// +kubebuilder:validation:Optional
	PeeringHaInsaneModeAz *string `json:"peeringHaInsaneModeAz,omitempty" tf:"peering_ha_insane_mode_az,omitempty"`

	// The software version of the HA gateway. If set, we will attempt to update the HA gateway to the specified version if current version is different. If left blank, the HA gateway upgrade can be managed with the aviatrix_controller_config resource. Type: String. Example: "6.5.821". Available as of provider version R2.20.0.
	// peering_ha_software_version can be used to set the desired software version of the HA gateway. If set, we will attempt to update the gateway to the specified version. If left blank, the gateway software version will continue to be managed through the aviatrix_controller_config resource.
	// +kubebuilder:validation:Optional
	PeeringHaSoftwareVersion *string `json:"peeringHaSoftwareVersion,omitempty" tf:"peering_ha_software_version,omitempty"`

	// Public subnet CIDR to create Peering HA Gateway in. Required if enabling Peering HA for AWS/AWSGov/AWS Top Secret/AWS Secret/Azure/AzureGov/Alibaba Cloud. Optional if enabling Peering HA for GCP. Example: AWS: "10.0.0.0/16".
	// Public Subnet Information while creating Peering HA Gateway, only subnet is accepted. Required to create peering ha gateway if cloud_type = 1 or 8 (AWS or Azure). Optional if cloud_type = 4 (GCP)
	// +kubebuilder:validation:Optional
	PeeringHaSubnet *string `json:"peeringHaSubnet,omitempty" tf:"peering_ha_subnet,omitempty"`

	// Zone to create Peering HA Gateway in. Required if enabling Peering HA for GCP. Example: GCP: "us-west1-c". Optional for Azure. Valid values for Azure gateways are in the form "az-n". Example: "az-2". Available for Azure as of provider version R2.17+.
	// Zone information for creating Peering HA Gateway. Required to create peering ha gateway if cloud_type = 4 (GCP). Optional for cloud_type = 8 (Azure).
	// +kubebuilder:validation:Optional
	PeeringHaZone *string `json:"peeringHaZone,omitempty" tf:"peering_ha_zone,omitempty"`

	// Whether to enforce Guard Duty IP blocking.  Only valid when enable_public_subnet_filtering attribute is true. Valid values: true or false. Default value: true. Available as of provider version R2.18+.
	// Whether to enforce Guard Duty IP blocking. Required when `enable_public_subnet_filtering` attribute is true. Valid values: true or false. Default value: true.
	// +kubebuilder:validation:Optional
	PublicSubnetFilteringGuardDutyEnforced *bool `json:"publicSubnetFilteringGuardDutyEnforced,omitempty" tf:"public_subnet_filtering_guard_duty_enforced,omitempty"`

	// Route tables whose associated public subnets are protected for the HA PSF gateway. Required when enable_public_subnet_filtering and peering_ha_subnet are set. Available as of provider version R2.18+.
	// Route tables whose associated public subnets are protected for the HA PSF gateway. Required when enable_public_subnet_filtering and peering_ha_subnet are set.
	// +kubebuilder:validation:Optional
	PublicSubnetFilteringHaRouteTables []*string `json:"publicSubnetFilteringHaRouteTables,omitempty" tf:"public_subnet_filtering_ha_route_tables,omitempty"`

	// Route tables whose associated public subnets are protected. Only valid when enable_public_subnet_filtering attribute is true. Available as of provider version R2.18+.
	// Route tables whose associated public subnets are protected. Required when `enable_public_subnet_filtering` attribute is true.
	// +kubebuilder:validation:Optional
	PublicSubnetFilteringRouteTables []*string `json:"publicSubnetFilteringRouteTables,omitempty" tf:"public_subnet_filtering_route_tables,omitempty"`

	// It sets the value (seconds) of the renegotiation interval. This renegotiation interval feature is enable only if this attribute is set, otherwise it is disabled. The entered value must be an integer number greater than 300. Available in provider version R2.17.1+.
	// Typed value when modifying renegotiation_interval. If it's -1, this feature is disabled.
	// +kubebuilder:validation:Optional
	RenegotiationInterval *float64 `json:"renegotiationInterval,omitempty" tf:"renegotiation_interval,omitempty"`

	// Gateway ethernet interface RX queue size. Applies on HA as well if enabled. Once set, can't be deleted or disabled. Available for AWS as of provider version R2.22+.
	// Gateway ethernet interface RX queue size. Supported for AWS related clouds only. Applies on HA as well if enabled.
	// +kubebuilder:validation:Optional
	RxQueueSize *string `json:"rxQueueSize,omitempty" tf:"rx_queue_size,omitempty"`

	// Enable/disable SAML. This field is available in Controller version 3.3 or later release. Valid values: true, false. Default value: false.
	// This field indicates whether to enable SAML or not.
	// +kubebuilder:validation:Optional
	SAMLEnabled *bool `json:"samlEnabled,omitempty" tf:"saml_enabled,omitempty"`

	// A list of domain names that will use the NameServer when a specific name is not in the destination when Split Tunnel Mode is enabled.
	// A list of domain names that will use the NameServer when a specific name is not in the destination when Split Tunnel Mode is enabled.
	// +kubebuilder:validation:Optional
	SearchDomains *string `json:"searchDomains,omitempty" tf:"search_domains,omitempty"`

	// If enabled, Controller monitors the health of the gateway and restarts the gateway if it becomes unreachable. Valid values: true, false. Default value: false.
	// Set to true if this feature is desired.
	// +kubebuilder:validation:Optional
	SingleAzHa *bool `json:"singleAzHa,omitempty" tf:"single_az_ha,omitempty"`

	// Enable Source NAT in "single ip" mode for this gateway. Valid values: true, false. Default value: false. NOTE: If using SNAT for FQDN use-case, please see notes
	// Enable Source NAT for this container.
	// +kubebuilder:validation:Optional
	SingleIPSnat *bool `json:"singleIpSnat,omitempty" tf:"single_ip_snat,omitempty"`

	// Enable/disable Split Tunnel Mode. Valid values: true, false. Default value: true. Please see here for more information on split tunnel.
	// Specify split tunnel mode.
	// +kubebuilder:validation:Optional
	SplitTunnel *bool `json:"splitTunnel,omitempty" tf:"split_tunnel,omitempty"`

	// Price for spot instance. NOT supported for production deployment.
	// Price for spot instance. NOT supported for production deployment.
	// +kubebuilder:validation:Optional
	SpotPrice *string `json:"spotPrice,omitempty" tf:"spot_price,omitempty"`

	// A VPC network address range selected from one of the available network ranges. Example: "172.31.0.0/20". NOTE: If using
	// A VPC Network address range selected from one of the available network ranges.
	// +kubebuilder:validation:Required
	Subnet *string `json:"subnet" tf:"subnet,omitempty"`

	// Map of tags to assign to the gateway. Only available for AWS, AWSGov, AWSChina, Azure, AzureGov, AzureChina, AWS Top Secret and AWS Secret gateways. Allowed characters vary by cloud type but always include: letters, spaces, and numbers. AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret allow the use of any character.  Azure, AzureGov and AzureChina allows the following special characters: + - = . _ : @. Example: {"key1" = "value1", "key2" = "value2"}.
	// A map of tags to assign to the gateway.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The IPsec tunnel down detection time for the Gateway in seconds. Must be a number in the range [20-600]. The default value is set by the controller (60 seconds if nothing has been changed). NOTE: The controller UI has an option to set the tunnel detection time for all gateways. Available in provider R2.19+.
	// The IPSec tunnel down detection time for the Gateway.
	// +kubebuilder:validation:Optional
	TunnelDetectionTime *float64 `json:"tunnelDetectionTime,omitempty" tf:"tunnel_detection_time,omitempty"`

	// VPC ID/VNet name of cloud provider. Example: AWS/AWSGov/AWSChina: "vpc-abcd1234", GCP: "vpc-gcp-test~-~project-id", Azure/AzureGov/AzureChina: "vnet_name:rg_name:resource_guid", OCI: "ocid1.vcn.oc1.iad.aaaaaaaaba3pv6wkcr4jqae5f44n2b2m2yt2j6rx32uzr4h25vqstifsfdsq".
	// ID of legacy VPC/Vnet to be connected.
	// +kubebuilder:validation:Required
	VPCID *string `json:"vpcId" tf:"vpc_id,omitempty"`

	// VPC region the gateway will be created in. Example: AWS: "us-east-1", GCP: "us-west2-a", Azure: "East US 2", OCI: "us-ashburn-1", AzureGov: "USGov Arizona", AWSGov: "us-gov-west-1", AWSChina: "cn-north-1", AzureChina: "China North", AWS Top Secret: "us-iso-east-1", AWS Secret: "us-isob-east-1".
	// Region where this gateway will be launched.
	// +kubebuilder:validation:Required
	VPCReg *string `json:"vpcReg" tf:"vpc_reg,omitempty"`

	// Enable user access through VPN to this gateway. Valid values: true, false.
	// Enable user access through VPN to this container.
	// +kubebuilder:validation:Optional
	VPNAccess *bool `json:"vpnAccess,omitempty" tf:"vpn_access,omitempty"`

	// VPN CIDR block for the gateway. Required if vpn_access is true. Example: "192.168.43.0/24".
	// VPN CIDR block for the container.
	// +kubebuilder:validation:Optional
	VPNCidr *string `json:"vpnCidr,omitempty" tf:"vpn_cidr,omitempty"`

	// Transport mode for VPN connection. All cloud_types support TCP with ELB, and UDP without ELB. AWS(1) additionally supports UDP with ELB. Valid values: "TCP", "UDP". If not specified, "TCP" will be used.
	// Elb protocol for VPN gateway with elb enabled. Only supports AWS provider. Valid values: 'TCP', 'UDP'. If not specified, 'TCP” will be used.
	// +kubebuilder:validation:Optional
	VPNProtocol *string `json:"vpnProtocol,omitempty" tf:"vpn_protocol,omitempty"`

	// Availability Zone. Only available for Azure (8), Azure GOV (32), Azure CHINA (2048) and Public Subnet Filtering gateway. Available for Azure as of provider version R2.17+.
	// Availability Zone. Only available for Azure (8), Azure GOV (32), Azure CHINA (2048) and Public Subnet Filtering gateway. Must be in the form 'az-n', for example, 'az-2'.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// GatewaySpec defines the desired state of Gateway
type GatewaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GatewayParameters `json:"forProvider"`
}

// GatewayStatus defines the observed state of Gateway.
type GatewayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GatewayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Gateway is the Schema for the Gateways API. Creates and manages Aviatrix gateways
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aviatrix}
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GatewaySpec   `json:"spec"`
	Status            GatewayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GatewayList contains a list of Gateways
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// Repository type metadata.
var (
	Gateway_Kind             = "Gateway"
	Gateway_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Gateway_Kind}.String()
	Gateway_KindAPIVersion   = Gateway_Kind + "." + CRDGroupVersion.String()
	Gateway_GroupVersionKind = CRDGroupVersion.WithKind(Gateway_Kind)
)

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
}
