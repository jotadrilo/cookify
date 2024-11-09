package add

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"

	"github.com/jotadrilo/cookify/app/api"
)

type Model struct {
	products           map[string]*api.Product
	unselectedProducts map[string]*api.Product

	nameForm      *huh.Form
	productsForm  *ProductForm
	productsTable *ProductsTable
	name          string

	lg     *lipgloss.Renderer
	styles *Styles
}

type modelOptions struct {
	products []api.Product
}

type ModelOption func(*modelOptions)

func WithModelProducts(products ...api.Product) ModelOption {
	return func(o *modelOptions) {
		o.products = products
	}
}

func makeOptions(s ...ModelOption) *modelOptions {
	var opts = &modelOptions{}
	for _, o := range s {
		o(opts)
	}
	return opts
}

func NewModel(s ...ModelOption) *Model {
	opts := makeOptions(s...)

	m := Model{
		lg:                 lipgloss.DefaultRenderer(),
		products:           make(map[string]*api.Product, len(opts.products)),
		unselectedProducts: make(map[string]*api.Product),
	}

	for _, product := range opts.products {
		m.products[product.Uuid.String()] = &product
		m.unselectedProducts[product.Uuid.String()] = &product
	}

	m.styles = NewStyles(m.lg)
	m.nameForm = huh.NewForm(huh.NewGroup(
		huh.NewInput().
			PlaceholderFunc(func() string {
				return m.name
			}, &m.name).
			Value(&m.name).
			WithTheme(huh.ThemeCatppuccin()),
	))
	m.productsForm = NewProductForm(m.unselectedProducts)
	m.productsTable = NewProductsTable()

	return &m
}

func (m *Model) AddProduct() tea.Cmd {
	cur := m.productsForm

	if cur != nil {
		m.productsTable.AddProduct(m.getProductByUUID(cur.GetUUID()), cur.GetQuantity())
		delete(m.unselectedProducts, cur.GetUUID())
	}

	return m.productsForm.SetProducts(m.unselectedProducts)
}

func (m *Model) DeleteProduct(product *api.Product) tea.Cmd {
	m.productsTable.DeleteProduct(product)
	m.unselectedProducts[product.Uuid.String()] = product
	return m.productsForm.SetProducts(m.unselectedProducts)
}

func (m *Model) getProductByUUID(uuid string) *api.Product {
	return m.products[uuid]
}

func (m *Model) GetSelectedProducts() map[string]*api.Product {
	return m.productsTable.GetProducts()
}

func (m *Model) GetSelectedQuantities() map[string]float32 {
	return m.productsTable.GetQuantities()
}

func (m *Model) GetName() string {
	return m.name
}

func (m *Model) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds, m.nameForm.Init())
	cmds = append(cmds, m.productsForm.Init())
	cmds = append(cmds, m.productsTable.Init())

	return tea.Batch(cmds...)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyTab:
			if m.productsTable.table.Focused() {
				m.productsTable.table.Blur()
			} else {
				m.productsTable.table.Focus()
			}
		}

		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	}

	// Process the name form
	if m.nameForm.State == huh.StateNormal {
		_, cmd := m.nameForm.Update(msg)
		cmds = append(cmds, cmd)
	} else if m.productsTable.table.Focused() {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "-":
				var product = m.productsTable.GetSelectedProduct()

				cmds = append(cmds, m.DeleteProduct(product))

				m.productsTable.table.Blur()

				_, cmd := m.productsForm.Update(nil)
				cmds = append(cmds, cmd)

				return m, tea.Batch(cmds...)
			}
		}

		// Process the table
		_, cmd := m.productsTable.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		var cmd tea.Cmd

		// Process the form
		_, cmd = m.productsForm.Update(msg)
		cmds = append(cmds, cmd)

		if m.productsForm.IsCompleted() {
			cmds = append(cmds, m.AddProduct())
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	var (
		header   = m.styles.Header.Width(200).Render("Add recipe")
		nameForm string
		form     string
		table    string
	)

	if t := m.nameForm; t != nil {
		st := m.styles.GroupHeaderFocus
		if m.productsTable.table.Focused() {
			st = m.styles.GroupHeader
		}

		nameForm = st.Width(100).Render("Enter name")
		nameForm += "\n\n"
		if m.nameForm.State == huh.StateNormal {
			nameForm += m.styles.Group.Render(m.nameForm.View())
		} else {
			nameForm += m.styles.Group.MarginLeft(2).Render(fmt.Sprintf("Adding %s recipe", m.name))
		}
	}

	if t := m.productsTable; t != nil {
		st := m.styles.GroupHeader
		if m.productsTable.table.Focused() {
			st = m.styles.GroupHeaderFocus
		}

		table = st.Width(100).Render("Current ingredients")
		table += "\n\n"
		table += m.styles.Group.Render(m.productsTable.View())
	}

	if t := m.productsForm; t != nil {
		st := m.styles.GroupHeaderFocus
		if m.productsTable.table.Focused() {
			st = m.styles.GroupHeader
		}

		form = st.Width(100).Render("Add ingredient")
		form += "\n\n"
		form += m.styles.Group.Render(t.View())
	}

	body := nameForm + "\n\n" + lipgloss.JoinHorizontal(lipgloss.Top, form, table)

	return m.styles.Base.Render("\n" + header + "\n\n" + body + "\n\n")
}
