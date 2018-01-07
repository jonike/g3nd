package gui

import (
	"fmt"

	"github.com/g3n/engine/gui"
	"github.com/g3n/g3nd/app"
	"github.com/g3n/g3nd/demos"
)

func init() {
	demos.Map["gui.tree"] = &GuiTree{}
}

type GuiTree struct {
}

func (t *GuiTree) Initialize(a *app.App) {

	labelId := 1
	nodeId := 1

	// Tree
	tree := gui.NewTree(200, 300)
	tree.SetPosition(10, 40)
	a.GuiPanel().Add(tree)
	tree.Subscribe(gui.OnChange, func(evname string, ev interface{}) {
		selected := tree.Selected()
		a.Log().Info("OnChange: selected:%T", selected)
	})

	addChild := func(child gui.IPanel) {
		// Get selected item.
		// If no item selected add to the tree
		sel := tree.Selected()
		if sel == nil {
			tree.Add(child)
			return
		}
		// If selected item is a node adds to this node
		node, ok := sel.(*gui.TreeNode)
		if ok {
			node.Add(child)
			return
		}
		// Add child to the parent node of the selected item
		par, pos := tree.FindChild(sel)
		if par != nil {
			par.InsertAt(pos, child)
		} else {
			tree.InsertAt(pos, child)
		}
	}

	// Add label button
	b1 := gui.NewButton("Add label")
	b1.SetPosition(10, 10)
	b1.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		item := gui.NewImageLabel(fmt.Sprintf("label %d", labelId))
		labelId++
		addChild(item)
	})
	a.GuiPanel().Add(b1)

	// Add checkbox button
	b2 := gui.NewButton("Add checkbox")
	b2.SetPosition(b1.Position().X+b1.Width()+10, 10)
	b2.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		item := gui.NewCheckBox(fmt.Sprintf("check %d", labelId))
		labelId++
		addChild(item)
	})
	a.GuiPanel().Add(b2)

	// Add node button
	b3 := gui.NewButton("Add node")
	b3.SetPosition(b2.Position().X+b2.Width()+10, 10)
	b3.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		// Get selected item.
		// If no item selected add to the tree
		sel := tree.Selected()
		if sel == nil {
			tree.AddNode(fmt.Sprintf("node %d", nodeId))
			nodeId++
			return
		}
		// If selected item is a node adds to this node
		node, ok := sel.(*gui.TreeNode)
		if ok {
			node.AddNode(fmt.Sprintf("node %d", nodeId))
			nodeId++
			return
		}
		// Add node to the parent node of the selected item
		par, pos := tree.FindChild(sel)
		if par != nil {
			par.InsertNodeAt(pos, fmt.Sprintf("node %d", nodeId))
			nodeId++
		} else {
			tree.InsertNodeAt(pos, fmt.Sprintf("node %d", nodeId))
			nodeId++
		}
	})
	a.GuiPanel().Add(b3)

	// Add remove button
	b4 := gui.NewButton("Remove")
	b4.SetPosition(b3.Position().X+b3.Width()+10, 10)
	b4.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		sel := tree.Selected()
		if sel == nil {
			return
		}
		tree.Remove(sel)
		sel.Dispose()
	})
	a.GuiPanel().Add(b4)

	// Add clear button
	b5 := gui.NewButton("Clear")
	b5.SetPosition(b4.Position().X+b4.Width()+10, 10)
	b5.Subscribe(gui.OnClick, func(evname string, ev interface{}) {
		tree.Clear()
	})
	a.GuiPanel().Add(b5)
}

func (t *GuiTree) Render(a *app.App) {
}
