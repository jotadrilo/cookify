package add

import (
	"fmt"
	"sort"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/jotadrilo/cookify/app/api"
)

type ProductsTable struct {
	products    map[string]*api.Product
	quantities  map[string]float32
	table       table.Model
	maxRowWidth int
}

func NewProductsTable() *ProductsTable {
	m := &ProductsTable{
		products:   make(map[string]*api.Product),
		quantities: make(map[string]float32),
		table: table.New(
			table.WithHeight(10),
		),
	}

	m.ResetDataGrid()
	m.Blur()

	return m
}

func (m *ProductsTable) Focus() {
	m.table.Focus()

	st := table.Styles{}

	st.Header = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Foreground(lipgloss.Color("#FAFAFA")).
		Padding(0, 1).
		Bold(true)

	st.Selected = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Padding(0, 1)

	st.Cell = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Padding(0, 1)

	m.table.SetStyles(st)
}

func (m *ProductsTable) Blur() {
	m.table.Blur()

	st := table.Styles{}

	st.Header = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#7D56F4")).
		Foreground(lipgloss.Color("#FAFAFA")).
		Padding(0, 1).
		Bold(true)

	st.Selected = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Faint(true).
		Padding(0, 1)

	st.Cell = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Faint(true).
		Padding(0, 1)

	m.table.SetStyles(st)
}

func (m *ProductsTable) AddProduct(product *api.Product, quantity float32) {
	m.products[product.Uuid.String()] = product
	m.quantities[product.Uuid.String()] = quantity
	m.ResetDataGrid()
}

func (m *ProductsTable) GetSelectedProduct() *api.Product {
	var uuid = m.table.SelectedRow()[1]
	return m.products[uuid]
}

func (m *ProductsTable) GetProducts() map[string]*api.Product {
	return m.products
}

func (m *ProductsTable) GetQuantities() map[string]float32 {
	return m.quantities
}

func (m *ProductsTable) DeleteProduct(product *api.Product) {
	if product == nil {
		return
	}
	delete(m.products, product.Uuid.String())
	delete(m.quantities, product.Uuid.String())
	m.ResetDataGrid()
}

func (m *ProductsTable) ResetDataGrid() {
	var (
		productsByName   = make(map[string]*api.Product, len(m.products))
		quantitiesByName = make(map[string]float32, len(m.quantities))
		names            []string
		rows             []table.Row
		maxNameLength    = 10
	)

	for _, product := range m.products {
		var name = product.Slug

		if product.DisplayNameLangEsEs != nil {
			name = *product.DisplayNameLangEsEs
		}

		names = append(names, name)
		productsByName[name] = product
		quantitiesByName[name] = m.quantities[product.Uuid.String()]

		if len(name) > maxNameLength {
			maxNameLength = len(name)
		}
	}

	sort.Strings(names)

	for ix, name := range names {
		var (
			product  = productsByName[name]
			quantity = fmt.Sprintf("%.2f", quantitiesByName[name])
			number   = fmt.Sprintf("%d", ix+1)
			id       = product.Uuid.String()
		)

		rows = append(rows, table.Row{number, id, name, quantity})
	}

	var (
		maxRowWidth = 10
		columns     = []table.Column{
			{Title: "#", Width: 4},
			{Title: "ID", Width: 36},
			{Title: "Name", Width: maxNameLength},
			{Title: "Quantity", Width: 10},
		}
	)

	for _, col := range columns {
		maxRowWidth += col.Width
	}

	m.table.SetColumns(columns)
	m.table.SetRows(rows)
	m.maxRowWidth = maxRowWidth
}

func (m *ProductsTable) GetMaxRowWidth() int {
	return m.maxRowWidth
}

func (m *ProductsTable) Init() tea.Cmd {
	return nil
}

func (m *ProductsTable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m *ProductsTable) View() string {
	return m.table.View() + "\n\n" + m.table.HelpView()
}
