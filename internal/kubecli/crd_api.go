package kubecli

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"strings"
)

func GetCrdCommand(regx string) []*command.Param {
	var (
		crds []*command.Param
	)
	crdList, err := getClient().dynamicInformerFactory.ForResource(GetCrdThisGroupVersionResource()).Lister().List(labels.Everything())
	if err != nil {
		return nil
	}
	for _, crd := range crdList {
		gvr := GetCrdGroupVersionResourceWithObject(crd)
		crds = append(crds, &command.Param{
			Name:        gvr.Resource,
			Description: fmt.Sprintf(regx, gvr.Resource),
			Extended: map[string]string{
				"group":   gvr.Group,
				"version": gvr.Version,
			},
		})
	}
	return crds
}

func GetCrdDefinitions(input string, limit int) []*command.Param {
	var (
		crds []*command.Param
	)
	crdList, err := getClient().dynamicInformerFactory.ForResource(GetCrdThisGroupVersionResource()).Lister().List(labels.Everything())
	if err != nil {
		return nil
	}
	count := 0
	for _, crd := range crdList {
		crdMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(&crd)
		name := crdMap["metadata"].(map[string]interface{})["name"].(string)
		if !strings.HasPrefix(name, input) {
			continue
		}
		crds = append(crds, &command.Param{
			Name:        name,
			Description: "",
		})
		count++
		if count > limit {
			break
		}
	}
	return crds
}

func GetCrdResource(crdGroup, crdVersion, crdName, namespace, namePrefix string, limit int) []*command.Param {
	var (
		resources []*command.Param
	)
	crdList, err := getClient().dynamicInformerFactory.ForResource(schema.GroupVersionResource{
		Group:    crdGroup,
		Version:  crdVersion,
		Resource: crdName,
	}).Lister().ByNamespace(namespace).List(labels.Everything())
	if err != nil {
		return nil
	}
	cnt := 0
	for _, crd := range crdList {
		crdMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(&crd)
		name := crdMap["metadata"].(map[string]interface{})["name"].(string)
		if !strings.HasPrefix(name, namePrefix) {
			continue
		}
		resources = append(resources, &command.Param{
			Name:        name,
			Description: crdMap["metadata"].(map[string]interface{})["creationTimestamp"].(string),
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return resources
}
