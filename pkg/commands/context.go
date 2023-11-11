package commands

type DockerContext struct {
	Name               string `json:"Name"`
	Description        string `json:"Description"`
	StackOrchestrator  string `json:"StackOrchestrator"`
	DockerEndpoint     string `json:"DockerEndpoint"`
	KubernetesEndpoint string `json:"KubernetesEndpoint"`
	Current            bool   `json:"Current"`
}
