/*
 * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gdk"
	_ "github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

// create About DLG
func setAboutDlg() {
	about, _ := gtk.AboutDialogNew()
	about.Activate()
	about.SetComments("Go Gtk(gotk3) MENU DEMO")
	about.SetVersion("0.1")
	about.SetName("simplemenudemo")
	about.SetCopyright("License MIT")
	about.AddCreditSection("", []string{"nodryo"})
	about.SetModal(true)
	about.ShowNow()
	about.Run()
	about.Destroy()
}
func reallyquit(win *gtk.Window) {
	dlg := gtk.MessageDialogNew(win, gtk.DIALOG_MODAL, gtk.MESSAGE_QUESTION, gtk.BUTTONS_YES_NO, "really quit")
	resp := dlg.Run()
	////  fmt.Printf("resp is  %v \n wait for %v", resp, gtk.RESPONSE_YES)
	if resp == gtk.RESPONSE_YES {
		dlg.Destroy()
		gtk.MainQuit()
	}
	dlg.Destroy()
}
func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new label widget to show in the window.
	l, err := gtk.LabelNew("Hello, gotk3!")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	// create main VerticalBox
	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)

	// create main menu
	menubar, _ := gtk.MenuBarNew()
	// add main File menu
	menufile, _ := gtk.MenuItemNewWithMnemonic("File")
	menubar.Add(menufile)
	// add file sub menu
	submenufile, _ := gtk.MenuNew()
	menufile.SetSubmenu(submenufile)
	menuopen, _ := gtk.MenuItemNewWithMnemonic("Open")
	menuclose, _ := gtk.MenuItemNewWithMnemonic("Close")
	menusep, _ := gtk.SeparatorMenuItemNew()
	menuquit, _ := gtk.MenuItemNewWithMnemonic("Quit")
	submenufile.Add(menuopen)
	submenufile.Add(menuclose)
	submenufile.Add(menusep)
	submenufile.Add(menuquit)

	// add help main menu
	menuhelp, _ := gtk.MenuItemNewWithLabel("Help")
	menubar.Add(menuhelp)
	// submenu about
	submenuhelp, _ := gtk.MenuNew()
	menuhelp.SetSubmenu(submenuhelp)
	menuabout, _ := gtk.MenuItemNewWithMnemonic("About")
	submenuhelp.Add(menuabout)

	// here create popup menu
	popupmenu, _ := gtk.MenuNew()
	pabout, _ := gtk.MenuItemNew()
	pabout.SetLabel("About")
	pquit, _ := gtk.MenuItemNew()
	pquit.SetLabel("Quit")
	popupmenu.Append(pabout)
	popupmenu.Append(pquit)

	// call open dlg
	menuopen.Connect("activate", func() {
		filedlg := gtk.OpenFileChooserNative("open file", win)
		fmt.Printf("file chosen %s \n", *filedlg)
	})

	// call menu quit
	menuquit.Connect("activate", func() {
		reallyquit(win)
	})
	pquit.Connect("activate", func() {
		reallyquit(win)
	})

	// call show ABOUT
	menuabout.Connect("activate", func() {
		setAboutDlg()
	})
	// call show ABOUT
	pabout.Connect("activate", func() {
		setAboutDlg()
	})

	//  show menu popup on button 3
	win.Connect("button-press-event", func(win *gtk.Window, ev *gdk.Event) {
		button := &gdk.EventButton{ev}
		if button.Button() == 3 {
			popupmenu.PopupAtPointer(ev)
			win.QueueDraw()
		}
	})

	// add menu and label to vertical contener
	vbox.PackStart(menubar, false, false, 0)
	vbox.Add(l)
	// add main contener in window
	win.Add(vbox)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	menubar.ShowAll()
	popupmenu.ShowAll()
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
