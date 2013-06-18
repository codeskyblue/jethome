<!DOCTYPE html>
<html>
  <head>
    <title>platqa jethome</title>
  </head>
  <body>
    <h1>周报生成平台</h1>
	<div>
		<ul>
		<li><a href="proj">add new project</a></li>
		<li><a href="user">add new user</a></li>
		<li><a href="report">本周周报</a></li>
		</ul>
	</div>
	<hr/>
    <div>
		<h2>项目列表</h2>
        <ul>
        {{range .ProjList}}
            <li><a href="/proj/{{.Name}}">{{.Name}}</a></li>
        {{end}}
        </ul>
    </div>
	<hr/>
	<div>
        <h2>Leave a message</h2>
		<form  method="post">
			Info: <input type="text" name="info"/><br/>
			<input type="submit" />
		</form>
	</div>
	<hr/>
	<div class="info">
		<ul>
		{{range .Infos}}
			<li>{{.}}</li>
		{{end}}
		</ul>
	</div>
  </body>
</html>
