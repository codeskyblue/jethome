<!DOCTYPE html>
<html>
  <head>
    <title>platqa jethome</title>
  </head>
  <body>
    <h1>Jethome</h1>
	<div>
		<form method="post">
			Name: <input type="text" name="name" /><br/>
			Email: <input type="text" name="email" /><br/>
			<input type="submit" value="Add" />
		</form>
	</div>
	<hr/>
	<div>
		<ul>
		{{range .UserList}}
			<li>Id: {{.Id}} Name: {{.Name}} Email:{{.Email}}</li>
		{{end}}
		</ul>
	</div>
  </body>
</html>
