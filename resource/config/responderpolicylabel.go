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
* Configuration for responder policy label resource.
 */
type Responderpolicylabel struct {
	/**
	* Name for the responder policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters. Cannot be changed after the responder policy label is added.
		The following requirement applies only to the Citrix ADC CLI:
		If the name includes one or more spaces, enclose the name in double or single quotation marks (for example, "my responder policy label" or my responder policy label').
	*/
	Labelname string `json:"labelname,omitempty"`
	/**
	* Type of responses sent by the policies bound to this policy label. Types are:
		* HTTP - HTTP responses.
		* OTHERTCP - NON-HTTP TCP responses.
		* SIP_UDP - SIP responses.
		* RADIUS - RADIUS responses.
		* MYSQL - SQL responses in MySQL format.
		* MSSQL - SQL responses in Microsoft SQL format.
		* NAT - NAT response.
		* MQTT - Trigger policies bind with MQTT type.
		* MQTT_JUMBO - Trigger policies bind with MQTT Jumbo type.
	*/
	Policylabeltype string `json:"policylabeltype,omitempty"`
	/**
	* Any comments to preserve information about this responder policy label.
	 */
	Comment string `json:"comment,omitempty"`
	/**
	* New name for the responder policy label. Must begin with a letter, number, or the underscore character (_), and must contain only letters, numbers, and the hyphen (-), period (.) hash (#), space ( ), at (@), equals (=), colon (:), and underscore characters.
	 */
	Newname string `json:"newname,omitempty"`

	//------- Read only Parameter ---------;

	Numpol                 string `json:"numpol,omitempty"`
	Hits                   string `json:"hits,omitempty"`
	Priority               string `json:"priority,omitempty"`
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Invokelabelname        string `json:"invoke_labelname,omitempty"`
}