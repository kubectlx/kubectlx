package kubecli

import (
	"context"
	"io"
	v1 "k8s.io/api/core/v1"
)

func OpenLogsFollowStream(ctx context.Context, namespace, pod, c string) (io.ReadCloser, error) {
	clientset := getClient().client
	req := clientset.CoreV1().Pods(namespace).GetLogs(pod, &v1.PodLogOptions{
		Container: c,
		Follow:    true,
	})
	return req.Stream(ctx)
}
