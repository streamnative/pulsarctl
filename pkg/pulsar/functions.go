// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package pulsar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"

	"github.com/streamnative/pulsarctl/pkg/cli"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

// Functions is admin interface for functions management
type Functions interface {
	// CreateFunc create a new function.
	CreateFunc(data *utils.FunctionConfig, fileName string) error

	// CreateFuncWithURL create a new function by providing url from which fun-pkg can be downloaded.
	// supported url: http/file
	// eg:
	//  File: file:/dir/fileName.jar
	//  Http: http://www.repo.com/fileName.jar
	//
	// @param functionConfig
	//      the function configuration object
	// @param pkgURL
	//      url from which pkg can be downloaded
	CreateFuncWithURL(data *utils.FunctionConfig, pkgURL string) error

	// StopFunction stop all function instances
	StopFunction(tenant, namespace, name string) error

	// StopFunctionWithID stop function instance
	StopFunctionWithID(tenant, namespace, name string, instanceID int) error

	// DeleteFunction delete an existing function
	DeleteFunction(tenant, namespace, name string) error

	// StartFunction start all function instances
	StartFunction(tenant, namespace, name string) error

	// StartFunctionWithID start function instance
	StartFunctionWithID(tenant, namespace, name string, instanceID int) error

	// RestartFunction restart all function instances
	RestartFunction(tenant, namespace, name string) error

	// RestartFunctionWithID restart function instance
	RestartFunctionWithID(tenant, namespace, name string, instanceID int) error

	// GetFunctions returns the list of functions
	GetFunctions(tenant, namespace string) ([]string, error)

	// GetFunction returns the configuration for the specified function
	GetFunction(tenant, namespace, name string) (utils.FunctionConfig, error)

	// GetFunctionStatus returns the current status of a function
	GetFunctionStatus(tenant, namespace, name string) (utils.FunctionStatus, error)

	// GetFunctionStatusWithInstanceID returns the current status of a function instance
	GetFunctionStatusWithInstanceID(tenant, namespace, name string, instanceID int) (
		utils.FunctionInstanceStatusData, error)

	// GetFunctionStats returns the current stats of a function
	GetFunctionStats(tenant, namespace, name string) (utils.FunctionStats, error)

	// GetFunctionStatsWithInstanceID gets the current stats of a function instance
	GetFunctionStatsWithInstanceID(tenant, namespace, name string, instanceID int) (utils.FunctionInstanceStatsData, error)

	// GetFunctionState fetch the current state associated with a Pulsar Function
	//
	// Response Example:
	// 		{ "value : 12, version : 2"}
	GetFunctionState(tenant, namespace, name, key string) (utils.FunctionState, error)

	// PutFunctionState puts the given state associated with a Pulsar Function
	PutFunctionState(tenant, namespace, name string, state utils.FunctionState) error

	// TriggerFunction triggers the function by writing to the input topic
	TriggerFunction(tenant, namespace, name, topic, triggerValue, triggerFile string) (string, error)

	// UpdateFunction updates the configuration for a function.
	UpdateFunction(functionConfig *utils.FunctionConfig, fileName string, updateOptions *utils.UpdateOptions) error

	// UpdateFunctionWithURL updates the configuration for a function.
	//
	// Update a function by providing url from which fun-pkg can be downloaded. supported url: http/file
	// eg:
	// File: file:/dir/fileName.jar
	// Http: http://www.repo.com/fileName.jar
	UpdateFunctionWithURL(functionConfig *utils.FunctionConfig, pkgURL string, updateOptions *utils.UpdateOptions) error
}

type functions struct {
	client   *pulsarClient
	request  *cli.Client
	basePath string
}

// Functions is used to access the functions endpoints
func (c *pulsarClient) Functions() Functions {
	return &functions{
		client:   c,
		request:  c.Client,
		basePath: "/functions",
	}
}

func (f *functions) createStringFromField(w *multipart.Writer, value string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s" `, value))
	h.Set("Content-Type", "application/json")
	return w.CreatePart(h)
}

