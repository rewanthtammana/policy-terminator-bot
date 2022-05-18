package kubernetes

import (
	"context"
	"fmt"

	"github.com/rewanthtammana/policy-terminator/slack"
	"github.com/rewanthtammana/policy-terminator/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/client-go/dynamic"
)

func DeleteResource(cliSet dynamic.Interface, ctx context.Context, group string, version string, resource string, namespace string, name string) {

	// deletePolicy := metav1.DeletePropagationForeground
	// deleteOptions := metav1.DeleteOptions{
	// 	PropagationPolicy: &deletePolicy,
	// }

	// Deletes the specificed resource
	err := cliSet.Resource(schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}).Namespace(namespace).Delete(ctx, name, metav1.DeleteOptions{})

	// // Lists available resources
	// out, err := cliSet.Resource(schema.GroupVersionResource{
	// 	Group:    group,
	// 	Version:  version,
	// 	Resource: resource,
	// }).Namespace(namespace).List(ctx, metav1.ListOptions{})

	// utils.CheckErrorsAndNotify(err, fmt.Sprintf("Terminating %s/%s in %s namespace", resource, name, namespace))
	isError := utils.CheckIfErrorAndLog(err)
	if !isError {
		slack.NotifyUser(fmt.Sprintf("Terminating %s/%s in %s namespace", resource, name, namespace))
	}
}
