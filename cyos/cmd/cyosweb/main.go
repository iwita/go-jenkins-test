package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"evolve/cyos"
)

func main() {

	port := flag.Int("port", 3000, "the port to start the application on")

	//Read the file using flags
	fileName := flag.String("file", "gopher.json", "the JSON file with the CYOS story")
	flag.Parse()
	fmt.Printf("Using the story in %s\n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err) // not the best way of error handling
	}

	story, err := cyos.JsonStory(f)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(storyHandlerTemplate))
	// Maps don't have a specific order, so in different executions we may end up
	// with different order in our printed stories
	//fmt.Printf("%+v\n", story)

	h := cyos.NewHandler(story, cyos.WithTemplate(tpl), cyos.WithPathFunc(pathFn))

	// Default handler
	//h1 := cyos.NewHandler(story)

	mux := http.NewServeMux()
	mux.Handle("/story/", h)

	fmt.Printf("Starting the server at: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	// "/intro" => "intro"
	return path[len("/story/"):]
}

var storyHandlerTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraph}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`