func (f *functions) createTextFromFiled(w *multipart.Writer, value string) (io.Writer, error) {
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s" `, value))
	h.Set("Content-Type", "text/plain")
	return w.CreatePart(h)
}

func (f *functions) CreateFunc(funcConf *utils.FunctionConfig, fileName string) error {
	endpoint := f.client.endpoint(f.basePath, funcConf.Tenant, funcConf.Namespace, funcConf.Name)

	// buffer to store our request as bytes
	bodyBuf := bytes.NewBufferString("")

	multiPartWriter := multipart.NewWriter(bodyBuf)

	jsonData, err := json.Marshal(funcConf)
	if err != nil {
		return err
	}

	stringWriter, err := f.createStringFromField(multiPartWriter, "functionConfig")
	if err != nil {
		return err
	}

	_, err = stringWriter.Write(jsonData)
	if err != nil {
		return err
	}

	if fileName != "" && !strings.HasPrefix(fileName, "builtin://") {
		// If the function code is built in, we don't need to submit here
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		part, err := multiPartWriter.CreateFormFile("data", filepath.Base(file.Name()))

		if err != nil {
			return err
		}

		// copy the actual file content to the filed's writer
		_, err = io.Copy(part, file)
		if err != nil {
			return err
		}
	}

	// In here, we completed adding the file and the fields, let's close the multipart writer
	// So it writes the ending boundary
	if err = multiPartWriter.Close(); err != nil {
		return err
	}

	contentType := multiPartWriter.FormDataContentType()
	err = f.request.PostWithMultiPart(endpoint, nil, bodyBuf, contentType)
	if err != nil {
		return err
	}

	return nil
}

func (f *functions) CreateFuncWithURL(funcConf *utils.FunctionConfig, pkgURL string) error {
	endpoint := f.client.endpoint(f.basePath, funcConf.Tenant, funcConf.Namespace, funcConf.Name)
	// buffer to store our request as bytes
	bodyBuf := bytes.NewBufferString("")

	multiPartWriter := multipart.NewWriter(bodyBuf)

	textWriter, err := f.createTextFromFiled(multiPartWriter, "url")
	if err != nil {
		return err
	}

	_, err = textWriter.Write([]byte(pkgURL))
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(funcConf)
	if err != nil {
		return err
	}

	stringWriter, err := f.createStringFromField(multiPartWriter, "functionConfig")
	if err != nil {
		return err
	}

	_, err = stringWriter.Write(jsonData)
	if err != nil {
		return err
	}

	if err = multiPartWriter.Close(); err != nil {
		return err
	}

	contentType := multiPartWriter.FormDataContentType()
	err = f.request.PostWithMultiPart(endpoint, nil, bodyBuf, contentType)
	if err != nil {
		return err
	}

	return nil
}

func (f *functions) StopFunction(tenant, namespace, name string) error {
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name)
	return f.request.Post(endpoint+"/stop", "")
}

func (f *functions) StopFunctionWithID(tenant, namespace, name string, instanceID int) error {
	id := fmt.Sprintf("%d", instanceID)
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name, id)

	return f.request.Post(endpoint+"/stop", "")
}

func (f *functions) DeleteFunction(tenant, namespace, name string) error {
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name)
	return f.request.Delete(endpoint)
}

func (f *functions) StartFunction(tenant, namespace, name string) error {
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name)
	return f.request.Post(endpoint+"/start", "")
}

func (f *functions) StartFunctionWithID(tenant, namespace, name string, instanceID int) error {
	id := fmt.Sprintf("%d", instanceID)
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name, id)

	return f.request.Post(endpoint+"/start", "")
}

func (f *functions) RestartFunction(tenant, namespace, name string) error {
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name)
	return f.request.Post(endpoint+"/restart", "")
}

func (f *functions) RestartFunctionWithID(tenant, namespace, name string, instanceID int) error {
	id := fmt.Sprintf("%d", instanceID)
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name, id)

	return f.request.Post(endpoint+"/restart", "")
}

