package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/as27/donottrack"
)

var (
	flagPort = flag.String("port", ":1313", "change the port of the server")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Name string
			Foo  string
			DNT  bool // include a boolean field for your template
		}{
			"Username",
			"foo",
			donottrack.IsSet(r), // call the IsSet function here
		}
		t, err := template.New("webpage").Parse(htmlTemplate)
		if err != nil {
			log.Fatal("error parsing template", err)
		}
		err = t.Execute(w, data)
		if err != nil {
			log.Fatal("error executing template", err)
		}
	})
	fmt.Printf("serving at http://localhost%s\n", *flagPort)
	err := http.ListenAndServe(*flagPort, nil)
	if err != nil {
		fmt.Println(err)
	}
}

const htmlTemplate = `<html><body>
{{ if not .DNT }}
<p>Your external code for analytics or ads comes here</p>
{{ end }}
<p>my page</p>
</body></html>`
