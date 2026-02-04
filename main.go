package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const FILENAME = "data.json"

var (
	helpStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	doneDateStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f00"))
)

type todo struct {
	Name string     `json:"name"`
	Done *time.Time `json:"done"`
}

func (t *todo) Toggle() {
	if t.Done != nil {
		t.Done = nil
	} else {
		now := time.Now()
		t.Done = &now
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
			fmt.Println("Saving data into data.json")
			if err := saveData(m.todos); err != nil {
				fmt.Printf("Error ocurred while saving the data: %v", err)
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
	s := "What have you done?\n\n" // Header

	for i := range m.todos {
		todo := &m.todos[i]
		cursor := " "
		if m.cursor == i && !m.isAdding {
			cursor = ">"
		}
		completed := " "
		final := "\n"
		if todo.Done != nil {
			completed = "x"
			final = " - " + doneDateStyle.Render(todo.Done.Format(time.RFC3339Nano)) + "\n"
		}
		s += fmt.Sprintf("%s [%s] %s%s", cursor, completed, todo.Name, final)
	}
	if m.isAdding {
		s += m.textInput.View()
	}
	return s + helpStyle.Render("\n\n   space: toggle • a: add • enter: submit • d: delete • u: undo deletion • esc: cancel • q: exit\n")
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "New todo..."
	ti.CharLimit = 156
	ti.Width = 25

	file, err := os.Open(FILENAME)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening '%s'. ERROR: %v", FILENAME, err)
		return model{
			todos:     []todo{{Name: "My first todo"}},
			cursor:    0,
			textInput: ti,
		}
	}
	defer file.Close()

	var todos []todo

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&todos)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding '%s'. ERROR: %v", FILENAME, err)
		os.Exit(1)
	}

	return model{
		todos:     todos,
		cursor:    0,
		textInput: ti,
	}
}

func saveData(todos []todo) error {
	file, err := os.Create(FILENAME)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(todos)
	if err != nil {
		return err
	}

	return nil
}
