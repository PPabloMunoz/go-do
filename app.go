package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "New todo..."
	ti.CharLimit = 156
	ti.Width = 25

	todos := loadData()

	return model{
		todos:     todos,
		cursor:    0,
		textInput: ti,
	}
}

type model struct {
	todos     []todo
	cursor    int
	textInput textinput.Model
	isAdding  bool
	undos     []todo
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg: // Key pressed?
		switch msg.String() {
		case "ctrl+c", "q":
			if err := saveData(m.todos); err != nil {
				fmt.Println(errorStyle.Render("Error: " + err.Error()))
			}
			return m, tea.Quit
		case "up", "k":
			if !m.isAdding {
				if m.cursor > 0 {
					m.cursor--
				}
			}
		case "down", "j":
			if !m.isAdding {
				if m.cursor < len(m.todos)-1 {
					m.cursor++
				}
			}
		case " ":
			if !m.isAdding && len(m.todos) > 0 {
				todo := &m.todos[m.cursor]
				todo.Toggle()
			}
		case "d":
			if !m.isAdding && len(m.todos) > 0 {
				todo := m.todos[m.cursor]
				m.undos = append(m.undos, todo)
				m.todos = append(m.todos[:m.cursor], m.todos[m.cursor+1:]...)
				if m.cursor > len(m.todos)-1 {
					m.cursor = len(m.todos) - 1
				}
			}
		case "a":
			if !m.isAdding {
				m.isAdding = true
				m.textInput.Reset()
				m.textInput.Focus()
				return m, nil
			}
		case "esc":
			m.isAdding = false
			m.textInput.Reset()
		case "enter":
			if m.isAdding {
				name := strings.TrimSpace(m.textInput.Value())
				if name != "" {
					m.todos = append(m.todos, todo{Name: name})
				}
				m.textInput.Blur()
				m.isAdding = false
			}
		case "u":
			if !m.isAdding && len(m.undos) > 0 {
				todo := m.undos[len(m.undos)-1]
				m.todos = append(m.todos, todo)
				m.cursor = len(m.todos) - 1
				m.undos = m.undos[:len(m.undos)-1]
			}
		}

	}

	if m.isAdding {
		m.textInput, cmd = m.textInput.Update(msg)
	}
	return m, cmd
}

func (m model) View() string {
	s := headerStyle.Render("What have you done?") + "\n"

	for i := range m.todos {
		todo := &m.todos[i]
		cursor := " "
		final := "\n"
		if m.cursor == i && !m.isAdding {
			cursor = cursorStyle.Render(">")
		}
		if todo.Done != nil {
			final = "  " + checkboxDoneStyle.Render("[x]") + " " + doneTextStyle.Render(todo.Name) + " " + doneDateStyle.Render("• "+todo.Done.Format("Jan 02 15:04")) + "\n"
		} else {
			final = "  " + checkboxPendingStyle.Render("[ ]") + " " + pendingTextStyle.Render(todo.Name) + "\n"
		}
		s += cursor + final
	}
	if m.isAdding {
		s += inputStyle.Render("> ") + m.textInput.View()
	}

	helpText := "\n " + keyStyle.Render("space") + ": toggle • " + keyStyle.Render("a") + ": add • " + keyStyle.Render("enter") + ": submit • " + keyStyle.Render("d") + keyStyle.Render("esc") + ": delete • " + keyStyle.Render("u") + ": undo • " + ": cancel • " + keyStyle.Render("q") + ": exit"
	return containerStyle.Render(s + helpStyle.Render(helpText))
}
