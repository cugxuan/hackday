var link;//网页链接
var comUrl = "http://api.cugxuan.cn/api/hackday/send_share";//分享事件端口

var queryInfo = {
    "active": true
};

chrome.tabs.query(queryInfo, function (tabs) {
    link = tabs[0].url;
});

//提交分享事件
$("#confirm").click(function () { 

    var sum = $("#summary").val();
    var com = $("#comment").val();

    // console.log('summary:' + sum +'comment:' + com + 'link:' + link);

    var data = {
        "summary": sum,
        "comment": com,
        "link": link,
    }

    //console.log(data);

    $.ajax({
        type: "POST",
        url: comUrl,
        data: data,
        dataType: "json",
        //成功调试 最后记得删除
        success: function (response) {
            console.log(response.parse());
        }
    });
});
