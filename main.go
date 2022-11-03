package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"

	"github.com/joho/godotenv"
)

var httpPassword *string
var mailer *Mailer

func processMail(
	senderName *string,
	destination *string,
	subject *string,
	text string,
	templateName *string,
) error {
	var buffer bytes.Buffer

	if templateName != nil {
		filePath := path.Join("templates", fmt.Sprintf("%s.html", *templateName))
		tmpl, err := template.ParseFiles(filePath)
		if err != nil {
			return err
		}

		values := strings.Split(text, ":")
		var data map[string]string = map[string]string{}
		for i := 0; i < len(values); i++ {
			data["v"+fmt.Sprint(i)] = values[i]
		}

		tmpl.Execute(&buffer, data)
	} else {
		buffer = *bytes.NewBuffer([]byte(text))
	}

	if err := mailer.Send(
		senderName,
		destination,
		subject,
		buffer.String(),
	); err != nil {
		return err
	}

	return nil
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Wrong method", http.StatusBadRequest)
		return
	}

	if httpPassword != nil {
		password := r.URL.Query().Get("password")

		if password != *httpPassword {
			http.Error(w, "Wrong password", http.StatusUnauthorized)
			return
		}
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Print(err.Error())
		http.Error(w, "Error request", http.StatusBadRequest)
	}

	payloads := strings.Split(string(bytes), "|")

	var senderName, destination, subject *string
	var text string

	if len(payloads[0]) > 0 {
		senderName = &payloads[0]
	}

	if len(payloads[1]) > 0 {
		destination = &payloads[1]
	}

	if len(payloads[2]) > 0 {
		subject = &payloads[2]
	}

	if len(payloads[3]) > 0 {
		text = payloads[3]
	}

	var templateName *string
	v := r.URL.Query().Get("template")
	if len(v) > 0 {
		templateName = &v
	}

	if err := processMail(
		senderName,
		destination,
		subject,
		text,
		templateName,
	); err != nil {
		log.Fatalf("senderName: %+v | destination: %+v | subject: %+v | text: %s | error: %s", senderName, destination, subject, text, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func runServer(
	host string,
	port string,
	password *string,
) {
	http.HandleFunc("/", handlerIndex)

	address := fmt.Sprintf("%s:%s", host, port)

	log.Printf("Application (HTTP server) running on %s", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalln("Failed to run HTTP server")
		log.Fatalln(err.Error())
		return
	}

}

func buildMailer() error {
	var smtpHost string
	var smtpPort uint16
	var smtpAuthUser string
	var smtpAuthPass *string
	var smtpSenderUser string
	var tmp string

	smtpHost = os.Getenv("APP_SMTP_HOST")
	if len(smtpHost) == 0 {
		smtpHost = "127.0.0.1"
	}

	tmp = os.Getenv("APP_SMTP_PORT")
	if len(tmp) == 0 {
		smtpPort = uint16(465)
	} else {
		value, err := strconv.ParseUint(tmp, 10, 16)

		if err != nil {
			return errors.New("invalid parse value on APP_SMTP_PORT")
		} else {
			smtpPort = uint16(value)
		}
	}

	smtpAuthUser = os.Getenv("APP_SMTP_AUTH_USER")
	if len(smtpAuthUser) == 0 {
		return errors.New("APP_SMTP_AUTH_USER is required")
	}

	smtpSenderUser = os.Getenv("APP_SMTP_SENDER_USER")
	if len(smtpSenderUser) == 0 {
		return errors.New("APP_SMTP_SENDER_USER is required")
	}

	tmp = os.Getenv("APP_SMTP_AUTH_PASS")
	if len(smtpHost) > 0 {
		smtpAuthPass = &tmp
	}

	mailer = NewMailer(
		smtpHost,
		smtpPort,
		smtpAuthUser,
		smtpAuthPass,
		smtpSenderUser,
	)

	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err.Error())
		log.Fatal("Error loading .env file. Please configure .env file correctly and re-launch this application.")
		return
	}

	var host, port string

	host = os.Getenv("APP_HTTP_HOST")
	if len(host) == 0 {
		host = "127.0.0.1"
	}

	port = os.Getenv("APP_HTTP_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	password := os.Getenv("APP_HTTP_PASSWORD")
	if len(password) > 0 {
		httpPassword = &password
	}

	if err := buildMailer(); err != nil {
		log.Fatal(err)
		return
	}

	runServer(host, port, httpPassword)
}
