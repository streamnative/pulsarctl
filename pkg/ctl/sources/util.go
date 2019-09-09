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
    `encoding/json`
    `fmt`
    `github.com/pkg/errors`
    `github.com/streamnative/pulsarctl/pkg/ctl/utils`
    `github.com/streamnative/pulsarctl/pkg/pulsar`
    `gopkg.in/yaml.v2`
    `io/ioutil`
    `os`
    `strings`
)

func processArguments(sourceData *pulsar.SourceData) error {
    // Initialize config builder either from a supplied YAML config file or from scratch
    if sourceData.SourceConf != nil {
        // no-op
    } else {
        sourceData.SourceConf = new(pulsar.SourceConfig)
    }

    if sourceData.SourceConfigFile != "" {
        yamlFile, err := ioutil.ReadFile(sourceData.SourceConfigFile)
        if err == nil {
            err = yaml.Unmarshal(yamlFile, sourceData.SourceConf)
            if err != nil {
                return fmt.Errorf("unmarshal yaml file error:%s", err.Error())
            }
        } else if err != nil && !os.IsNotExist(err) {
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
            sourceData.SourceConf.Resources = pulsar.NewDefaultResources()
        }

        sourceData.SourceConf.Resources.CPU = sourceData.CPU
    }

    if sourceData.Disk != 0 {
        if sourceData.SourceConf.Resources == nil {
            sourceData.SourceConf.Resources = pulsar.NewDefaultResources()
        }

        sourceData.SourceConf.Resources.Disk = sourceData.Disk
    }

    if sourceData.RAM != 0 {
        if sourceData.SourceConf.Resources == nil {
            sourceData.SourceConf.Resources = pulsar.NewDefaultResources()
        }

        sourceData.SourceConf.Resources.Ram = sourceData.RAM
    }

    if sourceData.SourceConfigString != "" {
        sourceData.SourceConf.Configs = parseConfigs(sourceData.SourceConfigString)
    }

    return nil
}

func validateSourceType(sourceType string) string {

    return ""
}

func parseConfigs(str string) map[string]interface{} {
    var resMap map[string]interface{}

    err := json.Unmarshal([]byte(str), &resMap)
    if err != nil {
        return nil
    }

    return resMap
}

func validateSourceConfigs(sourceConfig *pulsar.SourceConfig) error {
    if sourceConfig.Archive == "" {
        return errors.New("Source archive not specified")
    }

    utils.InferMissingSourceArguments(sourceConfig)

    if utils.IsPackageUrlSupported(sourceConfig.Archive) && strings.HasPrefix(sourceConfig.Archive, utils.BUILTIN) {
        if !utils.IsFileExist(sourceConfig.Archive) {
            return fmt.Errorf("source Archive %s does not exist", sourceConfig.Archive)
        }
    }

    if sourceConfig.Name == "" {
        return errors.New("source name not specified")
    }
    return nil
}

func checkArgsForUpdate(sourceConfig *pulsar.SourceConfig) {
    if sourceConfig.Tenant == "" {
        sourceConfig.Tenant = utils.PublicTenant
    }

    if sourceConfig.Namespace == "" {
        sourceConfig.Namespace = utils.DefaultNamespace
    }
}

func processNamespaceCmd(sourceData *pulsar.SourceData) {
    if sourceData.Tenant == "" || sourceData.Namespace == "" {
        sourceData.Tenant = utils.PublicTenant
        sourceData.Namespace = utils.DefaultNamespace
    }
}

func processBaseArguments(sourceData *pulsar.SourceData) error {
    processNamespaceCmd(sourceData)

    if sourceData.Name == "" {
        return errors.New("You must specify a name for the source")
    }

    return nil
}
