package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

var tpl *template.Template

var defaultTemplate string = `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
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

func init() {
	tpl = template.Must(template.New("").Parse(defaultTemplate))
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

/*This is the newHandler function
-- Returns the http.Handler type.... So the return has to implement the ServeHttp function... The values which will be used to create the struct which will be finally implementing
the ServeHttp method will have all the values from the struct
*/
func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//First lets find what is coming in the input requestt Path
	path := req.URL.Path

	if val, ok := h.s[strings.Trim(path, "/")]; ok && path != "intro" {
		tpl.Execute(w, val)
		fmt.Println("Check 1")
	} else {
		fmt.Println("Check 3")
		tpl.Execute(w, h.s["intro"])
	}

}

func JsonStory(read io.Reader) (Story, error) {
	decoder := json.NewDecoder(read)
	var story1 Story
	if err := decoder.Decode(&story1); err != nil {
		return nil, err
		log.Fatal("I am not able to decode the file into a Json")
	}

	return story1, nil

}
