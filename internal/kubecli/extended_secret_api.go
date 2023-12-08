package kubecli

import (
	"encoding/base64"
	"k8s.io/apimachinery/pkg/runtime"
)

func GetSecretAndBase64Decode(namespace, resName string) (map[string]interface{}, error) {
	var (
		secretMap map[string]interface{}
	)
	secret, err := getClient().informerFactory.Core().V1().Secrets().Lister().Secrets(namespace).Get(resName)
	if err != nil {
		return nil, err
	}
	secretMap, _ = runtime.DefaultUnstructuredConverter.ToUnstructured(&secret)
	metadataMap := secretMap["metadata"].(map[string]interface{})
	delete(metadataMap, "managedFields")
	secretMap["metadata"] = metadataMap
	return decodeSecretData(secretMap), err
}

func decodeSecretData(secretMap map[string]interface{}) map[string]interface{} {
	if dataMap, ok := secretMap["data"]; ok {
		curD := dataMap.(map[string]interface{})
		newDataMap := map[string]interface{}{}
		for key := range curD {
			if strV, ok := curD[key].(string); ok {
				if v, err := base64.StdEncoding.DecodeString(strV); err != nil {
					newDataMap[key] = "parserErr ==> " + strV
				} else {
					newDataMap[key] = string(v)
				}
			}
		}
		secretMap["data"] = newDataMap
	}
	return secretMap
}
