package kubecli

import (
	"context"
	"errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"sync"
	"time"
)

type clientHolder struct {
	client                 *kubernetes.Clientset
	informerFactory        informers.SharedInformerFactory
	dynamicClient          *dynamic.DynamicClient
	dynamicInformerFactory dynamicinformer.DynamicSharedInformerFactory
	kubeconfig             string
}

var clientMap sync.Map
var _kubeconfig string

func getClient() *clientHolder {
	if client, ok := clientMap.Load(_kubeconfig); ok {
		return client.(*clientHolder)
	} else {
		panic(errors.New("system error. client not found:" + _kubeconfig))
	}
}

func init() {
	clientMap = sync.Map{}
}

func newClientSet(kubeconfig string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	if clientSet, err := kubernetes.NewForConfig(config); err != nil {
		return nil, err
	} else if _, err := clientSet.ServerVersion(); err != nil {
		return nil, err
	} else {
		return clientSet, nil
	}
}

func newDynamicClient(kubeconfig string) (*dynamic.DynamicClient, error) {
	if config, err := clientcmd.BuildConfigFromFlags("", kubeconfig); err != nil {
		return nil, err
	} else if dynamicClient, err := dynamic.NewForConfig(config); err != nil {
		return nil, err
	} else {
		return dynamicClient, nil
	}
}

func InitKubeClient(kubeconfig string) error {
	if _, ok := clientMap.Load(kubeconfig); ok {
		_kubeconfig = kubeconfig
		return nil
	}

	client := &clientHolder{
		kubeconfig: kubeconfig,
	}
	if c, err := newClientSet(kubeconfig); err != nil {
		return err
	} else {
		client.client = c
	}
	if dc, err := newDynamicClient(kubeconfig); err != nil {
		return err
	} else {
		client.dynamicClient = dc
	}
	initInformerCache(client)
	initDynamicInformerCache(client)

	clientMap.Store(kubeconfig, client)
	_kubeconfig = kubeconfig
	return nil
}

func initInformerCache(c *clientHolder) {
	c.informerFactory = informers.NewSharedInformerFactory(c.client, time.Second*600)
	c.informerFactory.Core().V1().Nodes().Informer()
	c.informerFactory.Core().V1().Pods().Informer()
	c.informerFactory.Apps().V1().Deployments().Informer()
	c.informerFactory.Core().V1().Services().Informer()
	c.informerFactory.Core().V1().Namespaces().Informer()
	c.informerFactory.Apps().V1().DaemonSets().Informer()
	c.informerFactory.Apps().V1().ReplicaSets().Informer()
	c.informerFactory.Apps().V1().StatefulSets().Informer()
	c.informerFactory.Batch().V1().Jobs().Informer()
	c.informerFactory.Batch().V1().CronJobs().Informer()
	c.informerFactory.Core().V1().Events().Informer()
	c.informerFactory.Core().V1().ConfigMaps().Informer()
	c.informerFactory.Core().V1().Secrets().Informer()
	// 等待缓存同步完成
	c.informerFactory.Start(wait.NeverStop)
	c.informerFactory.WaitForCacheSync(wait.NeverStop)
}

func initDynamicInformerCache(c *clientHolder) {
	c.dynamicInformerFactory = dynamicinformer.NewDynamicSharedInformerFactory(c.dynamicClient, time.Second*600)
	c.dynamicInformerFactory.ForResource(GetCrdThisGroupVersionResource()).Informer()
	if crdList, err := c.dynamicClient.Resource(GetCrdThisGroupVersionResource()).List(context.TODO(), v1.ListOptions{}); err == nil {
		for _, crd := range crdList.Items {
			gvr := GetCrdGroupVersionResource(&crd)
			c.dynamicInformerFactory.ForResource(gvr).Informer()
		}
	}
	// 等待缓存同步完成
	c.dynamicInformerFactory.Start(wait.NeverStop)
	c.dynamicInformerFactory.WaitForCacheSync(wait.NeverStop)
}
