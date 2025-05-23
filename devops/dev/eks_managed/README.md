## To fix Error: configmaps "aws-auth" already exists

Use the `terraform import` command to import the existing `aws-auth` ConfigMap into your Terraform state. This will make Terraform aware of the existing resource and manage it correctly

## Kubectl Configuration

```
terraform import module.eks.kubernetes_config_map.aws_auth_ignore_changes\[0\] kube-system/aws-auth
```

- Ensure kubectl is configured correctly:

```
kubectl config view
```

- Update your kubeconfig to include the EKS cluster.

```
aws eks --region <region> update-kubeconfig --name <cluster-name>
```

## Context and Namespace Management

- Check cluster info:

```
kubectl cluster-info
```

- List all contexts.

```
kubectl config get-contexts
```

- Switch to a different context.

```
kubectl config use-context <context-name>
```

- Set the default namespace for the current context.

```
kubectl config set-context --current --namespace=<namespace>
```

- Setting a specific context:

```
kubectl config set-context my-cluster --cluster=my-cluster --user=my-user
```

- Listing clusters, users, and contexts:

```
kubectl config get-clusters
kubectl config get-users
kubectl config get-contexts
```

## Working with Nodes

- List all nodes:

```
kubectl get nodes
```

- Get detailed information about a node:

```
kubectl describe node <node-name>
```

- Label a node:

```
kubectl label node <node-name> <label-key>=<label-value>
```

- Remove a label from a node:

```
kubectl label node <node-name> <label-key>-
```

## Working with Pods

- List all pods in a namespace:

```
kubectl get pods -n <namespace>
```

- Get detailed information about a pod:

```
kubectl describe pod <pod-name> -n <namespace>
```

- Delete a pod:

```
kubectl delete pod <pod-name> -n <namespace>
```

- Execute a command in a pod:

```
kubectl exec -it <pod-name> -n <namespace> -- <command>
```

## Working with Deployments

- List all deployments in a namespace:

```
kubectl get deployments -n <namespace>
```

- Get detailed information about a deployment:

```
kubectl describe deployment <deployment-name> -n <namespace>
```

## Working with Services

- List all services in a namespace:

```
kubectl get services -n <namespace>
```

- Get detailed information about a service:

```
kubectl describe service <service-name> -n <namespace>
```

## Working with ConfigMaps and Secrets

- List all ConfigMaps in a namespace:

```
kubectl get configmaps -n <namespace>
```

- Get detailed information about a ConfigMap:

```
kubectl describe configmap <configmap-name> -n <namespace>
```

- kubectl get secrets -n <namespace>

```
kubectl get secrets -n <namespace>
```

## Working with Namespaces

- List all namespaces:

```
kubectl get namespaces
```

- Create a new namespace:

```
kubectl create namespace <namespace-name>
```

- Delete a namespace:

```
kubectl delete namespace <namespace-name>
```

## Logs and Events

- View logs of a pod:

```
kubectl logs <pod-name> -n <namespace>
```

- Stream logs of a pod:

```
kubectl logs -f <pod-name> -n <namespace>
```

- View events in a namespace:

```
kubectl get events -n <namespace>
```

## Apply and Delete Configurations

- Apply a configuration file:

```
kubectl apply -f <file-path>
```
- Delete resources from a configuration file:

```
kubectl delete -f <file-path>
```