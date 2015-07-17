<!-- Create Problem Page -->
<div class="margin-15">
	<h4>Create Problem</h4>
	<form method="POST">
		<label>Title: <input type="text" name="statement"></label><br/><br/>
		<label>Description: <br/><textarea name="description" style="width:50%"></textarea></label><br/><br/>
		<label>Constraints: <br/><textarea name="constraints" style="width:50%"></textarea></label><br/><br/>
		<label>Sample Input: <br/><textarea name="sample_input" style="width:50%"></textarea></label><br/><br/>
		<label>Sample Output: <br/><textarea name="sample_output" style="width:50%"></textarea></label><br/><br/>
		<label>Type: <input type="text" name="type"></label><br/><br/>
		<label>Difficulty: <select type="text" name="difficulty"><option value="Beginner">Beginner</option><option value="Novice">Novice</option><option value="Expert">Expert</option></select></label><br/><br/>
		<label>Points: <input type="number" name="points"></label><br/><br/>
		<input type="submit">
	</form>
</div>