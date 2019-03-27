package v1

import (
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/dynamic"
	"k8s.io/apimachinery/pkg/runtime/schema"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type List struct {
	Base
	Labels string
}


func (list List) Operate() *unstructured.UnstructuredList{
	fmt.Println("list")
	config, err := clientcmd.BuildConfigFromFlags(list.MasterUrl, list.Config)
	if err != nil{
		panic(err)
	}
	//  Create a Dynamic Client to interface with CRDs.
	dynamicClient, _ := dynamic.NewForConfig(config)

	//  Create a GVR which represents an Istio Virtual Service.
	virtualServiceGVR := schema.GroupVersionResource{
		Group:    list.Group,
		Version:  list.Version,
		Resource: list.Resource,
	}

	result, err := dynamicClient.Resource(virtualServiceGVR).Namespace(list.Namespaces).List(initOptions(list))
	if err!=nil{
		panic(err)
	}
	return result
}

func initOptions(list List) metav1.ListOptions{
	var options = metav1.ListOptions{};
	if list.Labels != ""{
		options.LabelSelector = list.Labels
	}
	return options
}