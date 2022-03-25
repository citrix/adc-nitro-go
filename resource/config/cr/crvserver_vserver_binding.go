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

package cr

/**
* Binding class showing the vserver that can be bound to crvserver.
 */
type Crvservervserverbinding struct {
	/**
	* The Default target server name.
	 */
	Lbvserver string `json:"lbvserver,omitempty"`
	/**
	* Number of hits.
	 */
	Hits uint32 `json:"hits,omitempty"`
	/**
	* Name of the cache redirection virtual server to which to bind the cache redirection policy.
	 */
	Name string `json:"name,omitempty"`
}
