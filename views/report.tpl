<!DOCTYPE html>
<html>
  <head>
    <title>platqa jethome</title>
  </head>
  <body>
    <h1>项目周报</h1>
	<hr/>
	<div>
		<h2>重点项目</h2>
		<table>
			<tr>
				<td>项目名称</td>
				<td>RD</td>
				<td>QA</td>
				<td>本周内容</td>
			</tr>
			{{range .HeavyJobs}}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.RD}}</td>
				<td>{{.QA}}</td>
				<td>{{.Content}}</td>
			</tr>	
			{{end}}
		</table>
	</div>
	<div>
		<h2>敏捷项目</h2>
		<table>
			<tr>
				<td>项目名称</td>
				<td>RD</td>
				<td>QA</td>
				<td>本周内容</td>
			</tr>
			{{range .QuickJobs}}
			<tr>
				<td>{{.Name}}</td>
				<td>{{.RD}}</td>
				<td>{{.QA}}</td>
				<td>{{.Content}}</td>
			</tr>	
			{{end}}
		</table>
	</div>
  </body>
</html>
