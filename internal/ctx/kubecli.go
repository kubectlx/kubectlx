package ctx

import (
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"time"
)

type clientHolder struct {
	client          *kubernetes.Clientset
	informerFactory informers.SharedInformerFactory
	kubeconfig      string
}

var client *clientHolder

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

type genericHandler struct {
}

func (gh *genericHandler) OnAdd(obj interface{}, isInInitialList bool) {
}

func (gh *genericHandler) OnUpdate(oldObj, newObj interface{}) {
}

func (gh *genericHandler) OnDelete(obj interface{}) {
}

func initInformerCache(c *clientHolder) {
	c.informerFactory = informers.NewSharedInformerFactory(c.client, time.Second*600)
	c.informerFactory.Core().V1().Pods().Informer().AddEventHandler(&genericHandler{})
	c.informerFactory.Apps().V1().Deployments().Informer().AddEventHandler(&genericHandler{})
	c.informerFactory.Core().V1().Services().Informer().AddEventHandler(&genericHandler{})
	c.informerFactory.Core().V1().Namespaces().Informer().AddEventHandler(&genericHandler{})
	c.informerFactory.Apps().V1().DaemonSets().Informer().AddEventHandler(&genericHandler{})
	c.informerFactory.Apps().V1().StatefulSets().Informer().AddEventHandler(&genericHandler{})
	c.informerFactory.Batch().V1().Jobs().Informer().AddEventHandler(&genericHandler{})
	c.informerFactory.Apps().V1().ReplicaSets().Informer().AddEventHandler(&genericHandler{})
	c.informerFactory.Core().V1().Events().Informer().AddEventHandler(&genericHandler{})
	c.informerFactory.Core().V1().ConfigMaps().Informer().AddEventHandler(&genericHandler{})
	c.informerFactory.Core().V1().Secrets().Informer().AddEventHandler(&genericHandler{})
	c.informerFactory.Core().V1().Nodes().Informer().AddEventHandler(&genericHandler{})
	// 等待缓存同步完成
	c.informerFactory.Start(wait.NeverStop)
	c.informerFactory.WaitForCacheSync(wait.NeverStop)
}

func initKubeClientWithContext(kubeconfig string) error {
	if client != nil && client.kubeconfig == kubeconfig {
		return nil
	}
	if c, err := newClientSet(kubeconfig); err != nil {
		return err
	} else {
		client = &clientHolder{
			client:     c,
			kubeconfig: kubeconfig,
		}
		initInformerCache(client)
	}
	return nil
}
