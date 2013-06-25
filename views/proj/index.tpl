<!DOCTYPE html>
<html>
  <head>
    <title>platqa jethome</title>
  </head>
  <body>
    <h1>Project</h1>
	<div>
		<form method="post">
			<table>
				<tr>
					<td>项目名</td>
					<td><input type="text" name="name" /></td>
				</tr>
				<tr>
					<td>QA</td>
					<td><input type="text" name="qa" />空格分隔</td>
				</tr>
				<tr>
					<td>RD</td>
					<td><input type="text" name="rd" />空格分隔</td>
				</tr>
				<tr>
					<td>项目背景</td>
					<td><textarea name="description" /></textarea></td>
				</tr>
				<tr>
					<td>项目种类</td>
					<td>
						<select name="type">
							<option value="0">敏捷项目</option>
							<option value="1">重点项目</option>
						</select>
					</td>
				</tr>
			</table>
			<input type="submit" value="Add" />
		</form>
	</div>
	<hr/>
	<div>
		<ul>
		{{range .ProjList}}
			<li>Id: {{.Id}} 类型{{.Type}} Name: <a href="/proj/{{.Name}}">{{.Name}}</a> QA: {{.QA}} RD:{{.RD}} Desc:{{.Description}}</li>
		{{end}}
		</ul>
	</div>
  </body>
</html>
