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

package namespace

import (
	"fmt"
	"io"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func getNamespaceProperties(namespace string) (utils.NameSpaceName, map[string]string, error) {
	ns, err := utils.GetNamespaceName(namespace)
	if err != nil {
		return utils.NameSpaceName{}, nil, err
	}
	admin := cmdutils.NewPulsarClient()
	properties, err := admin.Namespaces().GetProperties(*ns)
	if err != nil {
		return utils.NameSpaceName{}, nil, err
	}
	if properties == nil {
		properties = map[string]string{}
	}
	return *ns, properties, nil
}

func updateNamespaceProperties(ns utils.NameSpaceName, properties map[string]string) error {
	admin := cmdutils.NewPulsarClient()
	return admin.Namespaces().UpdateProperties(ns, properties)
}

func removeNamespaceProperties(ns utils.NameSpaceName) error {
	admin := cmdutils.NewPulsarClient()
	return admin.Namespaces().RemoveProperties(ns)
}

func GetPropertiesCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Get properties of a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    "Get properties of a namespace",
		Command: "pulsarctl namespaces get-properties tenant/namespace",
	})
	desc.CommandExamples = examples
	desc.CommandOutput = append(desc.CommandOutput, ArgError, NsNotExistError)
	desc.CommandOutput = append(desc.CommandOutput, NsErrors...)

	vc.SetDescription(
		"get-properties",
		"Get properties of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)
	vc.EnableOutputFlagSet()
	vc.SetRunFuncWithNameArg(func() error {
		return doGetProperties(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetProperties(vc *cmdutils.VerbCmd) error {
	_, properties, err := getNamespaceProperties(vc.NameArg)
	if err != nil {
		return err
	}
	oc := cmdutils.NewOutputContent().WithObject(properties)
	return vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
}

func SetPropertiesCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Set properties of a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    "Set properties of a namespace",
		Command: "pulsarctl namespaces set-properties tenant/namespace -p k1=v1,k2=v2",
	})
	desc.CommandExamples = examples
	desc.CommandOutput = append(desc.CommandOutput, ArgError, NsNotExistError)
	desc.CommandOutput = append(desc.CommandOutput, NsErrors...)

	vc.SetDescription(
		"set-properties",
		"Set properties of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	properties := map[string]string{}
	vc.FlagSetGroup.InFlagSet("Properties", func(set *pflag.FlagSet) {
		set.StringToStringVarP(&properties, "properties", "p", nil, "comma separated key=value pairs")
		_ = cobra.MarkFlagRequired(set, "properties")
	})

	vc.SetRunFuncWithNameArg(func() error {
		return doSetProperties(vc, properties)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doSetProperties(vc *cmdutils.VerbCmd, properties map[string]string) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}
	err = updateNamespaceProperties(*ns, properties)
	if err == nil {
		vc.Command.Printf("Updated properties successfully for [%s]\n", ns.String())
	}
	return err
}

func ClearPropertiesCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Clear properties of a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    "Clear properties of a namespace",
		Command: "pulsarctl namespaces clear-properties tenant/namespace",
	})
	desc.CommandExamples = examples
	desc.CommandOutput = append(desc.CommandOutput, ArgError, NsNotExistError)
	desc.CommandOutput = append(desc.CommandOutput, NsErrors...)

	vc.SetDescription(
		"clear-properties",
		"Clear properties of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)
	vc.SetRunFuncWithNameArg(func() error {
		return doClearProperties(vc)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doClearProperties(vc *cmdutils.VerbCmd) error {
	ns, err := utils.GetNamespaceName(vc.NameArg)
	if err != nil {
		return err
	}
	err = removeNamespaceProperties(*ns)
	if err == nil {
		vc.Command.Printf("Cleared properties successfully for [%s]\n", ns.String())
	}
	return err
}

func GetPropertyCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Get a single property of a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    "Get a single property of a namespace",
		Command: "pulsarctl namespaces get-property tenant/namespace -k key",
	})
	desc.CommandExamples = examples
	desc.CommandOutput = append(desc.CommandOutput, ArgError, NsNotExistError)
	desc.CommandOutput = append(desc.CommandOutput, NsErrors...)

	vc.SetDescription(
		"get-property",
		"Get a single property of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	var key string
	vc.FlagSetGroup.InFlagSet("Properties", func(set *pflag.FlagSet) {
		set.StringVarP(&key, "key", "k", "", "property key")
		_ = cobra.MarkFlagRequired(set, "key")
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doGetProperty(vc, key)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doGetProperty(vc *cmdutils.VerbCmd, key string) error {
	_, properties, err := getNamespaceProperties(vc.NameArg)
	if err != nil {
		return err
	}
	value, ok := properties[key]
	if !ok {
		return writeNullablePropertyValue(vc, key, nil)
	}
	return writeNullablePropertyValue(vc, key, &value)
}

func SetPropertyCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Set a single property of a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    "Set a single property of a namespace",
		Command: "pulsarctl namespaces set-property tenant/namespace -k key -v value",
	})
	desc.CommandExamples = examples
	desc.CommandOutput = append(desc.CommandOutput, ArgError, NsNotExistError)
	desc.CommandOutput = append(desc.CommandOutput, NsErrors...)

	vc.SetDescription(
		"set-property",
		"Set a single property of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	var key string
	var value string
	vc.FlagSetGroup.InFlagSet("Properties", func(set *pflag.FlagSet) {
		set.StringVarP(&key, "key", "k", "", "property key")
		set.StringVarP(&value, "value", "v", "", "property value")
		_ = cobra.MarkFlagRequired(set, "key")
		_ = cobra.MarkFlagRequired(set, "value")
	})

	vc.SetRunFuncWithNameArg(func() error {
		return doSetProperty(vc, key, value)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doSetProperty(vc *cmdutils.VerbCmd, key, value string) error {
	ns, properties, err := getNamespaceProperties(vc.NameArg)
	if err != nil {
		return err
	}
	properties[key] = value
	err = updateNamespaceProperties(ns, properties)
	if err == nil {
		vc.Command.Printf("Set property %q successfully for [%s]\n", key, ns.String())
	}
	return err
}

func RemovePropertyCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Remove a single property of a namespace"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc:    "Remove a single property of a namespace",
		Command: "pulsarctl namespaces remove-property tenant/namespace -k key",
	})
	desc.CommandExamples = examples
	desc.CommandOutput = append(desc.CommandOutput, ArgError, NsNotExistError)
	desc.CommandOutput = append(desc.CommandOutput, NsErrors...)

	vc.SetDescription(
		"remove-property",
		"Remove a single property of a namespace",
		desc.ToString(),
		desc.ExampleToString(),
	)

	var key string
	vc.FlagSetGroup.InFlagSet("Properties", func(set *pflag.FlagSet) {
		set.StringVarP(&key, "key", "k", "", "property key")
		_ = cobra.MarkFlagRequired(set, "key")
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doRemoveProperty(vc, key)
	}, "the namespace name is not specified or the namespace name is specified more than one")
}

func doRemoveProperty(vc *cmdutils.VerbCmd, key string) error {
	ns, properties, err := getNamespaceProperties(vc.NameArg)
	if err != nil {
		return err
	}
	value, ok := properties[key]
	if !ok {
		return writeNullablePropertyValue(vc, key, nil)
	}
	delete(properties, key)

	if len(properties) == 0 {
		err = removeNamespaceProperties(ns)
	} else {
		err = updateNamespaceProperties(ns, properties)
	}
	if err == nil {
		return writeNullablePropertyValue(vc, key, &value)
	}
	return err
}

func writeNullablePropertyValue(vc *cmdutils.VerbCmd, key string, value *string) error {
	var obj interface{}
	if value != nil {
		obj = map[string]string{key: *value}
	}

	oc := cmdutils.NewOutputContent().
		WithObject(obj).
		WithTextFunc(func(w io.Writer) error {
			if value == nil {
				_, err := io.WriteString(w, "null\n")
				return err
			}
			_, err := fmt.Fprintln(w, *value)
			return err
		})

	return vc.OutputConfig.WriteOutput(vc.Command.OutOrStdout(), oc)
}
