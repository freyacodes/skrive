package view

import (
	"fmt"

	"skrive/logic"
	"skrive/wrapper"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	returnToStart func() (tea.Model, tea.Cmd)
	doses         []logic.Dose
	err           error

	loadingIndicator spinner.Model
	doseTable        table.Model
}

func InitializeModel(returnToStart func() (tea.Model, tea.Cmd)) (tea.Model, tea.Cmd) {
	loadingIndicator := spinner.New()
	loadingIndicator.Spinner = spinner.Dot

	doseTable := createTable(make([]logic.Dose, 0))

	var doses []logic.Dose = nil
	var err error = nil

	model := model{
		returnToStart,
		doses,
		err,
		loadingIndicator,
		doseTable,
	}

	return model, model.Init()
}

func (m model) Init() tea.Cmd {
	return tea.Batch(load, m.loadingIndicator.Tick)
}

func createTable(doses []logic.Dose) table.Model {
	columns := []table.Column{
		{Title: "Time", Width: 20},
		{Title: "Amount", Width: 15},
		{Title: "Substance", Width: 25},
		{Title: "Route", Width: 25},
	}

	rows := make([]table.Row, len(doses))

	for i, dose := range doses {
		rows[i] = table.Row{
			dose.Time.Local().Format("2006-01-02 15:04:05"),
			dose.Quantity,
			dose.Substance,
			dose.Route,
		}
	}

	return table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q", "esc":
			return m.returnToStart()
		}
	case successfulLoadMsg:
		m.doses = msg.doses
		m.doseTable = createTable(msg.doses)
	case failedLoadMsg:
		m.err = msg.err
	}

	cmds := make([]tea.Cmd, 2)
	m.loadingIndicator, cmds[0] = m.loadingIndicator.Update(msg)
	m.doseTable, cmds[1] = m.doseTable.Update(msg)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var ui string

	if m.err != nil {
		ui = m.err.Error()
	} else if m.doses != nil {
		ui = m.doseTable.View()
	} else {
		ui = fmt.Sprintf("%s Loading", m.loadingIndicator.View())
	}

	return wrapper.Wrap(ui)
}
