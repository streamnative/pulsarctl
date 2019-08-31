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

type  ProcessingGuarantees int

type Runtime int

const (
    AtLeasetOnce ProcessingGuarantees = iota
    AtMostOnce
    EffectivelyOnce
)

const (
    Java Runtime = iota
    Python
    Go
)

type FunctionConfig struct {
    // Any flags that you want to pass to the runtime.
    // note that in thread mode, these flags will have no impact
    RuntimeFlags string `json:"runtimeFlags"`

    Tenant    string `json:"tenant"`
    Namespace string `json:"namespace"`
    Name      string `json:"name"`
    ClassName string `json:"className"`

    Inputs             []string          `json:"inputs"`
    CustomSerdeInputs  map[string]string `json:"customSerdeInputs"`
    TopicsPattern      *string           `json:"topicsPattern"`
    CustomSchemaInputs map[string]string `json:"customSchemaInputs"`

    // A generalized way of specifying inputs
    InputSpecs map[string]ConsumerConfig `json:"inputSpecs"`

    Output string `json:"output"`

    // Represents either a builtin schema type (eg: 'avro', 'json', ect) or the class name for a Schema implementation
    OutputSchemaType string `json:"outputSchemaType"`

    OutputSerdeClassName string                 `json:"outputSerdeClassName"`
    LogTopic             string                 `json:"logTopic"`
    ProcessingGuarantees ProcessingGuarantees   `json:"processingGuarantees"`
    RetainOrdering       bool                   `json:"retainOrdering"`
    UserConfig           map[string]interface{} `json:"userConfig"`

    // This is a map of secretName(aka how the secret is going to be
    // accessed in the function via context) to an object that
    // encapsulates how the secret is fetched by the underlying
    // secrets provider. The type of an value here can be found by the
    // SecretProviderConfigurator.getSecretObjectType() method.
    Secrets map[string]interface{} `json:"secrets"`

    Runtime           Runtime       `json:"runtime"`
    AutoAck           bool          `json:"autoAck"`
    MaxMessageRetries int           `json:"maxMessageRetries"`
    DeadLetterTopic   string        `json:"deadLetterTopic"`
    SubName           string        `json:"subName"`
    Parallelism       int           `json:"parallelism"`
    Resources         *Resources    `json:"resources"`
    FQFN              string        `json:"fqfn"`
    WindowConfig      *WindowConfig `json:"windowConfig"`
    TimeoutMs         *int64         `json:"timeoutMs"`
    Jar               string        `json:"jar"`
    Py                string        `json:"py"`
    Go                string        `json:"go"`
    // Whether the subscriptions the functions created/used should be deleted when the functions is deleted
    CleanupSubscription bool `json:"cleanupSubscription"`
}
