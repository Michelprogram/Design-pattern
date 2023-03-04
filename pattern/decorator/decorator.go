package decorator

import "fmt"

type Component interface {
	Send() int
}

type Notification struct {
}

func NewNotification() *Notification {
	return &Notification{}
}

func (n *Notification) Send() int {
	return 0
}

type NotificationFacebook struct {
	Wrappe Component
}

func NewNotificationFacebook(wrappe Component) *NotificationFacebook {
	return &NotificationFacebook{
		Wrappe: wrappe,
	}
}

func (n *NotificationFacebook) Send() int {
	fmt.Println("Send notification to facebook")
	return n.Wrappe.Send() + 1
}

type NotificationTwitter struct {
	Wrappe Component
}

func NewNotificationTwitter(wrappe Component) *NotificationTwitter {
	return &NotificationTwitter{
		Wrappe: wrappe,
	}
}

func (n *NotificationTwitter) Send() int {
	fmt.Println("Send notification to twitter")
	return n.Wrappe.Send() + 1
}

type NotificationEmail struct {
	Wrappe Component
}

func NewNotificationEmail(wrappe Component) *NotificationEmail {
	return &NotificationEmail{
		Wrappe: wrappe,
	}
}

func (n *NotificationEmail) Send() int {
	fmt.Println("Send notification to email")
	return n.Wrappe.Send() + 1
}
