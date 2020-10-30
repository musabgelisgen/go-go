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

var (
	ext_name = filepath.Base(os.Args[0])
	ext_cli = extension.NewClient(os.Getenv("AWS_LAMBDA_RUNTIME_API"))
	print_pref = fmt.Sprintf("[%s]", extensionName)
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
