/*
The standard library package log provides a basic infrastructure for log management in 
GO language that can be used for logging our GO programs. The main purpose of logging is 
to get a trace of what's happening in the program, where it's happening, and when it's happening.
 Logs can be providing code tracing, profiling, and analytics. Logging(eyes and ears of a programmer) 
 is a way to find those bugs and learn more about how the program is functioning.
*/

// Program in GO language to demonstrates how to use base log package.
package loggingPrograms
import (
	"log"
	"net/smtp"		
)
func init(){
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
func simpleLog() {
// Println writes to the standard logger.
	log.Println("main started")

// Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

// Panicln is Println() followed by a call to panic()
	log.Panicln("panic message")
}
/*
After executing this code, the output would look something like this:

Output
LOG: 2017/06/25 14:49:41.989813 C:/golang/example38.go:11: init started
LOG: 2017/06/25 14:49:41.990813 C:/golang/example38.go:15: main started
LOG: 2017/06/25 14:49:41.990813 C:/golang/example38.go:18: fatal message
exit status 1
*/


func init(){
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
func RealLog() {
// Connect to the remote SMTP server.
client, err := smtp.Dial("smtp.smail.com:25")
	if err != nil {	
		log.Fatalln(err)
	}
client.Data()
}

/*
Output
TRACE: 2017/06/25 14:54:42.662011 C:/golang/example39.go:9: init started
TRACE: 2017/06/25 14:55:03.685213 C:/golang/example39.go:15: dial tcp 23.27.98.252:25: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
exit status 1
*/