package decorator

import "testing"

func TestTwoNotifications(t *testing.T) {

	notif := NewNotificationEmail(NewNotificationFacebook(NewNotification()))

	sended := notif.Send()

	if sended != 2 {
		t.Errorf("Counter %d; want 2 have %d", sended, sended)
	}

}

func TestThreeNotifications(t *testing.T) {

	notif := NewNotificationEmail(NewNotificationFacebook(NewNotificationTwitter(NewNotification())))

	sended := notif.Send()

	if sended != 3 {
		t.Errorf("Counter %d; want 3 have %d", sended, sended)
	}

}
