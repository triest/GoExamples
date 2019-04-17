package main
import (
	"fmt"
	"io"
	"net/http"
	"os"
)
func main() {
	client := http.Client{}
	resp, err := client.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}