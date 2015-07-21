<!DOCTYPE html>
<html>
<head>
	<title>{{.title}} | OnlineJudge</title>
	<link rel="stylesheet" type="text/css" href="/static/css/reset.css">
	<link rel="stylesheet" type="text/css" href="/static/css/960.css">
	<link rel="stylesheet" type="text/css" href="/static/css/styles.css">
	<script type="text/javascript" src="/static/js/jquery-2.1.4.min.js"></script>
	<script type="text/javascript" src="/static/js/js.cookie.js"></script>
	{{ .HtmlHead }}
</head>

<body>
	<div class="container_24">
		<div id="header" class="grid_24" style="background-color:#eee;">
			<div class="grid_5 alpha">
				<h1 class="center" id="logo"><a href="/">OnlineJudge</a></h1>
			</div>
			<div class="grid_5 prefix_14 omega">
				<div class="login-button">
					{{ if .logged }}
						<a href="/user/{{ .login }}">{{ .login }}</a> <a href="/user/logout">(logout)</a>
						{{ if .isEditor }}
							<a href="/problem/create">+</a>
						{{ end }}
					{{ else }}
						<a href="/user/login">Login or Sign up</a>
					{{ end }}
				</div>
			</div>
		</div>

		<div class="clear"></div>

		<div class="grid_24">
			<p><span style="color:red;"> {{.flash.error}} </span>
				<span style="color:yellow;"> {{.flash.warning}} </span>
				<span style="color:blue;"> {{.flash.notice}} </span></p>
		</div>

		<div class="grid_18">
			<div id="content">
				{{ .LayoutContent }}
			</div>
		</div>
		<div class="grid_6">
			<div id="sidebar">
				{{ .Sidebar }}
			</div>
		</div>
	</div>
</body>
</html>