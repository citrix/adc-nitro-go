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

package cs

/**
* Binding class showing the csvserver that can be bound to cspolicy.
 */
type Cspolicycsvserverbinding struct {
	/**
	* The domain name. The string value can range to 63 characters.
	 */
	Domain string `json:"domain,omitempty"`
	/**
	* Content switching action that names the target load balancing virtual server to which the traffic is switched.
	 */
	Action string `json:"action,omitempty"`
	/**
	* URL string that is matched with the URL of a request. Can contain a wildcard character. Specify the string value in the following format: [[prefix] [*]] [.suffix].
	 */
	Url string `json:"url,omitempty"`
	/**
	* priority of bound policy
	 */
	Priority int `json:"priority,omitempty"`
	/**
	* Total number of hits.
	 */
	Hits int `json:"hits,omitempty"`
	/**
	* Total number of hits.
	 */
	Pihits int `json:"pihits,omitempty"`
	/**
	* bind hits for PI CS Policy.
	 */
	Pipolicyhits int `json:"pipolicyhits,omitempty"`
	/**
	* The invocation type.
	 */
	Labeltype string `json:"labeltype,omitempty"`
	/**
	* Name of the label invoked.
	 */
	Labelname string `json:"labelname,omitempty"`
	/**
	* Name of the content switching policy to display. If this parameter is omitted, details of all the policies are displayed.
	 */
	Policyname string `json:"policyname,omitempty"`
}
