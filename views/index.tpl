<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>PLAT QA周报平台</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="sunshengxiang01">

    <!-- Le styles -->
    <link href="/static/css/bootstrap.css" rel="stylesheet">
    <style type="text/css">
      body {
        padding-top: 60px;
        padding-bottom: 40px;
      }
      .sidebar-nav {
        padding: 9px 0;
      }
    </style>
    <link href="/static/css/bootstrap-responsive.css" rel="stylesheet">

    <!-- HTML5 shim, for IE6-8 support of HTML5 elements -->
    <!--[if lt IE 9]>
      <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->

    <!-- Fav and touch icons -->
    <link rel="apple-touch-icon-precomposed" sizes="144x144" href="/static/ico/apple-touch-icon-144-precomposed.png">
    <link rel="apple-touch-icon-precomposed" sizes="114x114" href="/static/ico/apple-touch-icon-114-precomposed.png">
    <link rel="apple-touch-icon-precomposed" sizes="72x72" href="/static/ico/apple-touch-icon-72-precomposed.png">
    <link rel="apple-touch-icon-precomposed" href="/static/ico/apple-touch-icon-57-precomposed.png">
    <link rel="shortcut icon" href="/static/ico/favicon.png">
  </head>

  <body>

    <div class="navbar navbar-inverse navbar-fixed-top">
      <div class="navbar-inner">
        <div class="container-fluid">
          <a class="btn btn-navbar" data-toggle="collapse" data-target=".nav-collapse">
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </a>
          <a class="brand" href="#">opqa周报管理平台</a>
          <div class="nav-collapse collapse">            
            <ul class="nav">
              <li class="active"><a href="/">首页</a></li>
              <li><a href="/admin">项目管理</a></li>
              <li><a href="/report">生成周报</a></li>
              <li><a href="/about">关于我们</a></li>
            </ul>
          </div><!--/.nav-collapse -->
        </div>
      </div>
    </div>

    <div class="container-fluid">
      <div class="row-fluid">
        <div class="span3">
          <div class="well sidebar-nav">
            <ul class="nav nav-list">
              <li class="nav-header">项目列表</li>
			{{if .Admin}}
				{{range .ProjList}}
			   	<li {{if .Cru}}class="active"{{end}}>
					<a href="/admin/p/{{.Name}}">{{.Name}}</a>
				</li>
				{{end}}
			{{else}}
				{{range .ProjList}}
			   	<li {{if .Cru}}class="active"{{end}}>
					<a href="/p/{{.Name}}">{{.Name}}</a>
				</li>
				{{end}}
			{{end}}
            </ul>
          </div><!--/.well -->
        </div><!--/span-->
        <div class="span9">
          <div class="hero-unit">
            {{if .Project}}
                <table class="table table-bordered">
                    <tr>
                        <td>QA</td>
                        <td>RD</td>
                    </tr>
                    <tr>
                        <td>{{range .QA}}{{.}} {{end}}</td>
                        <td>{{range .RD}}{{.}} {{end}}</td>
                    </tr>
                </table>
            {{end}}
            <div> {{markdown .Content}} </div>
			{{if .Admin}}
				<h2>添加新项目</h2>
				<form class="form-horizontal" method="post">
					<div class="control-group">
						<label class="control-label">项目名</label>
						<div class="controls">
							<input type="text" name="name" placeholder="eg: hermes" />
						</div>
					</div>
					<div class="control-group">
						<label class="control-label">QA</label>
						<div class="controls">
							<input type="text" name="qa" placeholder="qa name" />
						</div>
					</div>
					<div class="control-group">
						<label class="control-label">RD</label>
						<div class="controls">
							<input type="text" name="rd" placeholder="rd name" />
						</div>
					</div>
					<div class="control-group">
						<label class="control-label">项目类型</label>
						<div class="controls">
							<select name="type">
							<option value="quick">敏捷项目</option>
							<option value="heavy">重点项目</option>
							<option value="normal">普通项目</option>
							</select>	
						</div>
					</div>
					<div class="control-group">
						<label class="control-label">项目背景</label>
						<div class="controls">
							<textarea name="description" class="span8" rows="7" placeholder="项目背景"></textarea>
						</div>
					</div>
					<div class="form-actions">
						<button type="submit" class="btn btn-primary">添加</button>
						<button type="button" class="btn">取消</button>
					</div>
				</form>
			{{end}}
          </div>         
        </div><!--/span-->
      </div><!--/row-->

      <hr>

      <footer>
        <p>&copy; Company 2013</p>
      </footer>

    </div><!--/.fluid-container-->

    <!-- Le javascript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="/static/js/jquery.js"></script>
    <script src="/static/js/bootstrap-transition.js"></script>
    <script src="/static/js/bootstrap-alert.js"></script>
    <script src="/static/js/bootstrap-modal.js"></script>
    <script src="/static/js/bootstrap-dropdown.js"></script>
    <script src="/static/js/bootstrap-scrollspy.js"></script>
    <script src="/static/js/bootstrap-tab.js"></script>
    <script src="/static/js/bootstrap-tooltip.js"></script>
    <script src="/static/js/bootstrap-popover.js"></script>
    <script src="/static/js/bootstrap-button.js"></script>
    <script src="/static/js/bootstrap-collapse.js"></script>
    <script src="/static/js/bootstrap-carousel.js"></script>
    <script src="/static/js/bootstrap-typeahead.js"></script>

  </body>
</html>
