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

func SearchK8sResourceStatus(namespace, resType, resName string) []*ResourceStatus {
	var (
		rss []*ResourceStatus
	)
	factory := getClient().informerFactory
	switch resType {
	case "pods", "pod":
		if pods, err := factory.Core().V1().Pods().Lister().Pods(namespace).List(labels.Everything()); err == nil {
			for _, pod := range pods {
				if !strings.HasPrefix(pod.Name, resName) {
					continue
				}
				objMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(pod)
				rss = append(rss, &ResourceStatus{
					ResName: pod.Name,
					Status:  objMap["status"],
				})
			}
		}
	case "deployments", "deployment":
		if deployments, err := factory.Apps().V1().Deployments().Lister().Deployments(namespace).List(labels.Everything()); err == nil {
			for _, deployment := range deployments {
				if !strings.HasPrefix(deployment.Name, resName) {
					continue
				}
				objMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(deployment)
				rss = append(rss, &ResourceStatus{
					ResName: deployment.Name,
					Status:  objMap["status"],
				})
			}
		}
	case "daemonsets", "daemonset":
		if daemonsets, err := factory.Apps().V1().DaemonSets().Lister().DaemonSets(namespace).List(labels.Everything()); err == nil {
			for _, daemonset := range daemonsets {
				if !strings.HasPrefix(daemonset.Name, resName) {
					continue
				}
				objMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(daemonset)
				rss = append(rss, &ResourceStatus{
					ResName: daemonset.Name,
					Status:  objMap["status"],
				})
			}
		}
	case "statefulsets", "statefulset":
		if statefuls, err := factory.Apps().V1().StatefulSets().Lister().StatefulSets(namespace).List(labels.Everything()); err == nil {
			for _, stateful := range statefuls {
				if !strings.HasPrefix(stateful.Name, resName) {
					continue
				}
				objMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(stateful)
				rss = append(rss, &ResourceStatus{
					ResName: stateful.Name,
					Status:  objMap["status"],
				})
			}
		}
	case "services", "service":
		if services, err := factory.Core().V1().Services().Lister().Services(namespace).List(labels.Everything()); err == nil {
			for _, service := range services {
				if !strings.HasPrefix(service.Name, resName) {
					continue
				}
				objMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(service)
				rss = append(rss, &ResourceStatus{
					ResName: service.Name,
					Status:  objMap["status"],
				})
			}
		}
	case "jobs", "job":
		if jobs, err := factory.Batch().V1().Jobs().Lister().Jobs(namespace).List(labels.Everything()); err == nil {
			for _, job := range jobs {
				if !strings.HasPrefix(job.Name, resName) {
					continue
				}
				objMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(job)
				rss = append(rss, &ResourceStatus{
					ResName: job.Name,
					Status:  objMap["status"],
				})
			}
		}
	case "cronjobs", "cronjob":
		if cronjobs, err := factory.Batch().V1().CronJobs().Lister().CronJobs(namespace).List(labels.Everything()); err == nil {
			for _, cronjob := range cronjobs {
				if !strings.HasPrefix(cronjob.Name, resName) {
					continue
				}
				objMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(cronjob)
				rss = append(rss, &ResourceStatus{
					ResName: cronjob.Name,
					Status:  objMap["status"],
				})
			}
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
