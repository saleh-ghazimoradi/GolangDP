package DIP

// DIP states that the high-level modules should not depend on the low-level modules; they both should depend on the abstraction.

// This is implemented via interfaces and dependency injection mechanism in Golang

// This violates the DIP because the NotificationService depends directly on the EmailSender

type EmailSender struct{}

func (es EmailSender) SendEmail(to, subject, body string) error {
	return nil
}

type NotificationService struct {
	emailSender EmailSender
}

func (ns NotificationService) NotifyUser(userEmail, message string) error {
	return ns.emailSender.SendEmail(userEmail, "Notification", message)
}

// USING DIP
type MessageSender interface {
	Send(to, subject, body string) error
}

type EmailsSender struct{}

func (es EmailsSender) Send(to, subject, body string) error {
	return nil
}

type NotificationsService struct {
	sender MessageSender
}

func (ns NotificationsService) NotifyUser(userEmail, message string) error {
	return ns.sender.Send(userEmail, "Notification", message)
}

func NewNotificationsService(sender MessageSender) NotificationsService {
	return NotificationsService{sender: sender}
}
