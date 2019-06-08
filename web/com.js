var sumUrl = "http://api.cugxuan.cn:8888/api/hackday/get_summary";

var queryInfo = {
    "active": true
};

chrome.tabs.query(queryInfo, function (tabs) {
    _url = tabs[0].url;
    _title = tabs[0].title;
});

$(document).ready(function () {

    console.log("hello");

    $.ajax({
        type: "POST",
        url: sumUrl,
        data: _url,
        dataType: "json",
        success: function (response) {
            console.log(response);
        }
    });
});