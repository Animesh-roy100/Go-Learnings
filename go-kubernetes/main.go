package main

import (
	"context"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	r := gin.Default()

	// Setup Kubernetes client
	kubeconfig := os.Getenv("KUBECONFIG")
	if kubeconfig == "" {
		kubeconfig = filepath.Join(os.Getenv("HOME"), ".kube", "config")
	}
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Endpoint to create a BusyBox pod
	r.POST("/createpod", func(c *gin.Context) {
		podName := c.PostForm("podName")
		namespace := c.PostForm("namespace")

		createBusyBoxPod(clientset, podName, namespace)
		c.JSON(http.StatusOK, gin.H{"message": "Pod created successfully"})
	})

	// Endpoint to delete a pod
	r.POST("/deletepod", func(c *gin.Context) {
		podName := c.PostForm("podName")
		namespace := c.PostForm("namespace")

		deletePod(clientset, podName, namespace)
		c.JSON(http.StatusOK, gin.H{"message": "Pod deleted successfully"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func createBusyBoxPod(clientset *kubernetes.Clientset, podName, namespace string) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      podName,
			Namespace: namespace,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:    "busybox",
					Image:   "busybox",
					Command: []string{"sleep", "3600"},
				},
			},
		},
	}

	_, err := clientset.CoreV1().Pods(namespace).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
}

func deletePod(clientset *kubernetes.Clientset, podName, namespace string) {
	err := clientset.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	}
}
