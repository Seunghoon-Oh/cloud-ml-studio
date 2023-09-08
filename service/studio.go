package service

import "github.com/Seunghoon-Oh/cloud-ml-studio-manager/data"


func GetStudios() []string {
	return data.GetRedisData()
}

func CreateStudio() string {
	// TODO: Create Kubernetes Object
	return data.CreateRedisData()
}
