package bssopenapi

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ItemInDescribeSavingsPlansUsageDetail is a nested struct in bssopenapi response
type ItemInDescribeSavingsPlansUsageDetail struct {
	UserId          int64   `json:"UserId" xml:"UserId"`
	UserName        string  `json:"UserName" xml:"UserName"`
	InstanceId      string  `json:"InstanceId" xml:"InstanceId"`
	StartPeriod     string  `json:"StartPeriod" xml:"StartPeriod"`
	EndPeriod       string  `json:"EndPeriod" xml:"EndPeriod"`
	Status          string  `json:"Status" xml:"Status"`
	Type            string  `json:"Type" xml:"Type"`
	UsagePercentage float64 `json:"UsagePercentage" xml:"UsagePercentage"`
	PoolValue       float64 `json:"PoolValue" xml:"PoolValue"`
	DeductValue     float64 `json:"DeductValue" xml:"DeductValue"`
	PostpaidCost    float64 `json:"PostpaidCost" xml:"PostpaidCost"`
	SavedCost       float64 `json:"SavedCost" xml:"SavedCost"`
	Currency        string  `json:"Currency" xml:"Currency"`
}
