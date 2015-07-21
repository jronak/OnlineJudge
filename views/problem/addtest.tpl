<!-- Add Testcase Page -->
<div class="margin-15">
	<h4>Tests Present for <a href="/problem/{{.problem.Pid}}">{{.problem.Statement}}</a></h4>
	{{range $key, $val := .cases}}
		<div class="margin-15 bottom-border">
			<p><b>Input:</b><br/>{{str2html .Input}}<br/>
				<b>Output:</b><br/>{{str2html .Output}}<br/>
				<b>Timeout:</b><br/>{{.Timeout}}<br/>
			</p>
		</div>
	{{end}}
	<h4>Create Test Case</h4>
	<form method="POST">
		<label>Input: <br/><textarea name="input" style="width:50%"></textarea></label><br/><br/>
		<label>Output: <br/><textarea name="output" style="width:50%"></textarea></label><br/><br/>
		<label>Timeout: <select type="number" name="timeout">
			<option value="2">2</option>
			<option value="5">5</option>
			<option value="10">10</option>
		</select>
		</label><br/><br/>
		<input type="submit">
	</form>
</div>