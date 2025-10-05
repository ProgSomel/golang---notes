package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type MyResource struct {
	name string
}

func (m MyResource) Close() error {
	fmt.Println("Closing Resourse: ", m.name)
	return nil
}

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader: ", err)
	}
	fmt.Println(string(buf[:n])) //? 0 to n-1
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing in writer: ", err)
	}

}

func closeResource(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatalln("Error closing resource: ", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer
	buf.WriteString("Hello Buffer!")
	fmt.Println(buf.String())
}

func multiReaderExample() {
	r1 := strings.NewReader("Hello ")
	r2 := strings.NewReader("World!")
	mr := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) 
	_, err := buf.ReadFrom(mr)
	if err != nil {
		log.Fatalln("Error reading from reader: ", err)
	}
	fmt.Println(buf.String())
}

func writeToFile(filePath string, data string) {
	file, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening/creating file: ", err)
	}
	defer closeResource(file)
	
	_, err = file.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to file: ", err)
	}

	//? alternative way to write the above lines
	// Type(value)
	// writer := io.Writer(file)
	// _, err = writer.Write([]byte(data))
	// if err != nil {
	// 	log.Fatalln("Error to writing to file: ", err)
	// }

}

func pipeExample() {
	pr, pw := io.Pipe()
	//? If we add the go keyword before the function, it becomes a go routine
	//? Go routines are functions that are immediately executed and they are anonymous
	/* this go keyword extracts this function out of the main thread.
	So it this function is going to take let's say 30 minutes, the next statement will not wait for 
	30 minutes to get executed.
	this functio will be extracted away from the main thread and the execution will fall on the next line
	The subsequent line will be executed. And this function, once it is completed, then it will come back to
	the main thread.
	Like async, await in javascript
	*/
	go func() {
		pw.Write([]byte("Hello Pipe"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println(buf.String())
}

func main() {
	fmt.Println("=== Read from the Reader ===")
	readFromReader(strings.NewReader("Hello Reader!"))

	fmt.Println("=== Write to writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Hello Writer")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== Multi Reader Example ===")
	multiReaderExample()

	fmt.Println("=== Pipe Example ===")
	pipeExample()

	filePath := "io.txt"
	writeToFile(filePath, "Hello File!")

	resource := &MyResource{name: "Test Resource"}
	closeResource(resource)
}