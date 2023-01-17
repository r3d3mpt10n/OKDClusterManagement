package main

import (
	"context"
	"fmt"

	"github.com/openshift/client-go/config/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	configv1 "github.com/openshift/api/config/v1"
)

var (
   updateImage = "registry.okd.bne-shift.net:8443/origin/okd@sha256:26f4fa1c2a6db02acfd1d8cbaf100b34ca4bff5f5db56218353987a844573dc6"
   clusterUpdate = true
)

func update(updateImage string) {
        var update *configv1.Update
	update = &configv1.Update{
	     Version: "",
             Image:   updateImage,
        }
	update.Force = true

}


func main() {
	// Create a new client to the OpenShift config API
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/brendanshephard/.kube/config")
	if err != nil {
		panic(err)
	}
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	}

        // Get the ClusterVersion version object
	cv, err := clientset.ConfigV1().ClusterVersions().Get(context.TODO(), "version", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Cluster Version Spec: %v", cv.Spec)

	// Perform an update if there currently isn't one running
        if clusterUpdate == true && cv.Spec.DesiredUpdate == nil {
		fmt.Printf("Starting cluster update")
		update(updateImage)
	}

	fmt.Printf("Cluster Version Status: %v", cv.Status)

}

