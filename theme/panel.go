package theme

import (
	"gioui.org/layout"
	"gioui.org/widget"
)

type Panel interface {
	Layout(win *Window, gtx layout.Context) layout.Dimensions
	Title() string
	Closed() bool
	Detached() bool
	Attached() bool
	SetWindowed(bool)
}

type PanelButtons struct {
	close  widget.Clickable
	back   widget.Clickable
	detach widget.Clickable
	attach widget.Clickable

	windowed bool
}

func (pb *PanelButtons) SetWindowed(b bool) { pb.windowed = b }

func (pb *PanelButtons) Windowed() bool { return pb.windowed }

func (pb *PanelButtons) Closed() bool {
	return pb.close.Clicked()
}

func (pb *PanelButtons) Detached() bool {
	return pb.detach.Clicked()
}

func (pb *PanelButtons) Attached() bool {
	return pb.attach.Clicked()
}

func (pb *PanelButtons) Backed() bool {
	return pb.back.Clicked()
}

func (pb *PanelButtons) Layout(win *Window, gtx layout.Context) layout.Dimensions {
	type button struct {
		w     *widget.Clickable
		label string
	}

	var buttons []button
	if pb.windowed {
		buttons = []button{
			{&pb.attach, "Attach"},
			{&pb.close, "Close"},
		}
	} else {
		buttons = []button{
			{&pb.back, "Back"},
			{&pb.detach, "Detach"},
			{&pb.close, "Close"},
		}
	}

	children := make([]layout.FlexChild, 0, 3)
	for _, btn := range buttons {
		btn := btn
		children = append(children,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return Button(win.Theme, btn.w, btn.label).Layout(win, gtx)
			}),
			layout.Rigid(layout.Spacer{Width: 5}.Layout),
		)
	}

	return layout.Flex{}.Layout(gtx, children...)
}

// WidgetPanel turns any widget into a Panel. You are responsible for handling panel button events.
type WidgetPanel struct {
	w Widget
	PanelButtons
}

func (wp *WidgetPanel) Layout(win *Window, gtx layout.Context) layout.Dimensions {
	return layout.Flex{}.Layout(gtx, layout.Rigid(Dumb(win, wp.PanelButtons.Layout)), layout.Rigid(Dumb(win, wp.w)))
}
