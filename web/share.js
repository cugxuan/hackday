var comUrl = "http://api.cugxuan.cn:8888/api/hackday/send_share";

var _title,_source,_sourceUrl,_pic,_showcount,_desc,_summary,_site,_url,
            _width = 600,
            _height = 600,
            _top = (screen.height-_height)/2,
            _left = (screen.width-_width)/2,
            _pic = '';

var queryInfo = {
    "active": true
};

var sum,com;

chrome.tabs.query(queryInfo, function (tabs) {
    _url = tabs[0].url;
    _title = tabs[0].title;
});
 
//提交分享事件
$("#confirm").click(function () { 

    sum = $("#summary").val();
    com = $("#comment").val();

    var data = {
        "summary": sum,
        "comment": com,
        "link": _url,
        "title": _title,
    }

    console.log(data);

    $.ajax({
        type: "POST",
        url: comUrl,
        data: data,
        dataType: "json",
    });

    var code = {
        text: _url,
        size: 90,
        background: "#ffffff",
        foreground: "#000000" 
    
    }
    $("#qrcode").qrcode(code);

    $("#card").attr("src", "images/card.png");
    $("#card").css("width", "800px");
    $(".sum").remove();
    $(".com").remove();
    $("#confirm").remove();


    var text1 = document.createElement("p");
    var text2 = document.createElement("p");

    text1.className = "text1";
    text2.className = "text2";

    text1.innerText = sum;
    text2.innerText = com;

    document.body.appendChild(text1);
    document.body.appendChild(text2);


});




//     //分享到新浪微博
// $("#sina").click(function () { 
//     var _shareUrl = 'http://v.t.sina.com.cn/share/share.php?&appkey=895033136';     //真实的appkey，必选参数
//         _shareUrl += '&url='+ encodeURIComponent(_url);     //参数url设置分享的内容链接|默认当前页location，可选参数
//         _shareUrl += '&title=' + encodeURIComponent(_title);    //参数title设置分享的标题|默认当前页标题，可选参数
//         _shareUrl += '&source=' + encodeURIComponent(_source||'');
//         _shareUrl += '&sourceUrl=' + encodeURIComponent(_sourceUrl||'');
//         _shareUrl += '&content=' + 'utf-8';   //参数content设置页面编码gb2312|utf-8，可选参数
//         _shareUrl += '&pic=' + encodeURIComponent(_pic||'');  //参数pic设置图片链接|默认为空，可选参数
//         window.open(_shareUrl,'_blank','width='+_width+',height='+_height+',top='+_top+',left='+_left+',toolbar=no,menubar=no,scrollbars=no, resizable=1,location=no,status=0');
// });
//     //分享到QQ空间
// $("#qzone").click(function () { 
//         var _shareUrl = 'http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?';
//         _shareUrl += 'url=' + encodeURIComponent(_url);   //参数url设置分享的内容链接|默认当前页location
//         _shareUrl += '&showcount=' + _showcount||0;      //参数showcount是否显示分享总数,显示：'1'，不显示：'0'，默认不显示
//         _shareUrl += '&desc=' + encodeURIComponent(_desc||'分享的描述');    //参数desc设置分享的描述，可选参数
//         _shareUrl += '&summary=' + encodeURIComponent(_summary||'分享摘要');    //参数summary设置分享摘要，可选参数
//         _shareUrl += '&title=' + encodeURIComponent(_title);    //参数title设置分享标题，可选参数
//         _shareUrl += '&site=' + encodeURIComponent(_site||'');   //参数site设置分享来源，可选参数
//         _shareUrl += '&pics=' + encodeURIComponent(_pic||'');   //参数pics设置分享图片的路径，多张图片以＂|＂隔开，可选参数
//         window.open(_shareUrl,'_blank','width='+_width+',height='+_height+',top='+_top+',left='+_left+',toolbar=no,menubar=no,scrollbars=no,resizable=1,location=no,status=0');
// });

//     //分享到百度贴吧
// $("#tieba").click(function () { 
        
//         var _shareUrl = 'http://tieba.baidu.com/f/commit/share/openShareApi?';
//         _shareUrl += 'title=' + encodeURIComponent(_title||document.title);  //分享的标题
//         _shareUrl += '&url=' + encodeURIComponent(_url||document.location);  //分享的链接
//         _shareUrl += '&pic=' + encodeURIComponent(_pic||'');    //分享的图片
//         window.open(_shareUrl,'_blank','width='+_width+',height='+_height+',left='+_left+',top='+_top+',toolbar=no,menubar=no,scrollbars=no,resizable=1,location=no,status=0');
// });

//     //分享到豆瓣
// $("#douban").click(function () { 
//         var _shareUrl = 'http://shuo.douban.com/!service/share?';
//         _shareUrl += 'href=' + encodeURIComponent(_url||location.href);    //分享的链接
//         _shareUrl += '&name=' + encodeURIComponent(_title||document.title);    //分享的标题
//         _shareUrl += '&image=' + encodeURIComponent(_pic||'');    //分享的图片
//         window.open(_shareUrl,'_blank','width='+_width+',height='+_height+',left='+_left+',top='+_top+',toolbar=no,menubar=no,scrollbars=no,resizable=1,location=no,status=0');
// });


//     //分享到facebook
// $("#fackbook").click(function () { 
//         var _shareUrl = 'http://www.facebook.com/sharer/sharer.php?';
//         _shareUrl += 'u=' + encodeURIComponent(_url||location.href);    //分享的链接
//         _shareUrl += '&t=' + encodeURIComponent(_title||document.title);    //分享的标题
//         window.open(_shareUrl,'_blank','width='+_width+',height='+_height+',left='+_left+',top='+_top+',toolbar=no,menubar=no,scrollbars=no,resizable=1,location=no,status=0');
// });


//     //分享到Twitter
// $("#twitter").click(function () { 
//         var _shareUrl = 'http://twitter.com/intent/tweet?';
//         _shareUrl += 'url=' + encodeURIComponent(_url||location.href);    //分享的链接
//         _shareUrl += '&text=' + encodeURIComponent(_title||document.title);    //分享的标题
//         window.open(_shareUrl,'_blank','width='+_width+',height='+_height+',left='+_left+',top='+_top+',toolbar=no,menubar=no,scrollbars=no,resizable=1,location=no,status=0');
// });