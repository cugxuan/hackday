var tag;

var url = "http://api.cugxuan.cn/api/hackday/get_hot";

$(document).ready(function () {

    $.ajax({
        type: "POST",
        url: url,
        data: tag,
        dataType: "json",
        success: function (response) {
            if($(".text")) {
                $(".text").remove();
                console.log("321");
            } 
            var text = document.createElement("div");
            text.className = 'text';
            $(".rank").append(text);
            for (var i in response.data) {
                var x = document.createElement("div");
                x.className = "rank" + i;
                $(".text").append(x);
                var title = document.createElement("p");
                title.innerText = response.data[i].title;
                title.className = "title";
                x.appendChild(title);
                var count = document.createElement("p");
                count.innerText = response.data[i].count;
                count.className = "count";
                x.appendChild(count);
            }
            }
    })
            
});


$("#submit").click(function () { 
    tag = $("#input").val();
    console.log(tag);
    $.ajax({
        type: "POST",
        url: url,
        data: tag,
        dataType: "json",
        success: function (response) {
            if($(".text")) {
                $(".text").remove();
                console.log("123");
            }   
            var text = document.createElement("div");
            text.className = 'text';
            $(".rank").append(text);
            for (var i in response.data) {
                var x = document.createElement("div");
                x.className = "rank" + i;
                $(".text").append(x);
                var title = document.createElement("p");
                title.innerText = response.data[i].title;
                title.className = "title";
                x.appendChild(title);
                var count = document.createElement("p");
                count.innerText = response.data[i].count;
                count.className = "count";
                x.appendChild(count);
            }
            }
    })
    
});

        // var jsonObj = [
        //     {"title": "123","count": "1"},
        //     {"title": "12342","count": "4"},
        //     {"title": "123123","count": "7"},
        //     {"title": "123123","count": "8"}
        // ];
        
        






