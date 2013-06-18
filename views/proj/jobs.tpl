<!DOCTYPE html>
<html>
  <head>
    <title>platqa jethome</title>
  </head>
  <body>
    <h1>Jethome</h1>
	<div>
		<form method="post">
			<input type="hidden" name="name" value="{{.Jobs.Name}}" /><br/>
			本周内容: <textarea type="text" name="content"></textarea><br/>
			<input type="submit" value="Add Job" />
		</form>
	</div>
	<hr/>
	<div>
	<h3>项目名: {{.Jobs.Name}}</h2>
	<p>项目背景: {{.Jobs.Description}}</p>
		<ul>
		
		{{range .Jobs.Content}}
			<li>{{.}}</li>
		{{end}}
		</ul>
	</div>
  </body>
</html>
