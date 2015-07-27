<h4>Recently Solved by</h4>
{{ if .recentlySolvedUsersExist }}
	{{range $key, $val := .recentlySolvedUsers}}
	<a href="/user/show/{{.Username}}">{{.Username}}</a><br/>
	{{end}}
{{ else }}
	<span>Be the first to solve!</span>
{{ end }}
