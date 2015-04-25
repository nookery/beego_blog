<!DOCTYPE html>
<html>
<head>
  <title></title>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width,in-scale=1">
  <link rel="stylesheet" type="text/css" href="/static/bootstrap-3.3.4-dist/css/bootstrap.min.css">
  <link rel="stylesheet" type="text/css" href="/static/css/frontPage.css">
</head>
<body>
  <nav class="navbar navbar-inverse navbar-fixed-top" role="navigation">
    <div class="container">
      <div class="navbar-header">
        <button type="button" class="navbar-toggle" data-toggle="collapse"
           data-target="#example-navbar-collapse">
           <span class="sr-only">切换导航</span>
           <span class="icon-bar"></span>
           <span class="icon-bar"></span>
           <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" href="#">W3Cschool</a>
     </div>
     <div class="collapse navbar-collapse" id="example-navbar-collapse">
        <ul class="nav navbar-nav">
           <li><a href="#">iOS</a></li>
           <li><a href="#">SVN</a></li>
           <li><a href="#">SVN</a></li>
           <li><a href="#">SVN</a></li>
           <li><a href="#">SVN</a></li>
           <li><a href="#">SVN</a></li>
           <li><a href="#">SVN</a></li>
        </ul>
     </div>
    </div>
  </nav>

  <div class="jumbotron">
    <h1>Beego Framework & Bootstrap
    </h1>
    <p class="lead">上面的导航是一个响应式导航，浏览器变窄时会发生变化</p>
    <p>
      <a class="btn btn-lg btn-success" href="http://www.gulpjs.com.cn/docs/getting-started/" role="button" target="_blank">开始使用</a>
    </p>
  </div>

  <div class="container">
    <div class="row">
      <div class="col-md-6">
        <h2>易于使用</h2>
        <p>没有 nginx 和 apache 居然可以自己干这个事情？是的，Go 其实已经做了网络层的东西，所以可以做到不需要 nginx 和 apache。</p>
      </div>
      <div class="col-md-6">
        <h2>构建快速</h2>
        <p>利用node流的威力，你可以快速构建项目并减少频繁的 IO 操作。</p>
     </div>
    </div>
    <div class="row">
      <div class="col-md-6">
        <h2>插件高质</h2>
        <p>gulp严格的插件指南确保插件如你期望的那样简洁高质得工作。</p>
      </div>
      <div class="col-md-6">
        <h2>易于学习</h2>
        <p>通过最少的API，掌握Gulp毫不费力，构建工作尽在掌握：如同一系列流管道。</p>
      </div>
    </div>
  </div>

  <footer class="footer">
    <div class="container">
      <p>© {{.website}}</p>
      <p>{{.email}}</p>
    </div>
  </footer>
</body>
</html>
<script src="../assets/js/jquery.min.js"></script>
<script src="../assets/bootstrap-3.3.4-dist/js/bootstrap.min.js"></script>