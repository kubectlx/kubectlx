package kubecli

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"strings"
)

var k8sGVRs = []schema.GroupVersionResource{
	{
		Group:    "",
		Version:  "v1",
		Resource: "namespaces",
	},
	{
		Group:    "",
		Version:  "v1",
		Resource: "nodes",
	},
	{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	},
	{
		Group:    "",
		Version:  "v1",
		Resource: "services",
	},
	{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	},
	{
		Group:    "apps",
		Version:  "v1",
		Resource: "daemonsets",
	},
	{
		Group:    "apps",
		Version:  "v1",
		Resource: "replicasets",
	},
	{
		Group:    "apps",
		Version:  "v1",
		Resource: "statefulsets",
	},
	{
		Group:    "batch",
		Version:  "v1",
		Resource: "jobs",
	},
	{
		Group:    "batch",
		Version:  "v1",
		Resource: "cronjobs",
	},
	{
		Group:    "",
		Version:  "v1",
		Resource: "configmaps",
	},
	{
		Group:    "",
		Version:  "v1",
		Resource: "secrets",
	},
}

func IsK8sResource(group, resource string) bool {
	for _, gvr := range k8sGVRs {
		if gvr.Group == group && gvr.Resource == resource {
			return true
		}
	}
	return false
}

func GetK8sResourceCommand(regx string) []*command.Param {
	var ress []*command.Param
	for _, gvr := range k8sGVRs {
		ress = append(ress, &command.Param{
			Name:        gvr.Resource,
			Description: fmt.Sprintf(regx, gvr.Resource),
			Extended: map[string]string{
				"group":   gvr.Group,
				"version": gvr.Version,
			},
		})
	}
	return ress
}

func GetK8sResource(group, version, resource, namespace, namePrefix string, limit int) []*command.Param {
	var (
		result []*command.Param
		objs   []runtime.Object
		err    error
	)
	infr, err := getClient().informerFactory.ForResource(schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	})
	if err != nil {
		return result
	}
	cnt := 0
	if resource == "nodes" || resource == "namespaces" {
		objs, err = infr.Lister().List(labels.Everything())
	} else {
		objs, err = infr.Lister().ByNamespace(namespace).List(labels.Everything())
	}
	if err == nil {
		for _, obj := range objs {
			objMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
			name := objMap["metadata"].(map[string]interface{})["name"].(string)
			if !strings.HasPrefix(name, namePrefix) {
				continue
			}
			desc, _ := objMap["metadata"].(map[string]interface{})["creationTimestamp"].(string)
			switch resource {
			case "secrets":
				desc, _ = objMap["type"].(string)
			case "cronjobs":
				desc, _ = objMap["spec"].(map[string]interface{})["schedule"].(string)
			case "jobs":
				if conditions, ok := objMap["status"].(map[string]interface{})["conditions"]; ok {
					if cs, ok := conditions.([]interface{}); ok && len(cs) > 0 {
						desc = cs[0].(map[string]interface{})["type"].(string)
					}
				}
			case "deployments", "replicasets", "statefulsets":
				availableReplicas, ok := objMap["status"].(map[string]interface{})["availableReplicas"]
				if !ok {
					availableReplicas = 0
				}
				desc = fmt.Sprintf("%d/%d", availableReplicas,
					objMap["status"].(map[string]interface{})["replicas"])
			case "nodes", "namespaces", "pods":
				desc, _ = objMap["status"].(map[string]interface{})["phase"].(string)
			case "daemonsets":
				numberAvailable, ok := objMap["status"].(map[string]interface{})["numberAvailable"]
				if !ok {
					numberAvailable = 0
				}
				desc = fmt.Sprintf("%d", numberAvailable)
			}
			result = append(result, &command.Param{
				Name:        name,
				Description: desc,
			})
			cnt++
			if cnt == limit {
				break
			}
		}
	}
	return result
}

func GetNamespaces() []*command.Param {
	var (
		namespaceNames []*command.Param
	)
	namespaces, err := getClient().informerFactory.Core().V1().Namespaces().Lister().List(labels.Everything())
	if err != nil {
		return namespaceNames
	}
	for _, ns := range namespaces {
		namespaceNames = append(namespaceNames, &command.Param{
			Name:        ns.Name,
			Description: string(ns.Status.Phase),
		})
	}
	return namespaceNames
}

func GetPods(namespace, namePrefix string, limit int) []*command.Param {
	var (
		podNames []*command.Param
		pods     []*corev1.Pod
		err      error
	)
	pods, err = getClient().informerFactory.Core().V1().Pods().Lister().Pods(namespace).List(labels.Everything())
	if err != nil {
		return podNames
	}
	cnt := 0
	for _, pod := range pods {
		if !strings.HasPrefix(pod.Name, namePrefix) {
			continue
		}
		podNames = append(podNames, &command.Param{
			Name:        pod.Name,
			Description: string(pod.Status.Phase),
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return podNames
}
