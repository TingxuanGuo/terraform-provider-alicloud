package sls

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

// UpdateApp invokes the sls.UpdateApp API synchronously
// api document: https://help.aliyun.com/api/sls/updateapp.html
func (client *Client) UpdateApp(request *UpdateAppRequest) (response *UpdateAppResponse, err error) {
	response = CreateUpdateAppResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAppWithChan invokes the sls.UpdateApp API asynchronously
// api document: https://help.aliyun.com/api/sls/updateapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAppWithChan(request *UpdateAppRequest) (<-chan *UpdateAppResponse, <-chan error) {
	responseChan := make(chan *UpdateAppResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateApp(request)
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

// UpdateAppWithCallback invokes the sls.UpdateApp API asynchronously
// api document: https://help.aliyun.com/api/sls/updateapp.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateAppWithCallback(request *UpdateAppRequest, callback func(response *UpdateAppResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAppResponse
		var err error
		defer close(result)
		response, err = client.UpdateApp(request)
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

// UpdateAppRequest is the request struct for api UpdateApp
type UpdateAppRequest struct {
	*requests.RpcRequest
	AppName     string `position:"Query" name:"AppName"`
	DisplayName string `position:"Query" name:"DisplayName"`
	Config      string `position:"Body" name:"Config"`
}

// UpdateAppResponse is the response struct for api UpdateApp
type UpdateAppResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Success   string `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdateAppRequest creates a request to invoke UpdateApp API
func CreateUpdateAppRequest() (request *UpdateAppRequest) {
	request = &UpdateAppRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2019-10-23", "UpdateApp", "sls", "openAPI")
	return
}

// CreateUpdateAppResponse creates a response to parse from UpdateApp response
func CreateUpdateAppResponse() (response *UpdateAppResponse) {
	response = &UpdateAppResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
