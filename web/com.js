var sumUrl = "http://api.cugxuan.cn/api/hackday/get_summary";

var _url = window.location.href;
var _title = window.document.title;




$(document).ready(function () {

    

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

// $("p").style.backgroundColor = "yellow";

var highlight = document.getElementsByTagName("p");

for (var i in highlight) {
    highlight[i].style.backgroundColor = "yellow";
}




response.data[0].summary