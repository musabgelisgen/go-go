package main
import (
   "fmt"
   "context"
	"encoding/json"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)


func main() {
   fmt.Println("hello world")
}

func prettyPrint(v interface{}) string {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return ""
	}
	return string(data)
}
