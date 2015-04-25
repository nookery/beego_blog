
function loadNav(tab) {
  html = '<div class="navbar-header">'+
            '<a class="navbar-brand" href="index.html">logo</a>'+
          '</div>'+
          '<div>'+
            '<ul class="nav navbar-nav">'+
               '<li><a href="index.html">导航</a></li>'+
               '<li><a href="grid.html">栅格系统</a></li>'+
               '<li><a href="typography.html">排版</a></li>'+
               '<li><a href="form.html">表单</a></li>'+
               '<li><a href="list.html">列表</a></li>'+
               '<li><a href="components.html">小型组件</a></li>'+
               '<li><a href="class.html">辅助类</a></li>'+
               '<li><a href="glyphicons.html">字体图标</a></li>'+
               '<li class="dropdown">'+
                  '<a class="dropdown-toggle" data-toggle="dropdown" href="#">'+
                     '实例页面 <span class="caret"></span>'+
                  '</a>'+
                  '<ul class="dropdown-menu">'+
                     '<li><a href="examples/frontpage.html" target="_blank">首页</a></li>'+
                     '<li><a href="#">jMeter</a></li>'+
                     '<li><a href="#">EJB</a></li>'+
                     '<li class="divider"></li>'+
                     '<li><a href="#">分离的链接</a></li>'+
                  '</ul>'+
                 '</li>'+
            '</ul>'+
          '</div>';
  $('#nav').html(html);
  $('#nav').find("a[href='"+tab+".html']").parent().addClass('active');
}