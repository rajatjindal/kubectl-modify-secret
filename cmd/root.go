// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/rajatjindal/kubectl-modify-secret/pkg/editor"
	"github.com/rajatjindal/kubectl-modify-secret/pkg/secrets"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	namespace      string
	kubeconfigFile string
	debug          bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-modify-secret",
	Short: "kubectl-modify-secret allows user to directly modify the secret without worrying about base64 encoding/decoding",
	Run: func(cmd *cobra.Command, args []string) {
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		logrus.Debugf("args: %s", args)
		logrus.Debugf("environ: %s", os.Environ())

		if len(args) == 0 {
			logrus.Fatal("no secret name provided")
		}

		kubeclient, err := getKubernetesClient(kubeconfigFile)
		if err != nil {
			logrus.Fatal(err)
		}

		err = modify(kubeclient, args[0], namespace)
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "namespace of the secret")
	rootCmd.Flags().StringVar(&kubeconfigFile, clientcmd.RecommendedConfigPathFlag, "", "kubeconfig file")
	rootCmd.Flags().BoolVar(&debug, "debug", false, "print debug logs")
}

func modify(kubeclient *kubernetes.Clientset, name, namespace string) error {
	secret, err := secrets.Get(kubeclient, name, namespace)
	if err != nil {
		return err
	}

	data := make(map[string]string, len(secret.Data))
	for k, v := range secret.Data {
		data[k] = string(v)
	}

	tempfile, err := ioutil.TempFile("", fmt.Sprintf("%s-%s-", namespace, name))
	if err != nil {
		return err
	}
	defer os.Remove(tempfile.Name())

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	logrus.Info(string(yamlData))
	err = ioutil.WriteFile(tempfile.Name(), yamlData, 0644)
	if err != nil {
		return err
	}

	err = editor.Edit(tempfile.Name())
	if err != nil {
		return err
	}

	readData, err := ioutil.ReadFile(tempfile.Name())
	if err != nil {
		return err
	}

	var updateData map[string]string
	err = yaml.Unmarshal(readData, &updateData)
	if err != nil {
		return err
	}

	updateByteData := make(map[string][]byte, len(updateData))
	for k, v := range updateData {
		updateByteData[k] = []byte(v)
	}

	secret.Data = updateByteData

	_, err = secrets.Update(kubeclient, secret)
	if err != nil {
		return err
	}

	return nil
}

func getKubernetesClient(kubeconfig string) (*kubernetes.Clientset, error) {
	var config *rest.Config
	var err error

	if kubeconfig != "" {
		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			return nil, err
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
