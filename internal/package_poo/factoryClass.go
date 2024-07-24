package package_poo

import "fmt"

/*
***************************************************************
Example Class of Email and SMS notification sender.
Using Abstract Factory as a design pattern.
This is just an skeleton code to see the pattern implementation

Methods/Ways to send the notification:
- Email
- SMS

Channels/Technologies to send the notification:
- Twilio (For SMS)
- Amazon SES (For Email)
***************************************************************
*/

// Interface responsible for sending a notification, according to a method and a sending channel.
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

// Interface that collects the method and sending channel.
type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

/*
**********************************
	SMS
**********************************
*/
// SMS STRUCT
type SMSNotification struct {
}

// SMSNotification method
func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}

// SMSNotification method
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// STRUCT for sending SMS notifications
type SMSNotificationSender struct {
}

// SMSNotificationSender method
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

// SMSNotificationSender method
func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

/*
**********************************
	EMAIL
**********************************
*/
// Email STRUCT
type EmailNotification struct {
}

// EmailNotification method
func (EmailNotification) SendNotification() {
	fmt.Println("Sending Notification via Email")
}

// EmailNotification method
func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

// STRUCT for sending Email notifications
type EmailNotificationSender struct {
}

// EmailNotificationSender method
func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

// EmailNotificationSender method
func (EmailNotificationSender) GetSenderChannel() string {
	return "Amazon SES"
}

/*
**********************************
	Functionality
**********************************
*/
// Function that evaluates the type of notification to be sent
func GetNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("null notification type")
}

// Function that sends the notification
func SendNotification(f INotificationFactory) {
	f.SendNotification()
}

// Function that chooses the way/method by which the notification will be sent
func GetMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

// Function that chooses the channel/technology through which the notification will be sent
func GetChannel(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderChannel())
}
