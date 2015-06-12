var jjc = function(delay){
	if(delay == undefined){
		delay = 1000;
	}
	doCommand("jjc",{},delay);
}

var mineQueue = function(delay){
	if(delay == undefined){
		delay = 10000;
	}
	doCommand("mineQueue",{"mineType":$("#mineType").val()},delay);
}

var explore = function(delay){
	if(delay == undefined){
		delay = 10000;
	}
	doCommand("explore",{},delay);
}

var heartBeat = function(delay){
	if(delay == undefined){
		delay = 10000;
	}
	doCommand("heartBeat",{},delay,function(data){
		$heartBeatInfo.html(data.msg);
	});
}

var daywork = function(delay){
	if(delay == undefined){
		delay = 10000;
	}
	doCommand("daywork",{},delay);
}

var friends = function(delay){
	if(delay == undefined){
		delay = 10000;
	}
	doCommand("friends",{},delay);
}

function doCommand(command,args,delay,callback){
	if(args == undefined){
		args = {}
	}
	if(delay == undefined){
		delay = 10000
	}else if(delay == 0){
		args["action"] = command;
	}
	setTimeout(function(){
		$.post("/command/"+command,args,function(data,state){
			if(state === "success"){
				data.command = command;
				printLog(data);
				if(callback != undefined){
					callback(data);
				}
				errTip(false);
			}else{
				errTip(true);
			}
		});
	},getRandomNum(0,delay));
	insertCountdown("#count-down-"+command,functions[command].interval*1000);
}

function getRandomNum(min,max){   
	var range = max - min;   
	var rand = Math.random();   
	return(min + Math.round(rand * range));   
}

function printInfo(info){
	var line = "<p>"+info+"</p>";
	if(logCount >= maxLogCount){
		$logWrapper.children("p:last").remove();
	}else{
		logCount++
	}
	$logWrapper.prepend(line);
}

function printLog(log){
	var line = "";
	if(log.ret == 0){
		line += "<p>"+log.time+" - ["+functions[log.command].title+"]"+log.msg+"</p>";
	}else{
		line += "<p style=\"color:orange;\">"+log.time+" - ["+functions[log.command].title+"]"+log.msg+"</p>";
	}

	if(logCount >= maxLogCount){
		$logWrapper.children("p:last").remove();
	}else{
		logCount++
	}
	$logWrapper.prepend(line);
}

function clearLog(){
	$logWrapper.empty();
	logCount = 0;
}

function errTip(visible){
	if(visible){
		$("#top-error-tip").show();
	}else{
		$("#top-error-tip").hide();
	}
}

function eventBind(){
	//防止多次重复点击按钮
	$(document).on("click",".disabled-multi-click",function(){
		var $this = $(this);
		$this.attr("disabled","disabled");
		setTimeout(function(){
			$this.removeAttr("disabled");
		},1000);
	});
}

function timeFormat(time){
	var h=Math.floor(time/1000/60/60);
    var m=Math.floor(time/1000/60%60);
    var s=Math.floor(time/1000%60);
    var hStr = h>=10?h:"0"+h;
    var mStr = m>=10?m:"0"+m;
    var sStr = s>=10?s:"0"+s;
    var timeStr = hStr+":"+mStr+":"+sStr;
    return timeStr;
}

function insertCountdown(selector,timemax){
	var interval = intervals[selector];
	if(interval!=undefined){
		clearInterval(interval);
		intervals[selector] = undefined;
	}
	var $wrapper = $(selector);
    $wrapper.html(timeFormat(timemax));
	interval = setInterval(function(){
		timemax -= 1000;
	    $wrapper.html(timeFormat(timemax));
	},1000);
	intervals[selector] = interval;
	setTimeout(function(){
		clearInterval(interval);
		intervals[selector] = undefined;
	},timemax+500);
}

var logCount = 0;
var maxLogCount = 200;
var $logWrapper = undefined;
var $heartBeatInfo = undefined;
var functions = {};
var intervals = {};

$(function(){
	errTip(false);
	$logWrapper = $("#log-wrapper");
	$heartBeatInfo = $("#heartBeatInfo");
	//事件绑定
	eventBind();
	//功能列表
	functions["jjc"] = {title:"竞技场",interval:60*60+10,func:jjc};
	functions["mineQueue"] ={title:"矿队",interval:30*60,func:mineQueue};
	functions["explore"] ={title:"冒险",interval:10*60,func:explore};
	functions["heartBeat"] = {title:"打酱油",interval:2*60,func:heartBeat};
	functions["daywork"] = {title:"日常",interval:24*60*60,func:daywork};
	functions["friends"] = {title:"好友",interval:1*60*60,func:friends};
	for(var key in functions){
		setInterval(functions[key].func,functions[key].interval*1000);
		insertCountdown("#count-down-"+key,functions[key].interval*1000);
	}  
	printInfo("Enjoy~");
	heartBeat(1);
});