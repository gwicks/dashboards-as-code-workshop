package main

import (
	"fmt"

	"github.com/grafana/grafana-foundation-sdk/go/common"
	"github.com/grafana/grafana-foundation-sdk/go/stat"
)

func versionStat(service Service) *stat.PanelBuilder {
	return statPanel().
		Title("Version").
		WithTarget(
			instantPrometheusQuery(fmt.Sprintf("app_infos{service=\"%s\"}", service.Name)),
		).
		Datasource(prometheusDatasourceRef()).
		ReduceOptions(common.NewReduceDataOptionsBuilder().
			Values(false).
			Calcs([]string{"last"}).
			Fields("/^version$/"),
		)
}
