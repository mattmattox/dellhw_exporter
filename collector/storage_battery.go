/*
Copyright 2021 The dellhw_exporter Authors. All rights reserved.

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

package collector

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

type storageBatteryCollector struct {
	current *prometheus.Desc
}

func init() {
	Factories["storage_battery"] = NewStorageBatteryCollector
}

// NewStorageBatteryCollector returns a new storageBatteryCollector
func NewStorageBatteryCollector(cfg *Config) (Collector, error) {
	return &storageBatteryCollector{}, nil
}

// Update Prometheus metrics
func (c *storageBatteryCollector) Update(ch chan<- prometheus.Metric) error {
	storageBattery, err := or.StorageBattery()
	if err != nil {
		return err
	}
	for _, value := range storageBattery {
		float, err := strconv.ParseFloat(value.Value, 64)
		if err != nil {
			return err
		}
		c.current = prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", value.Name),
			"Status of storage controller backup batteries.",
			nil, value.Labels)
		ch <- prometheus.MustNewConstMetric(
			c.current, prometheus.GaugeValue, float)
	}

	return nil
}
