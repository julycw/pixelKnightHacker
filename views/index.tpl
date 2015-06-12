<!DOCTYPE html>
<html lang="zh-cn">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1, user-scalable=no">
    <title>像素骑士团 - 别浪费时间了，该玩玩该工作工作！^_^</title>
    <link href="static/bootstrap/css/bootstrap.min.css" rel="stylesheet">
    <link href="static/css/docs.min.css" rel="stylesheet">
    <link href="static/css/site.css" rel="stylesheet">
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
    <div class="alert alert-danger" role="alert" id="top-error-tip">
      <p style="text-align:center;"><strong>出错啦！</strong> 似乎网络断了</p>
    </div>
    <div class="container">
      <div class="row">
        <div class="col-md-12">
          <h5>
            <span>
              {{.UserName}}|
              <a href="#" data-toggle="modal" data-target="#qrCodeWrapper">
                转战手机
              </a>|<a href="/logout">退出</a> / v{{.Version}}
            </span> / 
            <span>
              概况:<span id="heartBeatInfo">获取中...</span>
            </span> / 
            <span>
              打酱油: <span id="count-down-heartBeat">00:00:00</span>
            </span>
          </h5>
        </div>
      </div>
      <div class="row function-list">
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:jjc(0);" class="btn btn-primary btn-block disabled-multi-click" role="button">竞技场</a>
            <div id="count-down-jjc" class="count-down">00:00:00</div>
          </div>
        </div>
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:mineQueue(0);" class="btn btn-primary btn-block disabled-multi-click" role="button">矿队</a>
            <div id="count-down-mineQueue" class="count-down">00:00:00</div>
          </div>
        </div>
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:explore(0);" class="btn btn-primary btn-block disabled-multi-click" role="button">冒险</a>
            <div id="count-down-explore" class="count-down">00:00:00</div>
          </div>
        </div>
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:daywork(0);" class="btn btn-primary btn-block disabled-multi-click" role="button">日常</a>
            <div id="count-down-daywork" class="count-down">00:00:00</div>
          </div>
        </div>
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:friends(0);" class="btn btn-primary btn-block disabled-multi-click" role="button">好友</a>
            <div id="count-down-friends" class="count-down">00:00:00</div>
          </div>
        </div>
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:void(0);" class="btn btn-primary btn-block disabled-multi-click" role="button" disabled="disabled">功能</a>
            <div class="count-down">00:00:00</div>
          </div>
        </div>
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:void(0);" class="btn btn-primary btn-block disabled-multi-click" role="button" disabled="disabled">功能</a>
            <div class="count-down">00:00:00</div>
          </div>
        </div>
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:void(0);" class="btn btn-primary btn-block disabled-multi-click" role="button" disabled="disabled">功能</a>
            <div class="count-down">00:00:00</div>
          </div>
        </div>
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:void(0);" class="btn btn-primary btn-block disabled-multi-click" role="button" disabled="disabled">功能</a>
            <div class="count-down">00:00:00</div>
          </div>
        </div>
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:void(0);" class="btn btn-primary btn-block disabled-multi-click" role="button" disabled="disabled">功能</a>
            <div class="count-down">00:00:00</div>
          </div>
        </div>
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:void(0);" class="btn btn-primary btn-block disabled-multi-click" role="button" disabled="disabled">功能</a>
            <div class="count-down">00:00:00</div>
          </div>
        </div>
        <div class="col-xs-3 col-sm-2 col-md-1">
          <div class="function-wrapper">
            <a href="javascript:clearLog();" class="btn btn-success btn-block" role="button">清屏</a>
          </div>
        </div>
      </div>
      <div class="devide"></div>
      <div class="row">
        <div class="col-xs-12 col-sm-12 col-md-6">
          <form class="form-horizontal">
            <div class="form-group">
              <label for="mineType" class="col-sm-2 control-label">矿物深度</label>
              <div class="col-sm-10">
                <select class="form-control" id="mineType">
                  <option value="1" selected>浅层</option>
                  <option value="2">中层</option>
                  <option value="3">深层 - 消耗1地雷</option>
                  <option value="4">地心 - 消耗1地雷</option>
                </select>
              </div>
            </div>
            <div class="form-group">
              <div class="col-sm-offset-2 col-sm-8">
                <p>QQ群：<font style="font-weight:bold;">102623091</font></p>
              </div>
            </div>
          </form>
        </div>
        <div class="col-xs-12 col-sm-12 col-md-6">
          <div class="well well-sm" id="log-wrapper">
          </div>
        </div>
      </div>
    </div>
    

    <!-- Modal -->
    <div class="modal fade" id="qrCodeWrapper" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="modal-header">
            <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
            <h4 class="modal-title" id="myModalLabel">扫一扫</h4>
          </div>
          <div class="modal-body">
            <div id="qrcodeCanvas" style="text-align:center;"></div>
          </div>
          <div class="modal-footer">
            <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
          </div>
        </div>
      </div>
    </div>
    <script src="static/js/jquery.min.js"></script>
    <script src="static/bootstrap/js/bootstrap.min.js"></script>
    <script src="static/js/jquery.qrcode.js"></script>
    <script src="static/js/qrcode.js"></script>
    <script src="static/js/site.js"></script>
    <script>
      $('#qrcodeCanvas').qrcode({
         text: window.location.href,
         width: 277,//二维码宽度
         height: 277//二维码高度
      });
    </script>
  </body>
</html>