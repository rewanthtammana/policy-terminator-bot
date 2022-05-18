package kubernetes

import (
	"context"

	"github.com/rewanthtammana/policy-terminator/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
)

// Watches for resource updates continously & triggers events
func WatchResource(cliSet dynamic.Interface, ctx context.Context, group string, version string, resource string) watch.Interface {
	watcher, err := cliSet.Resource(schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}).Watch(ctx, metav1.ListOptions{})

	utils.CheckIfError(err)

	return watcher
}
