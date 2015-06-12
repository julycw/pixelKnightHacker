<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=no">
    <title>像素骑士团 - 别浪费时间了，该玩玩该工作工作！^_^</title>
    <link href="static/bootstrap/css/bootstrap.min.css" rel="stylesheet">
    <link href="static/css/docs.min.css" rel="stylesheet">
    <!--[if lt IE 9]>
      <script src="https://oss.maxcdn.com/html5shiv/3.7.2/html5shiv.min.js"></script>
      <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
    <script>
    var _hmt = _hmt || [];
    (function() {
      var hm = document.createElement("script");
      hm.src = "//hm.baidu.com/hm.js?6ce9ff98ea323424c8b7654e0d3f1642";
      var s = document.getElementsByTagName("script")[0]; 
      s.parentNode.insertBefore(hm, s);
    })();
    </script>
  </head>
  <body>
    {{if .message}}
    <div class="alert alert-danger" role="alert" id="top-error-tip">
      <p style="text-align:center;"><strong>出错啦！</strong> {{.message}}</p>
    </div>
    {{end}}
    <h1 style="text-align:center;">Have Fun ~</h1>
    <div class="container">
      <div class="row">
        <div class="col-sm-offset-2 col-sm-8">
          <div class="bs-callout bs-callout-info" id="callout-alerts-no-default">
            <h4>使用方法 - QQ群：<font style="font-weight:bold;">102623091</font></h4>
            <p>1. 找到游戏的真实地址，如3DM的4区就是http://bbs.3dmgame.com/u77.php?serverid=4</p>
            <p>2. 将该地址复制到浏览器地址栏中，按回车键应该会跳转到一个<font style="font-weight:bold;">新的地址</font></p>
            <p>3. 将<font style="font-weight:bold;">新的地址</font>复制到下面的输入框中，Let's play！^_^</p>
          </div>
        </div>
        <div class="col-xs-12">
          <form class="form-horizontal" action="/login" method="post">
            <div class="form-group">
              <label for="url" class="col-sm-2 control-label">URL</label>
              <div class="col-sm-8">
                <div class="input-group">
                  <input type="text" class="form-control" name="url" id="url">
                  <span class="input-group-btn">
                    <button type="submit" class="btn btn-default">Go ~</button>
                  </span>
                </div>
              </div>
            </div>
          </form>
        </div>
      </div>
    </div>
    <script src="static/js/jquery.min.js"></script>
    <script src="static/bootstrap/js/bootstrap.min.js"></script>
  </body>
</html>