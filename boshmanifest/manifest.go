package boshmanifest

import "encoding/json"

type manifest struct {
	Name         string      `json:"name"`
	DirectorUuid string      `json:"director_uuid"`
	Release      release     `json:"release"`
	Compilation  compilation `json:"compilation"`
	Update       update      `json:"update"`
	Networks     []network   `json:"networks"`
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

type subnet struct {
	Range           string                 `json:"range"`
	Reserved        []string               `json:"reserved"`
	Static          []string               `json:"static"`
	Gateway         string                 `json:"gateway"`
	Dns             []string               `json:"dns"`
	CloudProperties networkCloudProperties `json:"cloud_properties"`
}

type networkCloudProperties struct {
	Name string `json:"name"`
}

func Build(input InputFields) (output string, err error) {
	release := release{"dummy", "latest"}
	cloudProperties := cloudProperties{512, 1024, 1}
	compilation := compilation{1, "default", cloudProperties}
	update := update{1, "3000-90000", "3000-90000", 4, 1}
	reserved := []string{"192.168.0.2 - 192.168.0.10"}
	static := []string{"192.168.0.11"}
	dns := []string{"8.8.8.8"}
	networkCloudProperties := networkCloudProperties{"default"}

	subnets := []subnet{{
		"192.68.0.1/24",
		reserved,
		static,
		"192.168.0.1",
		dns,
		networkCloudProperties,
	}}

	networks := []network{{"default", subnets}}

	manifest := manifest{
		"dummy",
		input.DirectorUuid,
		release,
		compilation,
		update,
		networks,
	}

	var bytes []byte
	bytes, err = json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return
	}

	output = string(bytes)

	return
}
