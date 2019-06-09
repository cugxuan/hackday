var sumUrl = "http://api.cugxuan.cn:8888/api/hackday/get_summary";

// var queryInfo = {
//     "active": true
// };

// chrome.tabs.query(queryInfo, function (tabs) {
//     _url = tabs[0].url;
//     _title = tabs[0].title;
// });

// $(document).ready(function () {


//     $.ajax({
//         type: "POST",
//         url: sumUrl,
//         data: _url,
//         dataType: "json",
//         success: function (response) {
//             console.log(response);
//         }
//     });




// });

var highlight = document.getElementsByTagName("p");

for (var i in highlight) {
    highlight[i].style.backgroundColor = "yellow";
}