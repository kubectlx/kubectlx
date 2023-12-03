package kubecli

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func GetCrdThisGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "apiextensions.k8s.io",
		Version:  "v1",
		Resource: "customresourcedefinitions",
	}
}

func GetCrdGroupVersionResource(crd *unstructured.Unstructured) schema.GroupVersionResource {
	crdMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(crd)
	name := crdMap["spec"].(map[string]interface{})["names"].(map[string]interface{})["plural"].(string)
	group := crdMap["spec"].(map[string]interface{})["group"].(string)
	version := crdMap["spec"].(map[string]interface{})["versions"].([]interface{})[0].(map[string]interface{})["name"].(string)
	return schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: name,
	}
}

func GetCrdGroupVersionResourceWithObject(crd runtime.Object) schema.GroupVersionResource {
	crdMap, _ := runtime.DefaultUnstructuredConverter.ToUnstructured(crd)
	name := crdMap["spec"].(map[string]interface{})["names"].(map[string]interface{})["plural"].(string)
	group := crdMap["spec"].(map[string]interface{})["group"].(string)
	version := crdMap["spec"].(map[string]interface{})["versions"].([]interface{})[0].(map[string]interface{})["name"].(string)
	return schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: name,
	}
}
