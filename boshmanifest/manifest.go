package boshmanifest

import "encoding/json"

type manifest struct {
	Name          string            `json:"name"`
	DirectorUuid  string            `json:"director_uuid"`
	Release       release           `json:"release"`
	Compilation   compilation       `json:"compilation"`
	Update        update            `json:"update"`
	Networks      []network         `json:"networks"`
	Properties    map[string]string `json:"properties"`
	ResourcePools []resourcePool    `json:"resource_pools"`
	Jobs          []job             `json:"jobs"`
}

type release struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type compilation struct {
	Workers         int             `json:"workers"`
	Network         string          `json:"network"`
	CloudProperties cloudProperties `json:"cloud_properties"`
}

type cloudProperties struct {
	Ram  int `json:"ram"`
	Disk int `json:"disk"`
	Cpu  int `json:"cpu"`
}

type update struct {
	Canaries        int    `json:"canaries"`
	CanaryWatchTime string `json:"canary_watch_time"`
	UpdateWatchTime string `json:"update_watch_time"`
	MaxInFlight     int    `json:"max_in_flight"`
	MaxErrors       int    `json:"max_errors"`
}

type network struct {
	Name    string   `json:"name"`
	Subnets []subnet `json:"subnets"`
}

type resourcePool struct {
	Name            string          `json:"name"`
	Network         string          `json:"network"`
	Size            int             `json:"size"`
	Stemcell        stemcell        `json:"stemcell"`
	CloudProperties cloudProperties `json:"cloud_properties"`
}

type stemcell struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type job struct {
	Name         string       `json:"name"`
	Template     string       `json:"template"`
	Instances    int          `json:"instances"`
	ResourcePool string       `json:"resource_pool"`
	Networks     []jobNetwork `json:"networks"`
}

type jobNetwork struct {
	Name string `json:"name"`
}

func Build(input InputFields) (output string, err error) {
	release := release{"dummy", "latest"}
	cloudProperties := cloudProperties{512, 1024, 1}
	compilation := compilation{1, "default", cloudProperties}
	update := update{1, "3000-90000", "3000-90000", 4, 1}

	subnet1, err := newSubnet(input.Cidr, []string{"8.8.8.8"}, "VM Network")
	if err != nil {
		return
	}
	subnets := []subnet{*subnet1}

	networks := []network{{"default", subnets}}

	stemcell := stemcell{"bosh-vsphere-esxi-ubuntu", "latest"}

	resourcePools := []resourcePool{{
		"default",
		"default",
		1,
		stemcell,
		cloudProperties,
	}}

	jobs := []job{
		{"dummy", "dummy", 1, "default", []jobNetwork{{"default"}}},
	}

	manifest := manifest{
		"dummy",
		input.DirectorUuid,
		release,
		compilation,
		update,
		networks,
		make(map[string]string),
		resourcePools,
		jobs,
	}

	var bytes []byte
	bytes, err = json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return
	}

	output = string(bytes)

	return
}
