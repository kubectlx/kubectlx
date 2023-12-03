package kubecli

import (
	"fmt"
	"github.com/cxweilai/kubectlx/internal/command"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"strings"
)

func GetNamespaces() []*command.Param {
	var (
		namespaceNames []*command.Param
	)
	namespaces, err := getClient().informerFactory.Core().V1().Namespaces().Lister().List(labels.Everything())
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

func GetNodes(namePrefix string, limit int) []*command.Param {
	var (
		nodeNames []*command.Param
		nodes     []*corev1.Node
		err       error
	)
	nodes, err = getClient().informerFactory.Core().V1().Nodes().Lister().List(labels.Everything())
	if err != nil {
		return nodeNames
	}
	cnt := 0
	for _, node := range nodes {
		if !strings.HasPrefix(node.Name, namePrefix) {
			continue
		}
		nodeNames = append(nodeNames, &command.Param{
			Name:        node.Name,
			Description: string(node.Status.Phase),
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return nodeNames
}

func GetPods(namespace, namePrefix string, limit int) []*command.Param {
	var (
		podNames []*command.Param
		pods     []*corev1.Pod
		err      error
	)
	pods, err = getClient().informerFactory.Core().V1().Pods().Lister().Pods(namespace).List(labels.Everything())
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

func GetServices(namespace, namePrefix string, limit int) []*command.Param {
	var (
		serviceNames []*command.Param
		services     []*corev1.Service
		err          error
	)
	services, err = getClient().informerFactory.Core().V1().Services().Lister().Services(namespace).List(labels.Everything())
	if err != nil {
		return serviceNames
	}
	cnt := 0
	for _, svc := range services {
		if !strings.HasPrefix(svc.Name, namePrefix) {
			continue
		}
		serviceNames = append(serviceNames, &command.Param{
			Name:        svc.Name,
			Description: "",
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return serviceNames
}

func GetDeployments(namespace, namePrefix string, limit int) []*command.Param {
	var (
		deploymentNames []*command.Param
		deployments     []*appsv1.Deployment
		err             error
	)
	deployments, err = getClient().informerFactory.Apps().V1().Deployments().Lister().Deployments(namespace).List(labels.Everything())
	if err != nil {
		return deploymentNames
	}
	cnt := 0
	for _, deployment := range deployments {
		if !strings.HasPrefix(deployment.Name, namePrefix) {
			continue
		}
		deploymentNames = append(deploymentNames, &command.Param{
			Name:        deployment.Name,
			Description: fmt.Sprintf("%d/%d", deployment.Status.AvailableReplicas, deployment.Status.Replicas),
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return deploymentNames
}

func GetDaemonSets(namespace, namePrefix string, limit int) []*command.Param {
	var (
		daemonSetNames []*command.Param
		daemonSets     []*appsv1.DaemonSet
		err            error
	)
	daemonSets, err = getClient().informerFactory.Apps().V1().DaemonSets().Lister().DaemonSets(namespace).List(labels.Everything())
	if err != nil {
		return daemonSetNames
	}
	cnt := 0
	for _, daemonSet := range daemonSets {
		if !strings.HasPrefix(daemonSet.Name, namePrefix) {
			continue
		}
		daemonSetNames = append(daemonSetNames, &command.Param{
			Name:        daemonSet.Name,
			Description: fmt.Sprintf("%d", daemonSet.Status.NumberAvailable),
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return daemonSetNames
}

func GetReplicaSets(namespace, namePrefix string, limit int) []*command.Param {
	var (
		replicaSetNames []*command.Param
		replicaSets     []*appsv1.ReplicaSet
		err             error
	)
	replicaSets, err = getClient().informerFactory.Apps().V1().ReplicaSets().Lister().ReplicaSets(namespace).List(labels.Everything())
	if err != nil {
		return replicaSetNames
	}
	cnt := 0
	for _, replicaSet := range replicaSets {
		if !strings.HasPrefix(replicaSet.Name, namePrefix) {
			continue
		}
		replicaSetNames = append(replicaSetNames, &command.Param{
			Name:        replicaSet.Name,
			Description: fmt.Sprintf("%d/%d", replicaSet.Status.AvailableReplicas, replicaSet.Status.Replicas),
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return replicaSetNames
}

func GetJobs(namespace, namePrefix string, limit int) []*command.Param {
	var (
		jobNames []*command.Param
		jobs     []*batchv1.Job
		err      error
	)
	jobs, err = getClient().informerFactory.Batch().V1().Jobs().Lister().Jobs(namespace).List(labels.Everything())
	if err != nil {
		return jobNames
	}
	cnt := 0
	for _, job := range jobs {
		if !strings.HasPrefix(job.Name, namePrefix) {
			continue
		}
		desc := ""
		if len(job.Status.Conditions) > 0 {
			desc = string(job.Status.Conditions[0].Type)
		}
		jobNames = append(jobNames, &command.Param{
			Name:        job.Name,
			Description: desc,
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return jobNames
}

func GetCronJobs(namespace, namePrefix string, limit int) []*command.Param {
	var (
		jobNames []*command.Param
		cronJobs []*batchv1.CronJob
		err      error
	)
	cronJobs, err = getClient().informerFactory.Batch().V1().CronJobs().Lister().CronJobs(namespace).List(labels.Everything())
	if err != nil {
		return jobNames
	}
	cnt := 0
	for _, job := range cronJobs {
		if !strings.HasPrefix(job.Name, namePrefix) {
			continue
		}
		jobNames = append(jobNames, &command.Param{
			Name:        job.Name,
			Description: job.Spec.Schedule,
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return jobNames
}

func GetStatefulSets(namespace, namePrefix string, limit int) []*command.Param {
	var (
		statefulSetNames []*command.Param
		statefulSets     []*appsv1.StatefulSet
		err              error
	)
	statefulSets, err = getClient().informerFactory.Apps().V1().StatefulSets().Lister().StatefulSets(namespace).List(labels.Everything())
	if err != nil {
		return statefulSetNames
	}
	cnt := 0
	for _, statefulSet := range statefulSets {
		if !strings.HasPrefix(statefulSet.Name, namePrefix) {
			continue
		}
		statefulSetNames = append(statefulSetNames, &command.Param{
			Name:        statefulSet.Name,
			Description: fmt.Sprintf("%d/%d", statefulSet.Status.AvailableReplicas, statefulSet.Status.Replicas),
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return statefulSetNames
}

func GetConfigMaps(namespace, namePrefix string, limit int) []*command.Param {
	var (
		configMapNames []*command.Param
		configMaps     []*corev1.ConfigMap
		err            error
	)
	configMaps, err = getClient().informerFactory.Core().V1().ConfigMaps().Lister().ConfigMaps(namespace).List(labels.Everything())
	if err != nil {
		return configMapNames
	}
	cnt := 0
	for _, configMap := range configMaps {
		if !strings.HasPrefix(configMap.Name, namePrefix) {
			continue
		}
		configMapNames = append(configMapNames, &command.Param{
			Name:        configMap.Name,
			Description: fmt.Sprintf(""),
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return configMapNames
}

func GetSecrets(namespace, namePrefix string, limit int) []*command.Param {
	var (
		secretNames []*command.Param
		secrets     []*corev1.Secret
		err         error
	)
	secrets, err = getClient().informerFactory.Core().V1().Secrets().Lister().Secrets(namespace).List(labels.Everything())
	if err != nil {
		return secretNames
	}
	cnt := 0
	for _, secret := range secrets {
		if !strings.HasPrefix(secret.Name, namePrefix) {
			continue
		}
		secretNames = append(secretNames, &command.Param{
			Name:        secret.Name,
			Description: string(secret.Type),
		})
		cnt++
		if cnt == limit {
			break
		}
	}
	return secretNames
}
