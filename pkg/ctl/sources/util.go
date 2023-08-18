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

package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	util "github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/ctl/utils"
)

func processArguments(sourceData *util.SourceData) error {
	// Initialize config builder either from a supplied YAML config file or from scratch
	if sourceData.SourceConf != nil {
		// no-op
	} else {
		sourceData.SourceConf = new(util.SourceConfig)
	}

	if sourceData.SourceConfigFile != "" {
		yamlFile, err := ioutil.ReadFile(sourceData.SourceConfigFile)
		if err == nil {
			err = yaml.Unmarshal(yamlFile, sourceData.SourceConf)
			if err != nil {
				return fmt.Errorf("unmarshal yaml file error:%s", err.Error())
			}
		} else if !os.IsNotExist(err) {
			return fmt.Errorf("load conf file failed, err:%s", err.Error())
		}
	}

	if sourceData.Tenant != "" {
		sourceData.SourceConf.Tenant = sourceData.Tenant
	}

	if sourceData.Namespace != "" {
		sourceData.SourceConf.Namespace = sourceData.Namespace
	}

	if sourceData.Name != "" {
		sourceData.SourceConf.Name = sourceData.Name
	}

	if sourceData.ClassName != "" {
		sourceData.SourceConf.ClassName = sourceData.ClassName
	}

	if sourceData.DestinationTopicName != "" {
		sourceData.SourceConf.TopicName = sourceData.DestinationTopicName
	}

	if sourceData.DeserializationClassName != "" {
		sourceData.SourceConf.SerdeClassName = sourceData.DeserializationClassName
	}

	if sourceData.SchemaType != "" {
		sourceData.SourceConf.SchemaType = sourceData.SchemaType
	}

	if sourceData.ProcessingGuarantees != "" {
		sourceData.SourceConf.ProcessingGuarantees = sourceData.ProcessingGuarantees
	}

	if sourceData.Parallelism != 0 {
		sourceData.SourceConf.Parallelism = sourceData.Parallelism
	} else {
		sourceData.SourceConf.Parallelism = 1
	}

	if sourceData.Archive != "" && sourceData.SourceType != "" {
		return errors.New("Cannot specify both archive and source-type")
	}

	if sourceData.Archive != "" {
		sourceData.SourceConf.Archive = sourceData.Archive
	}

	if sourceData.SourceType != "" {
		sourceData.SourceConf.Archive = validateSourceType(sourceData.SourceType)
	}

	if sourceData.CPU != 0 {
		if sourceData.SourceConf.Resources == nil {
			sourceData.SourceConf.Resources = util.NewDefaultResources()
		}

		sourceData.SourceConf.Resources.CPU = sourceData.CPU
	}

	if sourceData.Disk != 0 {
		if sourceData.SourceConf.Resources == nil {
			sourceData.SourceConf.Resources = util.NewDefaultResources()
		}

		sourceData.SourceConf.Resources.Disk = sourceData.Disk
	}

	if sourceData.RAM != 0 {
		if sourceData.SourceConf.Resources == nil {
			sourceData.SourceConf.Resources = util.NewDefaultResources()
		}

		sourceData.SourceConf.Resources.RAM = sourceData.RAM
	}

	if sourceData.SourceConfigString != "" {
		sourceData.SourceConf.Configs = parseConfigs(sourceData.SourceConfigString)
	}

	if sourceData.ProducerConfig != "" {
		producerConfig := &util.ProducerConfig{}
		err := json.Unmarshal([]byte(sourceData.ProducerConfig), producerConfig)
		if err != nil {
			return err
		}

		sourceData.SourceConf.ProducerConfig = producerConfig
	}

	if sourceData.BatchBuilder != "" {
		sourceData.SourceConf.BatchBuilder = sourceData.BatchBuilder
	}

	if sourceData.BatchSourceConfigString != "" {
		batchSourceConfig := &util.BatchSourceConfig{}
		err := json.Unmarshal([]byte(sourceData.BatchSourceConfigString), batchSourceConfig)
		if err != nil {
			return err
		}

		sourceData.SourceConf.BatchSourceConfig = batchSourceConfig
	}

	if sourceData.CustomRuntimeOptions != "" {
		sourceData.SourceConf.CustomRuntimeOptions = sourceData.CustomRuntimeOptions
	}

	if sourceData.Secrets != "" {
		secretsMap := make(map[string]interface{})
		err := json.Unmarshal([]byte(sourceData.Secrets), &secretsMap)
		if err != nil {
			return err
		}

		sourceData.SourceConf.Secrets = secretsMap
	}

	if sourceData.SourceConf.Secrets == nil {
		sourceData.SourceConf.Secrets = make(map[string]interface{})
	}

	return nil
}

func validateSourceType(sourceType string) string {
	availableSources := make([]string, 0, 10)
	admin := cmdutils.NewPulsarClientWithAPIVersion(config.V3)
	connectorDefinition, err := admin.Sources().GetBuiltInSources()
	if err != nil {
		log.Printf("get builtin sources error: %s", err.Error())
		return ""
	}

	for _, value := range connectorDefinition {
		availableSources = append(availableSources, value.Name)
	}

	availableSourcesString := strings.Join(availableSources, " ")
	if !strings.Contains(availableSourcesString, sourceType) {
		log.Printf("invalid source type [%s] -- Available sources are: %s", sourceType, availableSources)
		return ""
	}

	// Source type is a valid built-in connector type
	return "builtin://" + sourceType
}

func parseConfigs(str string) map[string]interface{} {
	var resMap map[string]interface{}

	err := json.Unmarshal([]byte(str), &resMap)
	if err != nil {
		return nil
	}

	return resMap
}

func validateSourceConfigs(sourceConfig *util.SourceConfig) error {
	if sourceConfig.Archive == "" {
		return errors.New("Source archive not specified")
	}

	utils.InferMissingSourceArguments(sourceConfig)

	if utils.IsPackageURLSupported(sourceConfig.Archive) && strings.HasPrefix(sourceConfig.Archive, utils.BUILTIN) {
		if !utils.IsFileExist(sourceConfig.Archive) {
			return fmt.Errorf("source Archive %s does not exist", sourceConfig.Archive)
		}
	}

	if sourceConfig.Name == "" {
		return errors.New("source name not specified")
	}
	return nil
}

func checkArgsForUpdate(sourceConfig *util.SourceConfig) {
	if sourceConfig.Tenant == "" {
		sourceConfig.Tenant = utils.PublicTenant
	}

	if sourceConfig.Namespace == "" {
		sourceConfig.Namespace = utils.DefaultNamespace
	}
}

func processNamespaceCmd(sourceData *util.SourceData) {
	if sourceData.Tenant == "" || sourceData.Namespace == "" {
		sourceData.Tenant = utils.PublicTenant
		sourceData.Namespace = utils.DefaultNamespace
	}
}

func processBaseArguments(sourceData *util.SourceData) error {
	processNamespaceCmd(sourceData)

	if sourceData.Name == "" {
		return errors.New("You must specify a name for the source")
	}

	return nil
}
