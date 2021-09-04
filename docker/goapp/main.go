package main
import ( // 라이브러리 목록
	"fmt"
	"os"
	"log"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request){ // request와 response에 대한 함수
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w,"hostname:", name)
}
func main() {
  fmt.Fprintln(os.Stdout,"Starting GoApp Server......")
	http.HandleFunc("/",handler) // http로 서비스
	log.Fatal(http.ListenAndServe(":8080",nil))
}