package notification

import (
	"go_project_Gin/internal/config"
	"go_project_Gin/internal/utils"
	"log"
	"net/smtp"
	"strconv"

	"go.uber.org/zap"
)

type EmailJob struct {
	Email  string
	UserId uint
	TaskId uint
	Action string
}

var emailQueue = make(chan EmailJob, 100)

func emailWorker(id int) {
	for job := range emailQueue {
		log.Printf("[Worker %d] Sending email for task %d (action: %s) to user %d", id, job.TaskId, job.Action, job.UserId)

		var subject, body string
		switch job.Action {
		case "welcome":
			subject = "Welcome to Task Manager!"
			body = "Hi! Your account has been successfully created. Start managing your tasks now!"
		case "created":
			subject = "New Task Created"
			body = "A new task has been created (Task ID: " + strconv.Itoa(int(job.TaskId)) + ")"
		case "updated":
			subject = "Task Updated"
			body = "Your task has been updated (Task ID: " + strconv.Itoa(int(job.TaskId)) + ")"
		case "deleted":
			subject = "Task Deleted"
			body = "Your task has been deleted (Task ID: " + strconv.Itoa(int(job.TaskId)) + ")"
		default:
			subject = "Task Notification"
			body = "Action: " + job.Action
		}
		err := sendEmailViaGmail(job.Email, subject, body)
		if err != nil {
			utils.Logger.Error("Email send failed", zap.Error(err), zap.String("email", job.Email))
		} else {
			utils.Logger.Info("Email sent successfully", zap.String("email", job.Email))
		}
	}
}

func init() {
	for i := 1; i <= 5; i++ {
		go emailWorker(i)
	}
	log.Println("Email worker started")
}

func SendTaskNotification(email string, userId uint, taskId uint, action string) {
	job := EmailJob{
		Email:  email,
		UserId: userId,
		TaskId: taskId,
		Action: action,
	}
	emailQueue <- job
}

func SendWelcomeEmail(email string, userID uint) {
	SendTaskNotification(email, userID, 0, "welcome")
}

func sendEmailViaGmail(to, subject, body string) error {
	from := config.App.FROM_EMAIL
	password := config.App.FROM_EMAIL_PASSWORD

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return err
	}
	return nil
}
