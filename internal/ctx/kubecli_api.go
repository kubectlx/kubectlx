package ctx

import (
	"github.com/cxweilai/kubectlx/internal/command"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"strings"
)

func GetNamespaces() []*command.Param {
	var namespaceNames []*command.Param
	namespaces, err := client.informerFactory.Core().V1().Namespaces().
		Lister().List(labels.Everything())
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

func GetPods(namePrefix string, limit int) []*command.Param {
	var (
		podNames []*command.Param
		pods     []*v1.Pod
		err      error
	)
	podList := client.informerFactory.Core().V1().Pods().Lister()
	pods, err = podList.Pods(GetNamespace()).List(labels.Everything())
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
