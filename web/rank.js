var tag;

var url = "http://api.cugxuan.cn:8888/api/hackday/get_hot";

$(document).ready(function () {

    writeText();
            
});


$("#submit").click(function () { 
    tag = $("#input").val();


    console.log(tag);

    writeText();
    
});


function writeText() {
    $.ajax({
        type: "POST",
        url: url,
        data: tag,
        dataType: "json",
        success: function (response) {
            console.log(response);
            if($(".text")) {
                $(".text").remove();
            } 
            var text = document.createElement("div");
            text.className = 'text';
            $(".rank").append(text);
            for (var i in response.data) {
                var x = document.createElement("div");
                x.className = "rank" + i;
                $(".text").append(x);
                var link = document.createElement("a");
                x.appendChild(link);
                link.href = response.data[i].link;
                var title = document.createElement("p");
                title.innerText = response.data[i].title;
                title.className = "title";
                link.appendChild(title);
                var count = document.createElement("p");
                count.innerText = response.data[i].count;
                count.className = "count";
                link.appendChild(count);
            }
            }
    })
};
        






