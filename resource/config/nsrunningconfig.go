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
* Configuration for running configuration resource.
 */
type Nsrunningconfig struct {
	/**
	* Include default values of parameters that have not been explicitly configured. If this argument is disabled, such parameters are not included.
	 */
	Withdefaults bool `json:"withdefaults,omitempty"`

	//------- Read only Parameter ---------;

	Response string `json:"response,omitempty"`
}