func (f *functions) GetFunctions(tenant, namespace string) ([]string, error) {
	var functions []string
	endpoint := f.client.endpoint(f.basePath, tenant, namespace)
	err := f.request.Get(endpoint, &functions)
	return functions, err
}

func (f *functions) GetFunction(tenant, namespace, name string) (utils.FunctionConfig, error) {
	var functionConfig utils.FunctionConfig
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name)
	err := f.request.Get(endpoint, &functionConfig)
	return functionConfig, err
}

func (f *functions) UpdateFunction(functionConfig *utils.FunctionConfig, fileName string,
	updateOptions *utils.UpdateOptions) error {
	endpoint := f.client.endpoint(f.basePath, functionConfig.Tenant, functionConfig.Namespace, functionConfig.Name)
	// buffer to store our request as bytes
	bodyBuf := bytes.NewBufferString("")

	multiPartWriter := multipart.NewWriter(bodyBuf)

	jsonData, err := json.Marshal(functionConfig)
	if err != nil {
		return err
	}

	stringWriter, err := f.createStringFromField(multiPartWriter, "functionConfig")
	if err != nil {
		return err
	}

	_, err = stringWriter.Write(jsonData)
	if err != nil {
		return err
	}

	if updateOptions != nil {
		updateData, err := json.Marshal(updateOptions)
		if err != nil {
			return err
		}

		updateStrWriter, err := f.createStringFromField(multiPartWriter, "updateOptions")
		if err != nil {
			return err
		}

		_, err = updateStrWriter.Write(updateData)
		if err != nil {
			return err
		}
	}

	if fileName != "" && !strings.HasPrefix(fileName, "builtin://") {
		// If the function code is built in, we don't need to submit here
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		part, err := multiPartWriter.CreateFormFile("data", filepath.Base(file.Name()))

		if err != nil {
			return err
		}

		// copy the actual file content to the filed's writer
		_, err = io.Copy(part, file)
		if err != nil {
			return err
		}
	}

	// In here, we completed adding the file and the fields, let's close the multipart writer
	// So it writes the ending boundary
	if err = multiPartWriter.Close(); err != nil {
		return err
	}

	contentType := multiPartWriter.FormDataContentType()
	err = f.request.PutWithMultiPart(endpoint, bodyBuf, contentType)
	if err != nil {
		return err
	}

	return nil
}

func (f *functions) UpdateFunctionWithURL(functionConfig *utils.FunctionConfig, pkgURL string,
	updateOptions *utils.UpdateOptions) error {
	endpoint := f.client.endpoint(f.basePath, functionConfig.Tenant, functionConfig.Namespace, functionConfig.Name)
	// buffer to store our request as bytes
	bodyBuf := bytes.NewBufferString("")

	multiPartWriter := multipart.NewWriter(bodyBuf)

	textWriter, err := f.createTextFromFiled(multiPartWriter, "url")
	if err != nil {
		return err
	}

	_, err = textWriter.Write([]byte(pkgURL))
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(functionConfig)
	if err != nil {
		return err
	}

	stringWriter, err := f.createStringFromField(multiPartWriter, "functionConfig")
	if err != nil {
		return err
	}

	_, err = stringWriter.Write(jsonData)
	if err != nil {
		return err
	}

	if updateOptions != nil {
		updateData, err := json.Marshal(updateOptions)
		if err != nil {
			return err
		}

		updateStrWriter, err := f.createStringFromField(multiPartWriter, "updateOptions")
		if err != nil {
			return err
		}

		_, err = updateStrWriter.Write(updateData)
		if err != nil {
			return err
		}
	}

	// In here, we completed adding the file and the fields, let's close the multipart writer
	// So it writes the ending boundary
	if err = multiPartWriter.Close(); err != nil {
		return err
	}

	contentType := multiPartWriter.FormDataContentType()
	err = f.request.PutWithMultiPart(endpoint, bodyBuf, contentType)
	if err != nil {
		return err
	}

	return nil
}

func (f *functions) GetFunctionStatus(tenant, namespace, name string) (utils.FunctionStatus, error) {
	var functionStatus utils.FunctionStatus
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name)
	err := f.request.Get(endpoint+"/status", &functionStatus)
	return functionStatus, err
}

