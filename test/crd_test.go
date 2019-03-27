package main


import (
	"testing"
	"fmt"
	"k8s-crd-operater/kevin/v1"

)

func TestOperate(t *testing.T){
	fmt.Println("test")
	var get = v1.Get{}
	fmt.Println(get.Operate())

}

func TestList(t *testing.T){
	var base = v1.Base{
		//Config:"/Users/liukai/go/src/k8s-crd-website/resources/read-test-kubeconfig",
		Group:"kevincrd.k8s.io",
		Version:"v1",
		Resource:"websites",
		Namespaces:"default",
		MasterUrl:"http://localhost:8088/",
	}
	var list = v1.List{Base:base}
	for _, item := range list.Operate().Items {
		fmt.Printf("item: %s\n", item.GetName())
		spec := item.Object["spec"].(map[string]interface{})
		fmt.Printf("item: %s\n",spec["image"])
		fmt.Println(item.GetResourceVersion())
	}
}