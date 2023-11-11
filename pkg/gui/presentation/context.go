package presentation

import "github.com/jesseduffield/lazydocker/pkg/commands"

func GetContextDisplayStrings(c *commands.DockerContext) []string {
	if c.Current {
		return []string{"* " + c.Name}
	}
	return []string{c.Name}
}
