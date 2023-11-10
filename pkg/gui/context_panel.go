package gui

import (
	"fmt"

	"github.com/jesseduffield/lazydocker/pkg/commands"
	"github.com/jesseduffield/lazydocker/pkg/gui/panels"
	"github.com/jesseduffield/lazydocker/pkg/gui/presentation"
	"github.com/jesseduffield/lazydocker/pkg/tasks"
)

func (gui *Gui) getContextPanel() *panels.SideListPanel[*commands.DockerContext] {
	return &panels.SideListPanel[*commands.DockerContext]{
		ContextState: &panels.ContextState[*commands.DockerContext]{
			GetMainTabs: func() []panels.MainTab[*commands.DockerContext] {
				return []panels.MainTab[*commands.DockerContext]{
					{
						Key:    "contexts",
						Title:  "Context",
						Render: gui.renderDockerContext,
					},
				}
			},
			GetItemContextCacheKey: func(project *commands.DockerContext) string {
				return "context:" + project.Name
			},
		},

		ListPanel: panels.ListPanel[*commands.DockerContext]{
			List: panels.NewFilteredList[*commands.DockerContext](),
			View: gui.Views.DockerContext,
		},
		NoItemsMessage: "",
		Gui:            gui.intoInterface(),

		Sort: func(a *commands.DockerContext, b *commands.DockerContext) bool {
			return false
		},
		GetTableCells: presentation.GetContextDisplayStrings,
		// It doesn't make sense to filter a list of only one item.
		DisableFilter: true,
	}
}

func (gui *Gui) renderDockerContext(dc *commands.DockerContext) tasks.TaskFunc {
	return gui.NewSimpleRenderStringTask(func() string {
		return fmt.Sprintf("Name: %s\nDescription: %s\nDockerEndpoint: %s\nKubernetesEndpoint: %s\nCurrent: %v\n",
			dc.Name, dc.Description, dc.DockerEndpoint, dc.KubernetesEndpoint, dc.Current)
	})
}

func (gui *Gui) refreshContexts() error {
	if gui.Views.DockerContext == nil {
		// if the containersView hasn't been instantiated yet we just return
		return nil
	}

	lst, err := gui.DockerCommand.GetAllContexts()

	if err != nil {
		return err
	}

	gui.Panels.Contexts.SetItems(lst)
	return gui.Panels.Contexts.RerenderList()
}
