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

package functions

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

const (
	HTTP    = "http"
	FILE    = "file"
	BUILTIN = "builtin"

	PublicTenant     = "public"
	DefaultNamespace = "default"
)

func isFunctionPackageUrlSupported(functionPkgUrl string) bool {
	return functionPkgUrl != "" && strings.HasPrefix(functionPkgUrl, HTTP) ||
		strings.HasPrefix(functionPkgUrl, FILE)
}

func inferMissingFunctionName(funcConf *pulsar.FunctionConfig) {
	className := funcConf.ClassName
	domains := strings.Split(className, "\\.")

	if len(domains) == 0 {
		funcConf.Name = funcConf.ClassName
	} else {
		funcConf.Name = domains[len(domains)-1]
	}
}

func inferMissingTenant(funcConf *pulsar.FunctionConfig) {
	funcConf.Tenant = PublicTenant
}

func inferMissingNamespace(funcConf *pulsar.FunctionConfig) {
	funcConf.Namespace = DefaultNamespace
}

func inferMissingSourceArguments(sourceConf *pulsar.SourceConfig) {
	if sourceConf.Tenant == "" {
		sourceConf.Tenant = PublicTenant
	}

	if sourceConf.Namespace == "" {
		sourceConf.Namespace = DefaultNamespace
	}

	if sourceConf.Parallelism == 0 {
		sourceConf.Parallelism = 1
	}
}

func inferMissingSinkeArguments(sinkConf *pulsar.SinkConfig) {
	if sinkConf.Tenant == "" {
		sinkConf.Tenant = PublicTenant
	}

	if sinkConf.Namespace == "" {
		sinkConf.Namespace = DefaultNamespace
	}

	if sinkConf.Parallelism == 0 {
		sinkConf.Parallelism = 1
	}
}

func parseFullyQualifiedFunctionName(fqfn string, functionConfig *pulsar.FunctionConfig) error {
	args := strings.Split(fqfn, "/")
	if len(args) != 3 {
		return errors.New("fully qualified function names (FQFNs) must be of the form tenant/namespace/name")
	}

	functionConfig.Tenant = args[0]
	functionConfig.Namespace = args[1]
	functionConfig.Name = args[2]

	return nil
}

func isFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println(info)
		return false
	}
	fmt.Println("exists", info.Name(), info.Size(), info.ModTime())
	return true
}

