<!-- Problems List Page -->
<div class="problems-list">
	{{range $key, $val := .problems}}
	<div class="problem">
		<h5><a href="/problem/{{.Pid}}">{{.Statement}}</a></h5>
		<p><span>Type: {{.Type}}</span><span>Difficulty: {{.Difficulty}}</span><span>Solved by: {{.Solve_count}}</span><span>Created: {{.Created_at}}</span></p>
	</div>
	{{end}}
</div>