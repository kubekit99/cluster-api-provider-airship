// Code generated by schema-generate. DO NOT EDIT.

package drydock

// Filesystem
type Filesystem2 struct {
	FsLabel      string `json:"fs_label,omitempty"`
	FsUuid       string `json:"fs_uuid,omitempty"`
	Fstype       string `json:"fstype,omitempty"`
	MountOptions string `json:"mount_options,omitempty"`
	Mountpoint   string `json:"mountpoint,omitempty"`
}

// InterfacesItem
type InterfacesItem2 struct {
	DeviceLink string   `json:"device_link,omitempty"`
	Networks   []string `json:"networks,omitempty"`
	Slaves     []string `json:"slaves,omitempty"`
	Sriov      *Sriov   `json:"sriov,omitempty"`
}

// KernelParams
type KernelParams2 struct {
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}

// LogicalVolumesItems
type LogicalVolumesItems2 struct {
	Filesystem *Filesystem `json:"filesystem,omitempty"`
	LvUuid     string      `json:"lv_uuid,omitempty"`
	Name       string      `json:"name,omitempty"`
	Size       string      `json:"size,omitempty"`
}

// Metadata
type Metadata2 struct {
	BootMac   string            `json:"boot_mac,omitempty"`
	OwnerData map[string]string `json:"owner_data,omitempty"`
	Rack      string            `json:"rack,omitempty"`
	Tags      []string          `json:"tags,omitempty"`
}

// Oob
type Oob2 struct {
	Account              string                 `json:"account,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
	Credetial            string                 `json:"credetial,omitempty"`
	Network              string                 `json:"network,omitempty"`
	Type                 string                 `json:"type,omitempty"`
}

// PartitionsItems
type PartitionsItems2 struct {
	Bootable    bool              `json:"bootable,omitempty"`
	Filesystem  *Filesystem2      `json:"filesystem,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Name        string            `json:"name,omitempty"`
	PartUuid    string            `json:"part_uuid,omitempty"`
	Size        string            `json:"size,omitempty"`
	VolumeGroup string            `json:"volume_group,omitempty"`
}

// PhysicalDevicesItem
type PhysicalDevicesItem2 struct {
	Labels      map[string]string   `json:"labels,omitempty"`
	Partitions  []*PartitionsItems2 `json:"partitions,omitempty"`
	VolumeGroup string              `json:"volume_group,omitempty"`
}

// Platform
type Platform2 struct {
	Image        string         `json:"image,omitempty"`
	Kernel       string         `json:"kernel,omitempty"`
	KernelParams *KernelParams2 `json:"kernel_params,omitempty"`
}

// Root
type HostProfileSpec struct {
	HardwareProfile string                     `json:"hardware_profile,omitempty"`
	HostProfile     string                     `json:"host_profile,omitempty"`
	Interfaces      map[string]*InterfacesItem `json:"interfaces,omitempty"`
	Metadata        *Metadata                  `json:"metadata,omitempty"`
	Oob             *Oob                       `json:"oob,omitempty"`
	Platform        *Platform                  `json:"platform,omitempty"`
	PrimaryNetwork  string                     `json:"primary_network,omitempty"`
	Storage         *Storage                   `json:"storage,omitempty"`
}

// Sriov
type Sriov struct {
	Trustmode bool    `json:"trustmode,omitempty"`
	VfCount   float64 `json:"vf_count,omitempty"`
}

// Storage
type Storage2 struct {
	PhysicalDevices map[string]*PhysicalDevicesItem2 `json:"physical_devices,omitempty"`
	VolumeGroups    map[string]*VolumeGroupsItem2    `json:"volume_groups,omitempty"`
}

// VolumeGroupsItem
type VolumeGroupsItem2 struct {
	LogicalVolumes []*LogicalVolumesItems2 `json:"logical_volumes,omitempty"`
	VgUuid         string                  `json:"vg_uuid,omitempty"`
}
