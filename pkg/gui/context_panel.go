package gui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
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
			GetItemContextCacheKey: func(ctx *commands.DockerContext) string {
				return "context:" + ctx.Name
			},
		},

		ListPanel: panels.ListPanel[*commands.DockerContext]{
			List: panels.NewFilteredList[*commands.DockerContext](),
			View: gui.Views.DockerContext,
		},
		NoItemsMessage: "",
		Gui:            gui.intoInterface(),

		Sort: func(a *commands.DockerContext, b *commands.DockerContext) bool {
			return a.Current
		},
		GetTableCells: presentation.GetContextDisplayStrings,
		DisableFilter: true,
		OnClick: func(ctx *commands.DockerContext) error {
			return gui.createConfirmationPanel(gui.Tr.Confirm, "确定要激活 Context: "+ctx.Name+"吗?", func(g *gocui.Gui, v *gocui.View) error {
				return gui.WithWaitingStatus(gui.Tr.StartingStatus, func() error {
					return gui.ChangeDockerContext(ctx.Name)
				})
			}, nil)
		},
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
		return nil
	}

	lst, err := gui.DockerCommand.GetAllContexts()
	if err != nil {
		return err
	}

	gui.Panels.Contexts.SetItems(lst)
	return gui.Panels.Contexts.RerenderList()
}

func (gui *Gui) handleContextUse(g *gocui.Gui, v *gocui.View) error {
	ctx, err := gui.Panels.Contexts.GetSelectedItem()
	if err != nil {
		return err
	}
	return gui.ChangeDockerContext(ctx.Name)
}
