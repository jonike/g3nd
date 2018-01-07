package gui

import (
	"github.com/g3n/engine/gui"
	"github.com/g3n/engine/math32"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["gui.checkradio"] = &CheckRadio{}
}

type CheckRadio struct{}

func (t *CheckRadio) Initialize(app *app.App) {

	// Checkbox 1
	cb1 := gui.NewCheckBox("Checkbox 1")
	cb1.SetPosition(10, 10)
	cb1.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("Checkbox1 OnChange. State:%v", cb1.Value())
	})
	app.GuiPanel().Add(cb1)

	// Checkbox 2
	cb2 := gui.NewCheckBox("Checkbox 2")
	cb2.Label.SetFontSize(16)
	cb2.SetPosition(10, 40)
	cb2.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("Checkbox1 OnChange. State:%v", cb2.Value())
	})
	app.GuiPanel().Add(cb2)

	// Checkbox 3
	cb3 := gui.NewCheckBox("Checkbox 3 (style changed)")
	cb3.SetPosition(10, 70)
	cb3.Label.SetFontSize(16)
	// Copy and change the default style
	styles := gui.StyleDefault().CheckRadio
	styles.Over.BgColor = math32.Color4Name("red", 1)
	styles.Over.FgColor = math32.ColorName("blue")
	cb3.SetStyles(&styles)
	cb3.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("Checkbox3 OnChange. State:%v", cb3.Value())
	})
	app.GuiPanel().Add(cb3)

	// Checkbox 4
	cb4Enabled := "Checkbox 4 (enabled)"
	cb4Disabled := "Checkbox 4 (disabled)"
	cb4 := gui.NewCheckBox(cb4Disabled)
	cb4.SetPosition(10, 100)
	cb4.Label.SetFontSize(16)
	cb4.SetEnabled(false)
	cb4.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("Checkbox4 OnChange. State:%v", cb4.Value())
	})
	app.GuiPanel().Add(cb4)

	// Checkbox 5
	cb5 := gui.NewCheckBox("Checkbox 4 (control)")
	cb5.SetPosition(200, 100)
	cb5.Label.SetFontSize(16)
	cb5.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		cb4.SetEnabled(cb5.Value())
		if cb4.Enabled() {
			cb4.Label.SetText(cb4Enabled)
		} else {
			cb4.Label.SetText(cb4Disabled)
		}
	})
	app.GuiPanel().Add(cb5)

	// RadioButton 1
	rb1 := gui.NewRadioButton("RadioButton 1")
	rb1.SetPosition(10, 200)
	rb1.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("RadioButton 1 OnChange. State:%v", rb1.Value())
	})
	app.GuiPanel().Add(rb1)

	// RadioButton 2
	rb2 := gui.NewRadioButton("RadioButton 2")
	rb2.SetPosition(10, 230)
	rb2.Label.SetFontSize(16)
	rb2.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("RadioButton 2 OnChange. State:%v", rb2.Value())
	})
	app.GuiPanel().Add(rb2)

	// RadioButton 3 (radio group)
	rb3 := gui.NewRadioButton("RadioButton 3 (group)")
	rb3.SetPosition(10, 270)
	rb3.Label.SetFontSize(16)
	rb3.SetGroup("group1")
	rb3.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("RadioButton 3 OnChange. State:%v", rb3.Value())
	})
	app.GuiPanel().Add(rb3)
	// RadioButton 4 (radio group)
	rb4 := gui.NewRadioButton("RadioButton 4 (group)")
	rb4.SetPosition(10, 295)
	rb4.Label.SetFontSize(16)
	rb4.SetGroup("group1")
	rb4.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("RadioButton 4 OnChange. State:%v", rb4.Value())
	})
	app.GuiPanel().Add(rb4)
	// RadioButton 4 (radio group)
	rb5 := gui.NewRadioButton("RadioButton 5 (group)")
	rb5.SetPosition(10, 320)
	rb5.Label.SetFontSize(16)
	rb5.SetGroup("group1")
	rb5.Subscribe(gui.OnChange, func(name string, ev interface{}) {
		app.Log().Info("RadioButton 5 OnChange. State:%v", rb5.Value())
	})
	app.GuiPanel().Add(rb5)
}

func (t *CheckRadio) Render(app *app.App) {
}
