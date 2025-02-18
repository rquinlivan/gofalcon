// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainDiscoverAPIHost Represents information about a managed, an unmanaged or an unsupported asset.
//
// swagger:model domain.DiscoverAPIHost
type DomainDiscoverAPIHost struct {

	// Whether the asset is account-enabled in Active Directory (Yes or No).
	AccountEnabled string `json:"account_enabled,omitempty"`

	// The user account control properties in Active Directory.
	AdUserAccountControl int32 `json:"ad_user_account_control,omitempty"`

	// The version of the Falcon sensor that's installed on the asset.
	AgentVersion string `json:"agent_version,omitempty"`

	// The agent ID of the Falcon sensor installed on the asset.
	Aid string `json:"aid,omitempty"`

	// The first and last name of the person who is assigned to this asset.
	AssignedTo string `json:"assigned_to,omitempty"`

	// The available disk space in the last 15 minutes on the host
	AvailableDiskSpace int32 `json:"available_disk_space,omitempty"`

	// The available disk space percent in the last 15 minutes on the host
	AvailableDiskSpacePct int32 `json:"available_disk_space_pct,omitempty"`

	// The average memory usage in the last 15 minutes on the host
	AverageMemoryUsage int32 `json:"average_memory_usage,omitempty"`

	// The average memory usage percent in the last 15 minutes on the host
	AverageMemoryUsagePct int32 `json:"average_memory_usage_pct,omitempty"`

	// The average processor usage in the last 15 minutes on the host
	AverageProcessorUsage int32 `json:"average_processor_usage,omitempty"`

	// The list of found sha256 and their measurement types
	BiosHashesData []*DomainDiscoverAPIBiosHashesData `json:"bios_hashes_data"`

	// The id of the bios on the host
	BiosID string `json:"bios_id,omitempty"`

	// The name of the asset's BIOS manufacturer.
	BiosManufacturer string `json:"bios_manufacturer,omitempty"`

	// The asset's BIOS version.
	BiosVersion string `json:"bios_version,omitempty"`

	// The asset's customer ID.
	// Required: true
	Cid *string `json:"cid"`

	// The name of the city where the asset is located.
	City string `json:"city,omitempty"`

	// How the server is classified, such as production, development, disaster recovery, or user acceptance testing.
	Classification string `json:"classification,omitempty"`

	// The level of confidence that the asset is a corporate asset (25 = low confidence, 50 = medium confidence, 75 = high confidence).
	Confidence int32 `json:"confidence,omitempty"`

	// The name of the country where the asset is located.
	Country string `json:"country,omitempty"`

	// The manufacturer of the asset's CPU.
	CPUManufacturer string `json:"cpu_manufacturer,omitempty"`

	// The name of the processor on the system
	CPUProcessorName string `json:"cpu_processor_name,omitempty"`

	// The time the asset was created in Active Directory, according to LDAP info.
	CreationTimestamp string `json:"creation_timestamp,omitempty"`

	// The last seen local IPv4 address of the asset.
	CurrentLocalIP string `json:"current_local_ip,omitempty"`

	// Where the data about the asset came from (such as CrowdStrike, ServiceNow, or Active Directory).
	DataProviders []string `json:"data_providers"`

	// How many services provided data about the asset.
	DataProvidersCount int32 `json:"data_providers_count,omitempty"`

	// The department where the asset is used.
	Department string `json:"department,omitempty"`

	// The descriptions of the asset in Active Directory (Cannot be used for filtering, sorting, or querying).
	Descriptions []string `json:"descriptions"`

	// The agent IDs of the Falcon sensors installed on the sources that discovered the asset.
	DiscovererAids []string `json:"discoverer_aids"`

	// The number of sources that discovered the asset.
	DiscovererCount int32 `json:"discoverer_count,omitempty"`

	// The platform names of the sources that discovered the asset.
	DiscovererPlatformNames []string `json:"discoverer_platform_names"`

	// The product type descriptions of the sources that discovered the asset.
	DiscovererProductTypeDescs []string `json:"discoverer_product_type_descs"`

	// The tags of the sources that discovered the asset.
	DiscovererTags []string `json:"discoverer_tags"`

	// The names and sizes of the disks on the asset
	DiskSizes []*DomainDiscoverAPIDiskSize `json:"disk_sizes"`

	// The email of the asset as listed in Active Directory.
	Email string `json:"email,omitempty"`

	// The list of encrypted drives on the host
	EncryptedDrives []string `json:"encrypted_drives"`

	// The count of encrypted drives on the host
	EncryptedDrivesCount int32 `json:"encrypted_drives_count,omitempty"`

	// The encryption status of the host
	EncryptionStatus string `json:"encryption_status,omitempty"`

	// The type of asset (managed, unmanaged, unsupported).
	EntityType string `json:"entity_type,omitempty"`

	// The external IPv4 address of the asset.
	ExternalIP string `json:"external_ip,omitempty"`

	// Lists the data providers for each property in the response (Cannot be used for filtering, sorting, or querying).
	FieldMetadata map[string]DomainDiscoverAPIFieldMetadata `json:"field_metadata,omitempty"`

	// The agent ID of the Falcon sensor on the source that first discovered the asset.
	FirstDiscovererAid string `json:"first_discoverer_aid,omitempty"`

	// The first time the asset was seen in your environment.
	FirstSeenTimestamp string `json:"first_seen_timestamp,omitempty"`

	// The fully qualified domain name of the asset.
	Fqdn string `json:"fqdn,omitempty"`

	// The host management groups the asset is part of.
	Groups []string `json:"groups"`

	// The asset's hostname.
	Hostname string `json:"hostname,omitempty"`

	// The unique ID of the asset.
	// Required: true
	ID *string `json:"id"`

	// Whether the asset is exposed to the internet (Yes or Unknown).
	InternetExposure string `json:"internet_exposure,omitempty"`

	// For Linux and Mac hosts: the major version, minor version, and patch version of the kernel for the asset. For Windows hosts: the build number of the asset.
	KernelVersion string `json:"kernel_version,omitempty"`

	// The agent ID of the Falcon sensor installed on the source that most recently discovered the asset.
	LastDiscovererAid string `json:"last_discoverer_aid,omitempty"`

	// The most recent time the asset was seen in your environment.
	LastSeenTimestamp string `json:"last_seen_timestamp,omitempty"`

	// Historical local IPv4 addresses associated with the asset.
	LocalIPAddresses []string `json:"local_ip_addresses"`

	// The number of historical local IPv4 addresses the asset has had.
	LocalIpsCount int32 `json:"local_ips_count,omitempty"`

	// The location of the asset.
	Location string `json:"location,omitempty"`

	// The number of logical cores available on the system
	LogicalCoreCount int32 `json:"logical_core_count,omitempty"`

	// Historical MAC addresses associated with the asset.
	MacAddresses []string `json:"mac_addresses"`

	// The domain name the asset is currently joined to.
	MachineDomain string `json:"machine_domain,omitempty"`

	// The first and last name of the person who manages this asset.
	ManagedBy string `json:"managed_by,omitempty"`

	// The max memory usage in the last 15 minutes on the host
	MaxMemoryUsage int32 `json:"max_memory_usage,omitempty"`

	// The max memory usage percent in the last 15 minutes on the host
	MaxMemoryUsagePct int32 `json:"max_memory_usage_pct,omitempty"`

	// The max processor usage in the last 15 minutes on the host
	MaxProcessorUsage int32 `json:"max_processor_usage,omitempty"`

	// The path, used and available space on mounted disks
	MountStorageInfo []*DomainDiscoverAPIMountStorageInfo `json:"mount_storage_info"`

	// The asset's network interfaces (Cannot be used for filtering, sorting, or querying).
	NetworkInterfaces []*DomainDiscoverAPINetworkInterface `json:"network_interfaces"`

	// The number of active physical drives available on the system.
	NumberOfDiskDrives int32 `json:"number_of_disk_drives,omitempty"`

	// The globally unique identifier (GUID) of the asset in Active Directory.
	ObjectGUID string `json:"object_guid,omitempty"`

	// The security identifier (SID) of the asset in Active Directory.
	ObjectSid string `json:"object_sid,omitempty"`

	// Whether the asset is at end of support (Yes, No, or Unknown).
	OsIsEol string `json:"os_is_eol,omitempty"`

	// The os security features of the asset
	OsSecurity *DomainDiscoverAPIOsSecurity `json:"os_security,omitempty"`

	// The OS service pack on the asset.
	OsServicePack string `json:"os_service_pack,omitempty"`

	// The OS version of the asset.
	OsVersion string `json:"os_version,omitempty"`

	// The organizational unit of the asset.
	Ou string `json:"ou,omitempty"`

	// The first and last name of the person who owns this asset.
	OwnedBy string `json:"owned_by,omitempty"`

	// The number of physical CPU cores available on the system.
	PhysicalCoreCount int32 `json:"physical_core_count,omitempty"`

	// The platform name of the asset (Windows, Mac, Linux).
	PlatformName string `json:"platform_name,omitempty"`

	// The number of physical processors available on the system.
	ProcessorPackageCount int32 `json:"processor_package_count,omitempty"`

	// The product type of the asset represented as a number (1 = Workstation, 2 = Domain Controller, 3 = Server).
	ProductType string `json:"product_type,omitempty"`

	// The product type of the asset (Workstation, Domain Controller, Server).
	ProductTypeDesc string `json:"product_type_desc,omitempty"`

	// Whether the asset is in reduced functionality mode (Yes or No).
	ReducedFunctionalityMode string `json:"reduced_functionality_mode,omitempty"`

	// The unique identifier of the asset from ServiceNow, if any.
	ServicenowID string `json:"servicenow_id,omitempty"`

	// The site name of the domain the asset is joined to (applies only to Windows hosts).
	SiteName string `json:"site_name,omitempty"`

	// The name of the U.S. state where the asset is located.
	State string `json:"state,omitempty"`

	// The asset's system manufacturer.
	SystemManufacturer string `json:"system_manufacturer,omitempty"`

	// The asset's system product name.
	SystemProductName string `json:"system_product_name,omitempty"`

	// The asset's system serial number.
	SystemSerialNumber string `json:"system_serial_number,omitempty"`

	// The sensor and cloud tags of the asset.
	Tags []string `json:"tags"`

	// The count of bios files measured by the firmware image
	TotalBiosFiles int32 `json:"total_bios_files,omitempty"`

	// Total amount of disk space available on the system
	TotalDiskSpace int32 `json:"total_disk_space,omitempty"`

	// The total memory of the asset
	TotalMemory int32 `json:"total_memory,omitempty"`

	// The list of unencrypted drives on the host
	UnencryptedDrives []string `json:"unencrypted_drives"`

	// The count of unencrypted drives on the host
	UnencryptedDrivesCount int32 `json:"unencrypted_drives_count,omitempty"`

	// The used disk space in the last 15 minutes on the host
	UsedDiskSpace int32 `json:"used_disk_space,omitempty"`

	// The used disk space percent in the last 15 minutes on the host
	UsedDiskSpacePct int32 `json:"used_disk_space_pct,omitempty"`

	// What the asset is used for, such as production, staging, or QA.
	UsedFor string `json:"used_for,omitempty"`
}

