package smartag

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeFlowLogs invokes the smartag.DescribeFlowLogs API synchronously
func (client *Client) DescribeFlowLogs(request *DescribeFlowLogsRequest) (response *DescribeFlowLogsResponse, err error) {
	response = CreateDescribeFlowLogsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFlowLogsWithChan invokes the smartag.DescribeFlowLogs API asynchronously
func (client *Client) DescribeFlowLogsWithChan(request *DescribeFlowLogsRequest) (<-chan *DescribeFlowLogsResponse, <-chan error) {
	responseChan := make(chan *DescribeFlowLogsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFlowLogs(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeFlowLogsWithCallback invokes the smartag.DescribeFlowLogs API asynchronously
func (client *Client) DescribeFlowLogsWithCallback(request *DescribeFlowLogsRequest, callback func(response *DescribeFlowLogsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFlowLogsResponse
		var err error
		defer close(result)
		response, err = client.DescribeFlowLogs(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeFlowLogsRequest is the request struct for api DescribeFlowLogs
type DescribeFlowLogsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Description          string           `position:"Query" name:"Description"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	OutputType           string           `position:"Query" name:"OutputType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	FlowLogId            string           `position:"Query" name:"FlowLogId"`
	FlowLogName          string           `position:"Query" name:"FlowLogName"`
	Status               string           `position:"Query" name:"Status"`
}

// DescribeFlowLogsResponse is the response struct for api DescribeFlowLogs
type DescribeFlowLogsResponse struct {
	*responses.BaseResponse
	TotalCount int      `json:"TotalCount" xml:"TotalCount"`
	PageSize   int      `json:"PageSize" xml:"PageSize"`
	RequestId  string   `json:"RequestId" xml:"RequestId"`
	PageNumber int      `json:"PageNumber" xml:"PageNumber"`
	FlowLogs   FlowLogs `json:"FlowLogs" xml:"FlowLogs"`
}

// CreateDescribeFlowLogsRequest creates a request to invoke DescribeFlowLogs API
func CreateDescribeFlowLogsRequest() (request *DescribeFlowLogsRequest) {
	request = &DescribeFlowLogsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Smartag", "2018-03-13", "DescribeFlowLogs", "smartag", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeFlowLogsResponse creates a response to parse from DescribeFlowLogs response
func CreateDescribeFlowLogsResponse() (response *DescribeFlowLogsResponse) {
	response = &DescribeFlowLogsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
