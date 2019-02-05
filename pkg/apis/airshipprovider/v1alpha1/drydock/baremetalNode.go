// Code generated by schema-generate. DO NOT EDIT.

package drydock

// AddressingItems
type AddressingItems struct {
	Address string `json:"address,omitempty"`
	Network string `json:"network,omitempty"`
}

// Filesystem
type Filesystem struct {
	FsLabel      string `json:"fs_label,omitempty"`
	FsUuid       string `json:"fs_uuid,omitempty"`
	Fstype       string `json:"fstype,omitempty"`
	MountOptions string `json:"mount_options,omitempty"`
	Mountpoint   string `json:"mountpoint,omitempty"`
}

// InterfacesItem
type InterfacesItem struct {
	DeviceLink string   `json:"device_link,omitempty"`
	Networks   []string `json:"networks,omitempty"`
	Slaves     []string `json:"slaves,omitempty"`
}

// KernelParams
type KernelParams struct {
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
}

// LogicalVolumesItems
type LogicalVolumesItems struct {
	Filesystem *Filesystem `json:"filesystem,omitempty"`
	LvUuid     string      `json:"lv_uuid,omitempty"`
	Name       string      `json:"name,omitempty"`
	Size       string      `json:"size,omitempty"`
}

// Metadata
type Metadata struct {
	BootMac   string            `json:"boot_mac,omitempty"`
	OwnerData map[string]string `json:"owner_data,omitempty"`
	Rack      string            `json:"rack,omitempty"`
	Tags      []string          `json:"tags,omitempty"`
}

// Oob
type Oob struct {
	Account              string                 `json:"account,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-,omitempty"`
	Credetial            string                 `json:"credetial,omitempty"`
	Network              string                 `json:"network,omitempty"`
	Type                 string                 `json:"type,omitempty"`
}

// PartitionsItems
type PartitionsItems struct {
	Bootable    bool              `json:"bootable,omitempty"`
	Filesystem  *Filesystem       `json:"filesystem,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
	Name        string            `json:"name,omitempty"`
	PartUuid    string            `json:"part_uuid,omitempty"`
	Size        string            `json:"size,omitempty"`
	VolumeGroup string            `json:"volume_group,omitempty"`
}

// PhysicalDevicesItem
type PhysicalDevicesItem struct {
	Labels      map[string]string  `json:"labels,omitempty"`
	Partitions  []*PartitionsItems `json:"partitions,omitempty"`
	VolumeGroup string             `json:"volume_group,omitempty"`
}

// Platform
type Platform struct {
	Image        string        `json:"image,omitempty"`
	Kernel       string        `json:"kernel,omitempty"`
	KernelParams *KernelParams `json:"kernel_params,omitempty"`
}

// Root
type BaremetalSpec struct {
	Addressing      []*AddressingItems         `json:"addressing,omitempty"`
	HardwareProfile string                     `json:"hardware_profile,omitempty"`
	HostProfile     string                     `json:"host_profile,omitempty"`
	Interfaces      map[string]*InterfacesItem `json:"interfaces,omitempty"`
	Metadata        *Metadata                  `json:"metadata,omitempty"`
	Oob             *Oob                       `json:"oob,omitempty"`
	Platform        *Platform                  `json:"platform,omitempty"`
	PrimaryNetwork  string                     `json:"primary_network,omitempty"`
	Storage         *Storage                   `json:"storage,omitempty"`
}

// Storage
type Storage struct {
	PhysicalDevices map[string]*PhysicalDevicesItem `json:"physical_devices,omitempty"`
	VolumeGroups    map[string]*VolumeGroupsItem    `json:"volume_groups,omitempty"`
}

// VolumeGroupsItem
type VolumeGroupsItem struct {
	LogicalVolumes []*LogicalVolumesItems `json:"logical_volumes,omitempty"`
	VgUuid         string                 `json:"vg_uuid,omitempty"`
}
