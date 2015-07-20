
<script src="/static/CodeMirror/lib/codemirror.js"></script>
<script src="/static/CodeMirror/addon/display/placeholder.js"></script>
<script src="/static/CodeMirror/mode/clike/clike.js"></script>
<script src="/static/CodeMirror/mode/go/go.js"></script>
<link rel="stylesheet" href="/static/CodeMirror/lib/codemirror.css">
<link rel="stylesheet" href="/static/CodeMirror/theme/ambiance.css">
<script type="text/javascript">
	$(document).ready(function () {
		config = {
			lineNumbers: true,
			mode: "clike",
			theme: "ambiance",
			indentWithTabs: true,
			value: "// Something",
		};
		editor = CodeMirror.fromTextArea(document.getElementById("paste_code"), config);

		$("[name='language']").change(function () {
			switch($("[name='language'] option:selected").text()) {
				case 'Cpp':
				case 'C':
				case 'Java': editor.setOption("mode", "clike"); break;
				case 'Go': editor.setOption("mode", "go"); break;
			}
		});
	})
</script>