// Validate validates this domain discover API host
func (m *DomainDiscoverAPIHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBiosHashesData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDiskSizes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFieldMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMountStorageInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsSecurity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDiscoverAPIHost) validateBiosHashesData(formats strfmt.Registry) error {
	if swag.IsZero(m.BiosHashesData) { // not required
		return nil
	}

	for i := 0; i < len(m.BiosHashesData); i++ {
		if swag.IsZero(m.BiosHashesData[i]) { // not required
			continue
		}

		if m.BiosHashesData[i] != nil {
			if err := m.BiosHashesData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bios_hashes_data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bios_hashes_data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainDiscoverAPIHost) validateCid(formats strfmt.Registry) error {

	if err := validate.Required("cid", "body", m.Cid); err != nil {
		return err
	}

	return nil
}

func (m *DomainDiscoverAPIHost) validateDiskSizes(formats strfmt.Registry) error {
	if swag.IsZero(m.DiskSizes) { // not required
		return nil
	}

	for i := 0; i < len(m.DiskSizes); i++ {
		if swag.IsZero(m.DiskSizes[i]) { // not required
			continue
		}

		if m.DiskSizes[i] != nil {
			if err := m.DiskSizes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disk_sizes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disk_sizes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainDiscoverAPIHost) validateFieldMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.FieldMetadata) { // not required
		return nil
	}

	for k := range m.FieldMetadata {

		if err := validate.Required("field_metadata"+"."+k, "body", m.FieldMetadata[k]); err != nil {
			return err
		}
		if val, ok := m.FieldMetadata[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("field_metadata" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("field_metadata" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainDiscoverAPIHost) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *DomainDiscoverAPIHost) validateMountStorageInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.MountStorageInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.MountStorageInfo); i++ {
		if swag.IsZero(m.MountStorageInfo[i]) { // not required
			continue
		}

		if m.MountStorageInfo[i] != nil {
			if err := m.MountStorageInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mount_storage_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mount_storage_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainDiscoverAPIHost) validateNetworkInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.NetworkInterfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkInterfaces); i++ {
		if swag.IsZero(m.NetworkInterfaces[i]) { // not required
			continue
		}

		if m.NetworkInterfaces[i] != nil {
			if err := m.NetworkInterfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("network_interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("network_interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainDiscoverAPIHost) validateOsSecurity(formats strfmt.Registry) error {
	if swag.IsZero(m.OsSecurity) { // not required
		return nil
	}

	if m.OsSecurity != nil {
		if err := m.OsSecurity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("os_security")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("os_security")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain discover API host based on the context it is used
func (m *DomainDiscoverAPIHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBiosHashesData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDiskSizes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFieldMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMountStorageInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworkInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOsSecurity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainDiscoverAPIHost) contextValidateBiosHashesData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BiosHashesData); i++ {

		if m.BiosHashesData[i] != nil {
			if err := m.BiosHashesData[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("bios_hashes_data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("bios_hashes_data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainDiscoverAPIHost) contextValidateDiskSizes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DiskSizes); i++ {

		if m.DiskSizes[i] != nil {
			if err := m.DiskSizes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disk_sizes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disk_sizes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainDiscoverAPIHost) contextValidateFieldMetadata(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.FieldMetadata {

		if val, ok := m.FieldMetadata[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *DomainDiscoverAPIHost) contextValidateMountStorageInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MountStorageInfo); i++ {

		if m.MountStorageInfo[i] != nil {
			if err := m.MountStorageInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mount_storage_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("mount_storage_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainDiscoverAPIHost) contextValidateNetworkInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NetworkInterfaces); i++ {

		if m.NetworkInterfaces[i] != nil {
			if err := m.NetworkInterfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("network_interfaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("network_interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DomainDiscoverAPIHost) contextValidateOsSecurity(ctx context.Context, formats strfmt.Registry) error {

	if m.OsSecurity != nil {
		if err := m.OsSecurity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("os_security")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("os_security")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainDiscoverAPIHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainDiscoverAPIHost) UnmarshalBinary(b []byte) error {
	var res DomainDiscoverAPIHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
