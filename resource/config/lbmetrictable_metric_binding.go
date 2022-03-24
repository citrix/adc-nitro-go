/*
* Copyright (c) 2021 Citrix Systems, Inc.
*
*   Licensed under the Apache License, Version 2.0 (the "License");
*   you may not use this file except in compliance with the License.
*   You may obtain a copy of the License at
*
*       http://www.apache.org/licenses/LICENSE-2.0
*
*  Unless required by applicable law or agreed to in writing, software
*   distributed under the License is distributed on an "AS IS" BASIS,
*   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*   See the License for the specific language governing permissions and
*   limitations under the License.
 */

package config

/**
* Binding class showing the metric that can be bound to lbmetrictable.
 */
type Lbmetrictablemetricbinding struct {
	/**
	* Name of the metric for which to change the SNMP OID.
	 */
	Metric string `json:"metric,omitempty"`
	/**
	* New SNMP OID of the metric.
	 */
	Snmpoid string `json:"Snmpoid,omitempty"`
	/**
	* Indication if it is a configured or internal
	 */
	Metrictype string `json:"metrictype,omitempty"`
	/**
	* Name of the metric table.
	 */
	Metrictable string `json:"metrictable,omitempty"`
}