func processArgs(funcData *pulsar.FunctionData) error {
	// Initialize config builder either from a supplied YAML config file or from scratch
	if funcData.FuncConf != nil {
		// no-op
	} else {
		funcData.FuncConf = new(pulsar.FunctionConfig)
	}

	if funcData.FunctionConfigFile != "" {
		yamlFile, err := ioutil.ReadFile(funcData.FunctionConfigFile)
		if err == nil {
			err = yaml.Unmarshal(yamlFile, funcData.FuncConf)
			if err != nil {
				return fmt.Errorf("unmarshal yaml file error:%s", err.Error())
			}
		} else if err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("load conf file failed, err:%s", err.Error())
		}
	}

	if funcData.FQFN != "" {
		err := parseFullyQualifiedFunctionName(funcData.FQFN, funcData.FuncConf)
		if err != nil {
			return err
		}
	} else {
		if funcData.Tenant != "" {
			funcData.FuncConf.Tenant = funcData.Tenant
		}
		if funcData.Namespace != "" {
			funcData.FuncConf.Namespace = funcData.Namespace
		}
		if funcData.FuncName != "" {
			funcData.FuncConf.Name = funcData.FuncName
		}
	}

	if funcData.Inputs != "" {
		inputTopics := strings.Split(funcData.Inputs, ",")
		funcData.FuncConf.Inputs = inputTopics
	}

	if funcData.CustomSerDeInputs != "" {
		customSerdeInputMap := make(map[string]string)
		err := json.Unmarshal([]byte(funcData.CustomSerDeInputs), &customSerdeInputMap)
		if err != nil {
			return err
		}
		funcData.FuncConf.CustomSerdeInputs = customSerdeInputMap
	}

	if funcData.CustomSchemaInput != "" {
		customSchemaInputMap := make(map[string]string)
		err := json.Unmarshal([]byte(funcData.CustomSchemaInput), &customSchemaInputMap)
		if err != nil {
			return err
		}
		funcData.FuncConf.CustomSchemaInputs = customSchemaInputMap
	}

	if funcData.TopicsPattern != "" {
		funcData.FuncConf.TopicsPattern = &funcData.TopicsPattern
	}

	if funcData.Output != "" {
		funcData.FuncConf.Output = funcData.Output
	}

	if funcData.LogTopic != "" {
		funcData.FuncConf.LogTopic = funcData.LogTopic
	}

	if funcData.ClassName != "" {
		funcData.FuncConf.ClassName = funcData.ClassName
	}

	if funcData.OutputSerDeClassName != "" {
		funcData.FuncConf.OutputSerdeClassName = funcData.OutputSerDeClassName
	}

	if funcData.SchemaType != "" {
		funcData.FuncConf.OutputSchemaType = funcData.SchemaType
	}

	// processingGuarantees default value is 0, means AtLeastOnce.
	if funcData.ProcessingGuarantees != "" {
		switch funcData.ProcessingGuarantees {
		case "ATMOST_ONCE":
			funcData.FuncConf.ProcessingGuarantees = pulsar.AtMostOnce
		case "EFFECTIVELY_ONCE":
			funcData.FuncConf.ProcessingGuarantees = pulsar.EffectivelyOnce
		case "ATLEAST_ONCE":
			funcData.FuncConf.ProcessingGuarantees = pulsar.AtLeasetOnce
		default:
			funcData.FuncConf.ProcessingGuarantees = pulsar.AtLeasetOnce
		}
	}

	if funcData.RetainOrdering {
		funcData.FuncConf.RetainOrdering = funcData.RetainOrdering
	}

	if funcData.SubsName != "" {
		funcData.FuncConf.SubName = funcData.SubsName
	}

	if funcData.UserConfig != "" {
		userConfigMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(funcData.UserConfig), &userConfigMap)
		if err != nil {
			return err
		}

		funcData.FuncConf.UserConfig = userConfigMap
	}

	if funcData.FuncConf.UserConfig == nil {
		funcData.FuncConf.UserConfig = make(map[string]interface{})
	}

	if funcData.Parallelism != 0 {
		funcData.FuncConf.Parallelism = funcData.Parallelism
	} else {
		funcData.FuncConf.Parallelism = 1
	}

	if funcData.CPU != 0 {
		if funcData.FuncConf.Resources == nil {
			funcData.FuncConf.Resources = pulsar.NewDefaultResources()
		}

		funcData.FuncConf.Resources.CPU = funcData.CPU
	}

	if funcData.Disk != 0 {
		if funcData.FuncConf.Resources == nil {
			funcData.FuncConf.Resources = pulsar.NewDefaultResources()
		}

		funcData.FuncConf.Resources.Disk = funcData.Disk
	}

	if funcData.RAM != 0 {
		if funcData.FuncConf.Resources == nil {
			funcData.FuncConf.Resources = pulsar.NewDefaultResources()
		}

		funcData.FuncConf.Resources.Ram = funcData.RAM
	}

	if funcData.TimeoutMs != 0 {
		funcData.FuncConf.TimeoutMs = &funcData.TimeoutMs
	}

	// window configs
	if funcData.WindowLengthCount != 0 {
		if funcData.FuncConf.WindowConfig == nil {
			funcData.FuncConf.WindowConfig = pulsar.NewDefaultWindowConfing()
		}

		funcData.FuncConf.WindowConfig.WindowLengthCount = funcData.WindowLengthCount
	}

	if funcData.WindowLengthDurationMs != 0 {
		if funcData.FuncConf.WindowConfig == nil {
			funcData.FuncConf.WindowConfig = pulsar.NewDefaultWindowConfing()
		}

		funcData.FuncConf.WindowConfig.WindowLengthDurationMs = funcData.WindowLengthDurationMs
	}

	if funcData.SlidingIntervalCount != 0 {
		if funcData.FuncConf.WindowConfig == nil {
			funcData.FuncConf.WindowConfig = pulsar.NewDefaultWindowConfing()
		}

		funcData.FuncConf.WindowConfig.SlidingIntervalCount = funcData.SlidingIntervalCount
	}

	if funcData.SlidingIntervalDurationMs != 0 {
		if funcData.FuncConf.WindowConfig == nil {
			funcData.FuncConf.WindowConfig = pulsar.NewDefaultWindowConfing()
		}

		funcData.FuncConf.WindowConfig.SlidingIntervalDurationMs = funcData.SlidingIntervalDurationMs
	}

	if funcData.AutoAck {
		funcData.FuncConf.AutoAck = funcData.AutoAck
	}

	if funcData.MaxMessageRetries != 0 {
		funcData.FuncConf.MaxMessageRetries = funcData.MaxMessageRetries
	}

	if funcData.DeadLetterTopic != "" {
		funcData.FuncConf.DeadLetterTopic = funcData.DeadLetterTopic
	}

	if funcData.Jar != "" {
		funcData.FuncConf.Jar = funcData.Jar
	}

	if funcData.Py != "" {
		funcData.FuncConf.Py = funcData.Py
	}

	if funcData.Go != "" {
		funcData.FuncConf.Go = funcData.Go
	}

	if funcData.FuncConf.Jar != "" {
		funcData.UserCodeFile = funcData.FuncConf.Jar
	} else if funcData.FuncConf.Py != "" {
		funcData.UserCodeFile = funcData.FuncConf.Py
	} else if funcData.FuncConf.Go != "" {
		funcData.UserCodeFile = funcData.FuncConf.Go
	}

	// check if configs are valid
	return validateFunctionConfigs(funcData.FuncConf)
}

