package tui

import (
	"github.com/rivo/tview"
	"githubclitracker/github"
	"strconv"
)

func RenderDashboard(repoData []*github.RepoData) {
	app := tview.NewApplication()

	table := tview.NewTable().
		SetBorders(true)

	table.SetCell(0, 0, tview.NewTableCell("Repository Name"))
	table.SetCell(0, 1, tview.NewTableCell("Stargazers"))
	table.SetCell(0, 2, tview.NewTableCell("Forks"))
	table.SetCell(0, 3, tview.NewTableCell("Open Issues"))

	for i, repo := range repoData {
		row := i + 1
		table.SetCell(row, 0, tview.NewTableCell(repo.FullName))
		table.SetCell(row, 1, tview.NewTableCell(strconv.Itoa(repo.Stargazers)))
		table.SetCell(row, 2, tview.NewTableCell(strconv.Itoa(repo.Forks)))
		table.SetCell(row, 3, tview.NewTableCell(strconv.Itoa(repo.OpenIssues)))
	}

	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