func (f *functions) GetFunctionStatusWithInstanceID(tenant, namespace, name string,
	instanceID int) (utils.FunctionInstanceStatusData, error) {
	var functionInstanceStatusData utils.FunctionInstanceStatusData
	id := fmt.Sprintf("%d", instanceID)
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name, id)
	err := f.request.Get(endpoint+"/status", &functionInstanceStatusData)
	return functionInstanceStatusData, err
}

func (f *functions) GetFunctionStats(tenant, namespace, name string) (utils.FunctionStats, error) {
	var functionStats utils.FunctionStats
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name)
	err := f.request.Get(endpoint+"/stats", &functionStats)
	return functionStats, err
}

func (f *functions) GetFunctionStatsWithInstanceID(tenant, namespace, name string,
	instanceID int) (utils.FunctionInstanceStatsData, error) {
	var functionInstanceStatsData utils.FunctionInstanceStatsData
	id := fmt.Sprintf("%d", instanceID)
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name, id)
	err := f.request.Get(endpoint+"/stats", &functionInstanceStatsData)
	return functionInstanceStatsData, err
}

func (f *functions) GetFunctionState(tenant, namespace, name, key string) (utils.FunctionState, error) {
	var functionState utils.FunctionState
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name, "state", key)
	err := f.request.Get(endpoint, &functionState)
	return functionState, err
}

func (f *functions) PutFunctionState(tenant, namespace, name string, state utils.FunctionState) error {
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name, "state", state.Key)

	// buffer to store our request as bytes
	bodyBuf := bytes.NewBufferString("")

	multiPartWriter := multipart.NewWriter(bodyBuf)

	stateData, err := json.Marshal(state)

	if err != nil {
		return err
	}

	stateWriter, err := f.createStringFromField(multiPartWriter, "state")
	if err != nil {
		return err
	}

	_, err = stateWriter.Write(stateData)

	if err != nil {
		return err
	}

	// In here, we completed adding the file and the fields, let's close the multipart writer
	// So it writes the ending boundary
	if err = multiPartWriter.Close(); err != nil {
		return err
	}

	contentType := multiPartWriter.FormDataContentType()

	err = f.request.PostWithMultiPart(endpoint, nil, bodyBuf, contentType)

	if err != nil {
		return err
	}

	return nil
}

func (f *functions) TriggerFunction(tenant, namespace, name, topic, triggerValue, triggerFile string) (string, error) {
	endpoint := f.client.endpoint(f.basePath, tenant, namespace, name, "trigger")

	// buffer to store our request as bytes
	bodyBuf := bytes.NewBufferString("")

	multiPartWriter := multipart.NewWriter(bodyBuf)

	if triggerFile != "" {
		file, err := os.Open(triggerFile)
		if err != nil {
			return "", err
		}
		defer file.Close()

		part, err := multiPartWriter.CreateFormFile("dataStream", filepath.Base(file.Name()))

		if err != nil {
			return "", err
		}

		// copy the actual file content to the filed's writer
		_, err = io.Copy(part, file)
		if err != nil {
			return "", err
		}
	}

	if triggerValue != "" {
		valueWriter, err := f.createTextFromFiled(multiPartWriter, "data")
		if err != nil {
			return "", err
		}

		_, err = valueWriter.Write([]byte(triggerValue))
		if err != nil {
			return "", err
		}
	}

	if topic != "" {
		topicWriter, err := f.createTextFromFiled(multiPartWriter, "topic")
		if err != nil {
			return "", err
		}

		_, err = topicWriter.Write([]byte(topic))
		if err != nil {
			return "", err
		}
	}

	// In here, we completed adding the file and the fields, let's close the multipart writer
	// So it writes the ending boundary
	if err := multiPartWriter.Close(); err != nil {
		return "", err
	}

	contentType := multiPartWriter.FormDataContentType()
	var str string
	err := f.request.PostWithMultiPart(endpoint, &str, bodyBuf, contentType)
	if err != nil {
		return "", err
	}

	return str, nil
}
