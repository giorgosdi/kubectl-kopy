# kubectl kopy

kubectl kopy will copy a resource with the exact same configuration without the generated fields (selfLink, uid etc)


## Installation

You can install this plugin by either downloading the corresponding asset from the latest release or from source by cloning the repository and build the project in the root directory.

```
go build .
```

`make` rules can also be used to generate all assets for all architectures (arm64, amd64) by running `make build`.

To remove all generated files use `make clean`.

A PR in [krew-index](https://github.com/kubernetes-sigs/krew-index) will be sumbitted shortly and you will be able to install directly from the krew plugin.

`kubectl krew install kopy`

## Usage

The plugin takes two arguments, the kubernetes object and the name of that object. It also takes two flags, the source namespace of the object which uses the conventional `kubectl` flag `--namespace` and the destination namespace `--target`

```
kubectl kopy secret my-secret -n kube-system --target monitoring
```

## Upcoming features

* Copy a namespace (the single object) and create a new namespace with the `--target` flag
* Add functionality to keep certain labels, annotations and any other generated field in the new object
* Support objects that are not namespaced like ClusterRoles etc.

## Similar plugins

A similar plugin you could use is [kubectl neat](https://github.com/itaysk/kubectl-neat). This plugin also removes the clutter kubernetes objects but instead of copying to another namespace, it will give you the neated manifest in your standard ouput. `kubectl neat` also neats local files.

## LICENSE

Apache 2.0. See [LICENSE](./LICENSE).