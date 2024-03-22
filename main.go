package main

import (
	"log"
	"net/http"
	"os"

	spelling "github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/cmd"
)

var rootCMD = &cobra.Command{runE: func(cms *cobra.Command, args []string) error { return cmd.Help() }}

func init() {
	rootCMD.AddCommand(spelling.CMD)
}
func main() {
	/* 	log.Println("Hi")
	   	d, err := io.ReadAll(r.Body)
	   	if err != nil {
	   		http.Error(rw, "Ooops", http.StatusBadRequest)
	   		return
	   	}
	   	fmt.FprintF(rw, "Hello %s\n", d)
	*/
	l := log.New(os.Stdout, "spelling-api", log.LstdFlags)
	sh := srv.spelling
	sm := http.NewServeMux()
	sm.Handle("/spelling", sh)
	s := &http.Server{}
	http.ListenAndServe(":9090", sm)

}
