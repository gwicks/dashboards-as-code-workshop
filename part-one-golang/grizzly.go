package main

import "github.com/grafana/grafana-foundation-sdk/go/dashboard"

type Manifest struct {
	APIVersion string            `json:"apiVersion" yaml:"apiVersion"`
	Kind       string            `json:"kind" yaml:"kind"`
	Metadata   map[string]string `json:"metadata" yaml:"metadata"`
	Spec       any               `json:"spec" yaml:"spec"`
}

func DashboardManifestFrom(folderUid string, dash dashboard.Dashboard) Manifest {
	return Manifest{
		APIVersion: "grizzly.grafana.com/v1alpha1",
		Kind:       "Dashboard",
		Metadata: map[string]string{
			"folder": folderUid,
			"name":   *dash.Uid,
		},
		Spec: dash,
	}
}
