<!-- User Login Page -->
<div class="margin-15">
	<div>
		<h4>Login</h4>
		<form action="/user/login/submit" method="POST">
			<label>Username: <input type="text" name="username"></label><br/><br/>
			<label>Password: <input type="password" name="password"></label><br/><br/>
			<input type="submit"><br/><br/>
		</form>
	</div>
	<div>
		<h4>Don't have an account? Sign up!</h4>
		<form action="/user/signup/submit" method="POST">
			<label>Username: <input type="text" name="username"></label><br/><br/>
			<label>Password: <input type="password" name="passkey"></label><br/><br/>
			<label>Name: <input type="text" name="name"></label><br/><br/>
			<label>College: <input type="text" name="college"></label><br/><br/>
			<label>EMail: <input type="email" name="email"></label><br/><br/>
			<input type="submit">
		</form>
	</div>
</div>