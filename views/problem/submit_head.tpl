
<script src="/static/CodeMirror/lib/codemirror.js"></script>
<script src="/static/CodeMirror/addon/display/placeholder.js"></script>
<script src="/static/CodeMirror/mode/clike/clike.js"></script>
<script src="/static/CodeMirror/mode/go/go.js"></script>
<script src="/static/CodeMirror/mode/python/python.js"></script>
<link rel="stylesheet" href="/static/CodeMirror/lib/codemirror.css">
<link rel="stylesheet" href="/static/CodeMirror/theme/ambiance.css">
<script type="text/javascript">
	$(document).ready(function () {
		config = {
			lineNumbers: true,
			mode: "clike",
			theme: "ambiance",
			indentWithTabs: true,
		};
		editor = CodeMirror.fromTextArea(document.getElementById("paste_code"), config);

		if(Cookies.get( ((document.location.pathname).split("/"))[2] + "_C") != null)
			editor.getDoc().setValue(Cookies.get( ((document.location.pathname).split("/"))[2] + "_C"));

		$("[name='language']").change(function () {
			switch($("[name='language'] option:selected").text()) {
				case 'Cpp':
				case 'C':
				case 'Java': editor.setOption("mode", "clike"); break;
				case 'Go': editor.setOption("mode", "go"); break;
				case 'Python3':
				case 'Python2': editor.setOption("mode", "python"); break;
			}
			if(Cookies.get( ((document.location.pathname).split("/"))[2] + "_" + $("[name='language'] option:selected").text()) != null)
				editor.getDoc().setValue(Cookies.get( ((document.location.pathname).split("/"))[2] + "_" + $("[name='language'] option:selected").text()));
			else
				editor.getDoc().setValue("");
		});

		$("[value='Save Draft'").click(function () {
			editor.save();
			Cookies.set( ((document.location.pathname).split("/"))[2] + "_" + $("[name='language'] option:selected").text() , $("#paste_code").val());
		});

		$("[value='Clear Draft'").click(function () {
			Cookies.remove( ((document.location.pathname).split("/"))[2] + "_" + $("[name='language'] option:selected").text());
			editor.getDoc().setValue("");
		});
	});
</script>
