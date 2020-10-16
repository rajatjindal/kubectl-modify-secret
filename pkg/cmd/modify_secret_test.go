package cmd

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/fake"
)

func TestModifySecrets(t *testing.T) {
	const (
		name      = "mysecret"
		namespace = "mynamespace"
	)

	logrus.SetOutput(ioutil.Discard)
	origEditor := os.Getenv("EDITOR")
	if origEditor != "" {
		defer os.Setenv("EDITOR", origEditor)
	} else {
		defer os.Unsetenv("EDITOR")
	}

	testcases := []struct {
		name     string
		command  string
		secret   map[string][]byte
		expected map[string][]byte
	}{
		{
			name:     "no changes to empty secret",
			command:  "touch",
			secret:   map[string][]byte{},
			expected: map[string][]byte{},
		}, {
			name:     "no changes to populated secret",
			command:  "touch",
			secret:   map[string][]byte{"key": []byte("value")},
			expected: map[string][]byte{"key": []byte("value")},
		}, {
			name:     "changes to populated secret",
			command:  "sed -i= s/value/updated/",
			secret:   map[string][]byte{"key": []byte("value")},
			expected: map[string][]byte{"key": []byte("updated")},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("EDITOR", tc.command)

			client := fake.NewSimpleClientset(&v1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:        name,
					Namespace:   namespace,
					Annotations: map[string]string{},
				},
				Data: tc.secret,
			})

			modify := ModifySecretOptions{
				kubeclient: client,
				secretName: name,
				namespace:  namespace,
			}
			require.NoError(t, modify.Run())

			object, err := client.Tracker().Get(
				schema.GroupVersionResource{
					Version:  "v1",
					Resource: "secrets",
				},
				namespace, name,
			)
			require.NoError(t, err)
			assert.Equal(t, tc.expected, object.(*v1.Secret).Data)
		})
	}
}
