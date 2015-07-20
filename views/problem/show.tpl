<!-- Problem Details page -->
<div class="problem-details">
	<h4>{{.problem.Statement}}</h4>
	<p><b>Problem Description</b><br/>{{.problem.Description}}</p>
	<p><b>Constraints</b><br/>{{.problem.Constraints}}</p>
	<p><b>Sample Input</b><br/>{{.problem.Sample_input}}</p>
	<p><b>Sample Output</b><br/>{{.problem.Sample_output}}</p>
</div>
<div class="write-code">
	<h5>Submit your code</h5>
	<form action="/problem/{{.problem.Pid}}/submit" method="POST">
		<textarea id="paste_code" name="code" placeholder="// Place Code here"></textarea><br/>
		<select name="language">
			<option value="C">C</option>
			<option value="Cpp">C++</option>
			<option value="Java">Java</option>
			<option value="Go">Go</option>
		</select>
		<div class="right"><input type="button" name="save" value="Save Draft"/><input type="button" name="run" value="Run"/><input type="submit"/></div>
	</form>
</div>