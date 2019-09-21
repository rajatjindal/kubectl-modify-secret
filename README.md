# kubectl-modify-secret
kubectl-modify-secret allows user to directly modify the secret without worrying about base64 encoding/decoding

once installed as plugin, you can run it as follows:

`kubectl modify secret secret-name -n kube-system --kubeconfig /path/to/kube/config`

It will open vim editor, and you can edit the secrets and upon saving and exiting the vim, the changes will be applied to kubernetes.
