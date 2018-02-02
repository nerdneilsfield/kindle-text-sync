package main

import (
	"log"

	epub "github.com/bmaupin/go-epub"
	"github.com/gin-gonic/gin"
	"gopkg.in/mail.v2"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

//make a epub file on the disk
func mkepub(title, author, input string) string {
	output := blackfriday.Run([]byte(input))
	e := epub.NewEpub(title)

	// Set the author
	e.SetAuthor(author)

	// Add a section
	section1Body := string(output)
	e.AddSection(section1Body, "Contents", "", "")

	// Write the EPUB
	err := e.Write("/tmp/" + title + ".epub")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("EPUB create successful")
	}
	return "/tmp" + title + ".epub"
}

//send file

func send(title, from_email, target_email, epubpath, address, users, passwd string, port int) {

	//make a email
	m := mail.NewMessage()
	m.SetHeader("From", from_email)
	m.SetHeader("To", target_email)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", "Nothing")
	m.Attach(epubpath)

	//send the email
	d := mail.NewDialer(address, 587, users, passwd)
	err := d.DialAndSend(m)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(title + " sent successful!")
	}

}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "Hello")
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