func validateFunctionConfigs(functionConfig *pulsar.FunctionConfig) error {
	// go doesn't need className
	if functionConfig.Runtime == pulsar.Python || functionConfig.Runtime == pulsar.Java {
		if functionConfig.ClassName == "" {
			return errors.New("no Function Classname specified")
		}
	}

	if functionConfig.Name == "" {
		inferMissingFunctionName(functionConfig)
	}

	if functionConfig.Tenant == "" {
		inferMissingTenant(functionConfig)
	}

	if functionConfig.Namespace == "" {
		inferMissingNamespace(functionConfig)
	}

	if functionConfig.Jar != "" && functionConfig.Py != "" && functionConfig.Go != "" {
		return errors.New("either a Java jar or a Python file or a Go executable binary needs to " +
			"be specified for the function, Cannot specify both")
	}

	if functionConfig.Jar == "" && functionConfig.Py == "" && functionConfig.Go == "" {
		return errors.New("either a Java jar or a Python file or a Go executable binary needs to " +
			"be specified for the function. Please specify one")
	}

	if functionConfig.Jar != "" && !isFunctionPackageUrlSupported(functionConfig.Jar) &&
		!isFileExist(functionConfig.Jar) {
		return errors.New("the specified jar file does not exist")
	}

	if functionConfig.Py != "" && !isFunctionPackageUrlSupported(functionConfig.Py) &&
		!isFileExist(functionConfig.Py) {
		return errors.New("the specified py file does not exist")
	}

	if functionConfig.Go != "" && !isFunctionPackageUrlSupported(functionConfig.Go) &&
		!isFileExist(functionConfig.Go) {
		return errors.New("the specified go file does not exist")
	}

	return nil
}
