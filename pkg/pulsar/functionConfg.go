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
    RuntimeFlags string

    Tenant    string
    Namespace string
    Name      string
    ClassName string

    Inputs             []string
    CustomSerdeInputs  map[string]string
    TopicsPattern      string
    CustomSchemaInputs map[string]string

    // A generalized way of specifying inputs
    InputSpecs map[string]ConsumerConfig

    Output string

    // Represents either a builtin schema type (eg: 'avro', 'json', ect) or the class name for a Schema implementation
    OutputSchemaType string

    OutputSerdeClassName string
    LogTopic             string
    ProcessingGuarantees ProcessingGuarantees
    RetainOrdering       bool
    UserConfig           map[string]interface{}

    // This is a map of secretName(aka how the secret is going to be
    // accessed in the function via context) to an object that
    // encapsulates how the secret is fetched by the underlying
    // secrets provider. The type of an value here can be found by the
    // SecretProviderConfigurator.getSecretObjectType() method.
    Secrets map[string]interface{}

    Runtime           Runtime
    AutoAck           bool
    MaxMessageRetries int
    DeadLetterTopic   string
    SubName           string
    Parallelism       int
    Resources         *Resources
    FQFN              string
    WindowConfig      *WindowConfig
    TimeoutMs         int64
    Jar               string
    Py                string
    Go                string
    // Whether the subscriptions the functions created/used should be deleted when the functions is deleted
    CleanupSubscription bool
}
