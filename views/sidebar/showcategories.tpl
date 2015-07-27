<h4>Categories</h4>
{{range $key, $val := .types}}
	<a href="/problem/{{.}}/1">{{.}}</a><br/>
{{end}}