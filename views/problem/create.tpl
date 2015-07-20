<!-- Create Problem Page -->
<div class="margin-15">
	<h4>Create Problem</h4>
	<form method="POST">
		<label>Title: <input type="text" name="statement"></label><br/><br/>
		<label>Description: <br/><textarea name="description" style="width:50%"></textarea></label><br/><br/>
		<label>Constraints: <br/><textarea name="constraints" style="width:50%"></textarea></label><br/><br/>
		<label>Sample Input: <br/><textarea name="sample_input" style="width:50%"></textarea></label><br/><br/>
		<label>Sample Output: <br/><textarea name="sample_output" style="width:50%"></textarea></label><br/><br/>
		<label>Type: <select type="text" name="type">
			<option value="Data Structure" >Data Structure</option>
			<option value="Sorting">Sorting</option>
			<option value="Graph Theory">Graph Theory</option>
		</select></label><br/><br/>
		<label>Difficulty: <select type="text" name="difficulty">
			<option value="Beginner">Beginner</option>
			<option value="Novice">Novice</option>
			<option value="Expert">Expert</option>
		</select></label><br/><br/> 
		<label>Points: <select type="number" name="points">
			<option value=10>10</option>
			<option value=20>20</option>
			<option value=30>30</option>
		</label><br/><br/>
		<input type="submit">
	</form>
</div>