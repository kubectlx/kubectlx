package kubecli

import (
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"strings"
)

type ResourceStatus struct {
	ResName string
	Status  interface{}
}

func SearchK8sResourceStatus(namespace, group, version, name string, resName string) []*ResourceStatus {
	var (
		rss  []*ResourceStatus
		objs []runtime.Object
		err  error
	)
	factory := getClient().informerFactory
	infr, err := factory.ForResource(schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: name,
	})
	if err != nil {
		return rss
	}
	if name == "nodes" || name == "namespaces" {
		objs, err = infr.Lister().List(labels.Everything())
	} else {
		objs, err = infr.Lister().ByNamespace(namespace).List(labels.Everything())
	}
	if err == nil {
		for _, obj := range objs {
			crdMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
			rName := crdMap["metadata"].(map[string]interface{})["name"].(string)
			if !strings.HasPrefix(rName, resName) {
				continue
			}
			rss = append(rss, &ResourceStatus{
				ResName: rName,
				Status:  crdMap["status"],
			})
		}
	}
	return rss
}

func SearchCrdResourceStatus(namespace, group, version, name string, resName string) []*ResourceStatus {
	var (
		rss []*ResourceStatus
	)
	factory := getClient().dynamicInformerFactory
	if objs, err := factory.ForResource(schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: name,
	}).Lister().ByNamespace(namespace).List(labels.Everything()); err == nil {
		for _, obj := range objs {
			crdMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(&obj)
			rName := crdMap["metadata"].(map[string]interface{})["name"].(string)
			if !strings.HasPrefix(rName, resName) {
				continue
			}
			rss = append(rss, &ResourceStatus{
				ResName: rName,
				Status:  crdMap["status"],
			})
		}
	}
	return rss
}
