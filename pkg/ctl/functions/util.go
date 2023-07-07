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
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/pkg/errors"
	util "github.com/streamnative/pulsar-admin-go/pkg/utils"
	"gopkg.in/yaml.v2"

	"github.com/streamnative/pulsarctl/pkg/ctl/utils"
)

func parseFullyQualifiedFunctionName(fqfn string, functionConfig *util.FunctionConfig) error {
	args := strings.Split(fqfn, "/")
	if len(args) != 3 {
		return errors.New("fully qualified function names (FQFNs) must be of the form tenant/namespace/name")
	}

	functionConfig.Tenant = args[0]
	functionConfig.Namespace = args[1]
	functionConfig.Name = args[2]

	return nil
}

func processArgs(funcData *util.FunctionData) error {
	// Initialize config builder either from a supplied YAML config file or from scratch
	if funcData.FuncConf != nil {
		// no-op
	} else {
		funcData.FuncConf = new(util.FunctionConfig)
	}

	if funcData.FunctionConfigFile != "" {
		yamlFile, err := ioutil.ReadFile(funcData.FunctionConfigFile)
		if err == nil {
			err = yaml.Unmarshal(yamlFile, funcData.FuncConf)
			if err != nil {
				return fmt.Errorf("unmarshal yaml file error:%s", err.Error())
			}
		} else if !os.IsNotExist(err) {
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

	if funcData.ProcessingGuarantees != "" {
		funcData.FuncConf.ProcessingGuarantees = funcData.ProcessingGuarantees
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

	if funcData.RetainOrdering {
		funcData.FuncConf.RetainOrdering = funcData.RetainOrdering
	}

	if funcData.RetainKeyOrdering {
		funcData.FuncConf.RetainKeyOrdering = funcData.RetainKeyOrdering
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
			funcData.FuncConf.Resources = util.NewDefaultResources()
		}

		funcData.FuncConf.Resources.CPU = funcData.CPU
	}

	if funcData.Disk != 0 {
		if funcData.FuncConf.Resources == nil {
			funcData.FuncConf.Resources = util.NewDefaultResources()
		}

		funcData.FuncConf.Resources.Disk = funcData.Disk
	}

	if funcData.RAM != 0 {
		if funcData.FuncConf.Resources == nil {
			funcData.FuncConf.Resources = util.NewDefaultResources()
		}

		funcData.FuncConf.Resources.RAM = funcData.RAM
	}

	if funcData.TimeoutMs != 0 {
		funcData.FuncConf.TimeoutMs = &funcData.TimeoutMs
	}

	// window configs
	if funcData.WindowLengthCount != 0 {
		if funcData.FuncConf.WindowConfig == nil {
			funcData.FuncConf.WindowConfig = util.NewDefaultWindowConfing()
		}

		funcData.FuncConf.WindowConfig.WindowLengthCount = &funcData.WindowLengthCount
	}

	if funcData.WindowLengthDurationMs != 0 {
		if funcData.FuncConf.WindowConfig == nil {
			funcData.FuncConf.WindowConfig = util.NewDefaultWindowConfing()
		}

		funcData.FuncConf.WindowConfig.WindowLengthDurationMs = &funcData.WindowLengthDurationMs
	}

	if funcData.SlidingIntervalCount != 0 {
		if funcData.FuncConf.WindowConfig == nil {
			funcData.FuncConf.WindowConfig = util.NewDefaultWindowConfing()
		}

		funcData.FuncConf.WindowConfig.SlidingIntervalCount = &funcData.SlidingIntervalCount
	}

	if funcData.SlidingIntervalDurationMs != 0 {
		if funcData.FuncConf.WindowConfig == nil {
			funcData.FuncConf.WindowConfig = util.NewDefaultWindowConfing()
		}

		funcData.FuncConf.WindowConfig.SlidingIntervalDurationMs = &funcData.SlidingIntervalDurationMs
	}

	if funcData.AutoAck {
		funcData.FuncConf.AutoAck = funcData.AutoAck
	} else {
		funcData.FuncConf.AutoAck = true
	}

	if funcData.MaxMessageRetries != 0 {
		funcData.FuncConf.MaxMessageRetries = &funcData.MaxMessageRetries
	}

	if funcData.DeadLetterTopic != "" {
		funcData.FuncConf.DeadLetterTopic = funcData.DeadLetterTopic
	}

	if funcData.FunctionType != "" {
		jar := fmt.Sprintf("builtin://%s", funcData.FunctionType)
		funcData.FuncConf.Jar = &jar
	} else if *funcData.FuncConf.FunctionType != "" {
		jar := fmt.Sprintf("builtin://%s", *funcData.FuncConf.FunctionType)
		funcData.FuncConf.Jar = &jar
	}

	if funcData.Jar != "" {
		funcData.FuncConf.Jar = &funcData.Jar
	}

	if funcData.Py != "" {
		funcData.FuncConf.Py = &funcData.Py
	}

	if funcData.Go != "" {
		funcData.FuncConf.Go = &funcData.Go
	}

	if funcData.FuncConf.Go != nil {
		funcData.UserCodeFile = *funcData.FuncConf.Go
	}

	if funcData.FuncConf.Py != nil {
		funcData.UserCodeFile = *funcData.FuncConf.Py
	}

	if funcData.FuncConf.Jar != nil {
		funcData.UserCodeFile = *funcData.FuncConf.Jar
	}

	funcData.FuncConf.CleanupSubscription = funcData.CleanupSubscription

	if funcData.ProducerConfig != "" {
		producerConfig := util.ProducerConfig{}
		err := json.Unmarshal([]byte(funcData.ProducerConfig), &producerConfig)
		if err != nil {
			return err
		}

		funcData.FuncConf.ProducerConfig = producerConfig
	}

	if funcData.CustomSchemaOutput != "" {
		schemaOutputs := make(map[string]string)
		err := json.Unmarshal([]byte(funcData.CustomSchemaOutput), &schemaOutputs)
		if err != nil {
			return err
		}

		funcData.FuncConf.CustomSchemaOutputs = schemaOutputs
	}
	if funcData.FuncConf.CustomSchemaOutputs == nil {
		funcData.FuncConf.CustomSchemaOutputs = make(map[string]string)
	}

	if funcData.InputSpecs != "" {
		inputSpecs := make(map[string]util.ConsumerConfig)
		err := json.Unmarshal([]byte(funcData.InputSpecs), &inputSpecs)
		if err != nil {
			return err
		}

		funcData.FuncConf.InputSpecs = inputSpecs
	}

	if funcData.InputTypeClassName != "" {
		funcData.FuncConf.InputTypeClassName = funcData.InputTypeClassName
	}

	if funcData.OutputTypeClassName != "" {
		funcData.FuncConf.OutputTypeClassName = funcData.OutputTypeClassName
	}

	if funcData.BatchBuilder != "" {
		funcData.FuncConf.BatchBuilder = funcData.BatchBuilder
	}

	funcData.FuncConf.ForwardSourceMessageProperty = funcData.ForwardSourceMessageProperty

	if funcData.SubsPosition != "" {
		funcData.FuncConf.SubscriptionPosition = funcData.SubsPosition
	}

	if funcData.SkipToLatest {
		funcData.FuncConf.SkipToLatest = funcData.SkipToLatest
	}

	if funcData.CustomRuntimeOptions != "" {
		funcData.FuncConf.CustomRuntimeOptions = funcData.CustomRuntimeOptions
	}

	if funcData.Secrets != "" {
		secretsMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(funcData.Secrets), &secretsMap)
		if err != nil {
			return err
		}

		funcData.FuncConf.Secrets = secretsMap
	}

	if funcData.FuncConf.Secrets == nil {
		funcData.FuncConf.Secrets = make(map[string]interface{})
	}

	return nil
}

func validateFunctionConfigs(functionConfig *util.FunctionConfig) error {
	if functionConfig.Name == "" {
		utils.InferMissingFunctionName(functionConfig)
	}

	if functionConfig.Tenant == "" {
		utils.InferMissingTenant(functionConfig)
	}

	if functionConfig.Namespace == "" {
		utils.InferMissingNamespace(functionConfig)
	}

	switch utils.NumProvidedStrings(functionConfig.Jar, functionConfig.Py, functionConfig.Go) {
	case 0:
		return errors.New("either a Java jar or a Python file or a Go executable binary needs to " +
			"be specified for the function. Please specify one")
	case 1:
		// proceed
	default:
		return errors.New("either a Java jar or a Python file or a Go executable binary needs to " +
			"be specified for the function, cannot specify more than one")
	}

	if functionConfig.Jar != nil && !strings.HasPrefix(*functionConfig.Jar, "builtin://") &&
		!utils.IsPackageURLSupported(*functionConfig.Jar) &&
		!utils.IsFileExist(*functionConfig.Jar) {
		return errors.New("the specified jar file does not exist")
	}

	if functionConfig.Py != nil && !utils.IsPackageURLSupported(*functionConfig.Py) &&
		!utils.IsFileExist(*functionConfig.Py) {
		return errors.New("the specified py file does not exist")
	}

	if functionConfig.Go != nil && !utils.IsPackageURLSupported(*functionConfig.Go) &&
		!utils.IsFileExist(*functionConfig.Go) {
		return errors.New("the specified go file does not exist")
	}

	if functionConfig.Go != nil {
		functionConfig.Runtime = util.GoRuntime
	}

	if functionConfig.Py != nil {
		functionConfig.Runtime = util.PythonRuntime
	}

	if functionConfig.Jar != nil {
		functionConfig.Runtime = util.JavaRuntime
	}

	// go doesn't need className
	if functionConfig.Runtime == util.JavaRuntime || functionConfig.Runtime == util.PythonRuntime {
		if functionConfig.ClassName == "" {
			return errors.New("no Function Classname specified")
		}
	}

	return nil
}

func processBaseArguments(funcData *util.FunctionData) error {
	usesSetters := funcData.Tenant != "" || funcData.Namespace != "" || funcData.FuncName != ""
	usesFqfn := funcData.FQFN != ""

	// return error if --fqfn is set alongside any combination of --tenant, --namespace, and --name
	if usesFqfn && usesSetters {
		return errors.New("you must specify either a Fully Qualified Function Name (FQFN)" +
			" or tenant, namespace, and function name")
	}

	if usesFqfn {
		// If the --fqfn flag is used, parse tenant, namespace, and name using that flag
		fqfnParts := strings.Split(funcData.FQFN, "/")
		if len(fqfnParts) != 3 {
			return errors.New("fully qualified function names (FQFNs) must be of the form" +
				" tenant/namespace/name")
		}

		funcData.Tenant = fqfnParts[0]
		funcData.Namespace = fqfnParts[1]
		funcData.FuncName = fqfnParts[2]
	} else {
		if funcData.Tenant == "" {
			funcData.Tenant = utils.PublicTenant
		}

		if funcData.Namespace == "" {
			funcData.Namespace = utils.DefaultNamespace
		}

		if funcData.FuncName == "" {
			return errors.New("you must specify a name for the function or a Fully Qualified" +
				" Function Name (FQFN)")
		}
	}

	return nil
}

func processNamespaceCmd(funcData *util.FunctionData) {
	if funcData.Tenant == "" || funcData.Namespace == "" {
		funcData.Tenant = utils.PublicTenant
		funcData.Namespace = utils.DefaultNamespace
	}
}

func checkArgsForUpdate(functionConfig *util.FunctionConfig) error {
	if functionConfig.ClassName == "" {
		if functionConfig.Name == "" {
			return errors.New("function Name not provided")
		}
	} else if functionConfig.Name == "" {
		utils.InferMissingFunctionName(functionConfig)
	}

	if functionConfig.Tenant == "" {
		utils.InferMissingTenant(functionConfig)
	}

	if functionConfig.Namespace == "" {
		utils.InferMissingNamespace(functionConfig)
	}

	return nil
}
