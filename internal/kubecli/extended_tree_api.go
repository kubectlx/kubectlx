package kubecli

import (
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
)

type ResourceTree struct {
	Type   string
	Name   string
	Status string
	Child  []*ResourceTree
}

func GetDeploymentTree(namespace, name string) (*ResourceTree, error) {
	factory := getClient().informerFactory
	deployment, err := factory.Apps().V1().Deployments().Lister().Deployments(namespace).Get(name)
	if err != nil {
		return nil, err
	}
	tree := &ResourceTree{
		Type:   "deployments",
		Name:   deployment.Name,
		Status: fmt.Sprintf("%d/%d", deployment.Status.AvailableReplicas, deployment.Status.Replicas),
		Child:  make([]*ResourceTree, 0),
	}
	ss := labels.Set{}
	for k, v := range deployment.Spec.Selector.MatchLabels {
		ss[k] = v
	}
	labelSelected := labels.SelectorFromSet(ss)
	pods, err := factory.Core().V1().Pods().Lister().Pods(namespace).List(labelSelected)
	if err != nil {
		return nil, err
	}
	for _, pod := range pods {
		tree.Child = append(tree.Child, &ResourceTree{
			Type:   "pods",
			Name:   pod.Name,
			Status: string(pod.Status.Phase),
		})
	}
	return tree, nil
}

func GetServiceTree(namespace, name string) (*ResourceTree, error) {
	factory := getClient().informerFactory
	service, err := factory.Core().V1().Services().Lister().Services(namespace).Get(name)
	if err != nil {
		return nil, err
	}
	tree := &ResourceTree{
		Type:   "services",
		Name:   service.Name,
		Status: string(service.Spec.Type),
		Child:  make([]*ResourceTree, 0),
	}
	ss := labels.Set{}
	for k, v := range service.Spec.Selector {
		ss[k] = v
	}
	labelSelected := labels.SelectorFromSet(ss)
	pods, err := factory.Core().V1().Pods().Lister().Pods(namespace).List(labelSelected)
	if err != nil {
		return nil, err
	}
	for _, pod := range pods {
		tree.Child = append(tree.Child, &ResourceTree{
			Type:   "pods",
			Name:   pod.Name,
			Status: string(pod.Status.Phase),
		})
	}
	return tree, nil
}

func GetStatefulSetTree(namespace, name string) (*ResourceTree, error) {
	factory := getClient().informerFactory
	statefulSet, err := factory.Apps().V1().StatefulSets().Lister().StatefulSets(namespace).Get(name)
	if err != nil {
		return nil, err
	}
	tree := &ResourceTree{
		Type:   "statefulsets",
		Name:   statefulSet.Name,
		Status: fmt.Sprintf("%d/%d", statefulSet.Status.AvailableReplicas, statefulSet.Status.Replicas),
		Child:  make([]*ResourceTree, 0),
	}
	ss := labels.Set{}
	for k, v := range statefulSet.Spec.Selector.MatchLabels {
		ss[k] = v
	}
	labelSelected := labels.SelectorFromSet(ss)
	pods, err := factory.Core().V1().Pods().Lister().Pods(namespace).List(labelSelected)
	if err != nil {
		return nil, err
	}
	for _, pod := range pods {
		tree.Child = append(tree.Child, &ResourceTree{
			Type:   "pods",
			Name:   pod.Name,
			Status: string(pod.Status.Phase),
		})
	}
	return tree, nil
}

func GetDaemonSetTree(namespace, name string) (*ResourceTree, error) {
	factory := getClient().informerFactory
	daemonSet, err := factory.Apps().V1().DaemonSets().Lister().DaemonSets(namespace).Get(name)
	if err != nil {
		return nil, err
	}
	tree := &ResourceTree{
		Type:   "daemonsets",
		Name:   daemonSet.Name,
		Status: fmt.Sprintf("Available:%d", daemonSet.Status.NumberAvailable),
		Child:  make([]*ResourceTree, 0),
	}
	ss := labels.Set{}
	for k, v := range daemonSet.Spec.Selector.MatchLabels {
		ss[k] = v
	}
	labelSelected := labels.SelectorFromSet(ss)
	pods, err := factory.Core().V1().Pods().Lister().Pods(namespace).List(labelSelected)
	if err != nil {
		return nil, err
	}
	for _, pod := range pods {
		tree.Child = append(tree.Child, &ResourceTree{
			Type:   "pods",
			Name:   pod.Name,
			Status: string(pod.Status.Phase),
		})
	}
	return tree, nil
}

func GetJobTree(namespace, name string) (*ResourceTree, error) {
	factory := getClient().informerFactory
	job, err := factory.Batch().V1().Jobs().Lister().Jobs(namespace).Get(name)
	if err != nil {
		return nil, err
	}
	status := ""
	if len(job.Status.Conditions) > 0 {
		status = string(job.Status.Conditions[0].Type)
	}
	tree := &ResourceTree{
		Type:   "jobs",
		Name:   job.Name,
		Status: status,
		Child:  make([]*ResourceTree, 0),
	}
	ss := labels.Set{}
	for k, v := range job.Spec.Selector.MatchLabels {
		ss[k] = v
	}
	labelSelected := labels.SelectorFromSet(ss)
	pods, err := factory.Core().V1().Pods().Lister().Pods(namespace).List(labelSelected)
	if err != nil {
		return nil, err
	}
	for _, pod := range pods {
		tree.Child = append(tree.Child, &ResourceTree{
			Type:   "pods",
			Name:   pod.Name,
			Status: string(pod.Status.Phase),
		})
	}
	return tree, nil
}
