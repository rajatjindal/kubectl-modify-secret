package secrets

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Get gets the secret from Kubernetes
func Get(ctx context.Context, kubeclient kubernetes.Interface, name, namespace string) (*v1.Secret, error) {
	return kubeclient.CoreV1().Secrets(namespace).Get(ctx, name, metav1.GetOptions{})
}

// Update updates the secret to Kubernetes
func Update(ctx context.Context, kubeclient kubernetes.Interface, secret *v1.Secret) (*v1.Secret, error) {
	return kubeclient.CoreV1().Secrets(secret.Namespace).Update(ctx, secret, metav1.UpdateOptions{})
}
