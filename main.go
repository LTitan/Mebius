package main

import (
	"context"
	"fmt"

	"github.com/LTitan/Mebius/pkg/clients/clientset/mebius"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		panic(err)
	}
	cli, err := mebius.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	obj, err := cli.MebiusApis().Clusters("default").Get(context.TODO(), "test-cluster", v1.GetOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println(obj)
}
