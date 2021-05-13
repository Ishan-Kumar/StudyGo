package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("https://www.google.com/")
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	//readUsingReaderInterface(resp)
	readUsingWriteInterface(resp)

}

// Printing simple rsp, doesn't print the entire response code
// fmt.Println(resp)
// to print the entire response code make a empty byte slice of any size say 9999
func readUsingReaderInterface(resp *http.Response) {
	bs := make([]byte, 99999)
	resp.Body.Read(bs)      // Call Read function from 'body' property in response struct and and put in byte slice
	fmt.Println(string(bs)) // print the byteslice in string fmt
}

// below one line will do the exact same thing as in above function
// @ check io.copy arg types, it takes 2 arguments : something that implements write, something that implements read
// @ os.stdout somewhere inside contains writer function
// 	Which will take the resp.Body and write out using it's interface chain
func readUsingWriteInterface(resp *http.Response) {
	io.Copy(os.Stdout, resp.Body)
}
