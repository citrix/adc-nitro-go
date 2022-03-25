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

package system

/**
* Configuration for global counter data resource.
 */
type Systemglobaldata struct {
	/**
	* Specify the counters to be collected.
	 */
	Counters string `json:"counters,omitempty"`
	/**
	* Specify the (counter) group name which contains all the counters specific to this particular group.
	 */
	Countergroup string `json:"countergroup,omitempty"`
	/**
	* Specify start time in mmddyyyyhhmm to start collecting values from that timestamp.
	 */
	Starttime string `json:"starttime,omitempty"`
	/**
	* Specify end time in mmddyyyyhhmm upto which values have to be collected.
	 */
	Endtime string `json:"endtime,omitempty"`
	/**
	* Last is literal way of saying a certain time period from the current moment. Example: -last 1 hour, -last 1 day, et cetera.
	 */
	Last int `json:"last,omitempty"`
	/**
	* Specify the  time period from current moment. Example 1 x where x = hours/ days/ years.
	 */
	Unit string `json:"unit,omitempty"`
	/**
	* Specifies the source which contains all the stored counter values.
	 */
	Datasource string `json:"datasource,omitempty"`
	/**
	* Specify core ID of the PE in nCore.
	 */
	Core int `json:"core,omitempty"`

	//------- Read only Parameter ---------;

	Response    string `json:"response,omitempty"`
	Startupdate string `json:"startupdate,omitempty"`
	Lastupdate  string `json:"lastupdate,omitempty"`
}
