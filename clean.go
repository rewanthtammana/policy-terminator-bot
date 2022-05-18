package main

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	kubernetes "github.com/rewanthtammana/policy-terminator/kubernetes"
	"github.com/tidwall/gjson"
	"k8s.io/client-go/dynamic"
	ctrl "sigs.k8s.io/controller-runtime"
)

// Main function to kick-start the bot
func main() {
	ctx := context.Background()
	config := ctrl.GetConfigOrDie()
	cliSet := dynamic.NewForConfigOrDie(config)

	// Watches for the below resource
	group := "wgpolicyk8s.io"
	version := "v1alpha2"
	resource := "policyreports"
	watchKyvernoPolicyReports := kubernetes.WatchResource(cliSet, ctx, group, version, resource)

	// Analyzing all events sent by watcher
	for event := range watchKyvernoPolicyReports.ResultChan() {
		// utils.FormatToJSON(event.Object)
		eventObject, _ := json.MarshalIndent(event.Object, "", "\t")

		// Extracts list of all resources that failed compliance
		allResourcesList := gjson.Get(string(eventObject), "results.#(result==\"fail\")#")

		// Iterate each resource to validate for resource type & extract required fields
		allResourcesList.ForEach(func(key, value gjson.Result) bool {
			result := gjson.GetMany(string(value.String()), "message", "policy", "resources", "severity")
			result[2].ForEach(func(key, value gjson.Result) bool {
				result := gjson.GetMany(value.String(), "apiVersion", "kind", "name", "namespace")

				// Performs operations on apiVersion field to extract group & version information for future operations
				apiVersion := result[0].String()
				regexPattern := regexp.MustCompile(`([^/]*)/?(.*)`)
				regexMatches := regexPattern.FindStringSubmatch(apiVersion)
				var group string
				var version string
				if regexMatches[2] == "" {
					group = ""
					version = regexMatches[1]
				} else {
					group = regexMatches[1]
					version = regexMatches[2]
				}
				name := result[2].String()
				namespace := result[3].String()

				kind := strings.ToLower(result[1].String())
				// if namespace == "loki" || namespace == "lala" {
				if kind == "deployment" {
					fmt.Println("XGroup = ", group, "Version = ", version, "kind = ", kind, "namespace = ", namespace, "name = ", name)
					kubernetes.DeleteResource(cliSet, ctx, group, version, "deployments", namespace, name)
				} else if kind == "pod" {
					fmt.Println("YGroup = ", group, "Version = ", version, "kind = ", kind, "namespace = ", namespace, "name = ", name)
					kubernetes.DeleteResource(cliSet, ctx, group, version, "pods", namespace, name)
				}
				// Can be expanded to daemonset, repliaset, statefulset....
				// }
				return true
			})
			return true
		})
	}
}
