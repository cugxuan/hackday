var _width = 500,
    _height = 300,
    _top = (screen.height - _height) / 2,
    _left = (screen.width - _width) / 2;

 $(".share").click(function () {
    window.open("share.html", "_blank", 'width=' + _width + ',height=' + _height + ',top=' + _top + ',left=' + _left + ',toolbar = 0, menubar = no, scrollbars = no, resizable = 1, location = no, status = 0');
});

$(".rank").click(function () {
    window.open("rank.html", "_blank", 'width=' + _width + ',height=' + _height + ',top=' + _top + ',left=' + _left + ',toolbar = 0, menubar = no, scrollbars = no, resizable = 1, location = no, status = 0');
});

$(".com").click(function () {
    window.open("com.html", "_blank", 'width=' + _width + ',height=' + _height + ',top=' + _top + ',left=' + _left + ',toolbar = 0, menubar = no, scrollbars = no, resizable = 1, location = no, status = 0');
});
