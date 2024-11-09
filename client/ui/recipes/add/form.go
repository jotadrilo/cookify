package add

import (
	"fmt"
	"sort"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"

	"github.com/jotadrilo/cookify/app/api"
)

type ProductForm struct {
	form     *huh.Form
	uuid     string
	quantity string
}

func NewProductForm(products map[string]*api.Product) *ProductForm {
	m := &ProductForm{}
	m.SetProducts(products)
	return m
}

func (m *ProductForm) SetProducts(products map[string]*api.Product) tea.Cmd {
	var selOpts []huh.Option[string]

	for id, product := range products {
		var key = product.Slug

		if product.DisplayNameLangEsEs != nil {
			key = *product.DisplayNameLangEsEs
		}

		selOpts = append(selOpts, huh.NewOption(key, id))
	}

	sort.SliceStable(selOpts, func(i, j int) bool {
		return selOpts[i].Key < selOpts[j].Key
	})

	m.uuid = ""
	m.quantity = ""
	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Height(20).
				Title("Select one product").
				Options(selOpts...).
				Value(&m.uuid).
				WithTheme(huh.ThemeCatppuccin()),
			huh.NewInput().
				Key("quantity").
				TitleFunc(func() string {
					return fmt.Sprintf("Enter the %s quantity", m.uuid)
				}, &m.uuid).
				Description("Examples: 10, 10.5").
				Validate(func(s string) error {
					_, err := strconv.ParseFloat(s, 32)
					if err != nil {
						return fmt.Errorf("%q is not a valid quantity: %s", s, err.Error())
					}
					return nil
				}).
				Value(&m.quantity).
				WithTheme(huh.ThemeCatppuccin()),
		),
	).WithWidth(100)

	return m.form.Init()
}

func (m *ProductForm) GetUUID() string {
	return m.uuid
}

func (m *ProductForm) GetQuantity() float32 {
	q, _ := strconv.ParseFloat(m.form.GetString("quantity"), 32)
	return float32(q)
}

func (m *ProductForm) IsCompleted() bool {
	return m.form.State == huh.StateCompleted
}

func (m *ProductForm) Init() tea.Cmd {
	return m.form.Init()
}

func (m *ProductForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	form, cmd := m.form.Update(msg)

	if f, ok := form.(*huh.Form); ok {
		m.form = f
	}

	return form, cmd
}

func (m *ProductForm) View() string {
	return m.form.View()
}
