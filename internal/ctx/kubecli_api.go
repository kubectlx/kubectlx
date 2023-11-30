package ctx

import (
	"k8s.io/apimachinery/pkg/labels"
	"strings"
)

func GetAllNamespace() []string {
	var namespaceNames []string
	namespaces, err := client.informerFactory.Core().V1().Namespaces().Lister().List(labels.Everything())
	if err != nil {
		return namespaceNames
	}
	for _, ns := range namespaces {
		namespaceNames = append(namespaceNames, ns.Name)
	}
	return namespaceNames
}

func GetPods(namePrefix string, limit int) []string {
	var podNames []string
	pods, err := client.informerFactory.Core().V1().Pods().Lister().Pods(GetNamespace()).List(labels.Everything())
	if err != nil {
		return podNames
	}
	cnt := 0
	for _, pod := range pods {
		if !strings.HasPrefix(pod.Name, namePrefix) {
			continue
		}
		podNames = append(podNames, pod.Name)
		cnt++
		if cnt == limit {
			break
		}
	}
	return podNames
}
