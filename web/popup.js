 $(".share").click(function () {
     _width = 840;
     _height = 540;
     _top = (screen.height - _height) / 2,
     _left = (screen.width - _width) / 2;
    window.open("share.html", "_blank", 'width=' + _width + ',height=' + _height + ',top=' + _top + ',left=' + _left + ',toolbar = 0, menubar = no, scrollbars = no, resizable = 1, location = no, status = 0');
});

$(".rank").click(function () {
    _width = 1160;
    _height = 1080;
    _top = (screen.height - _height) / 2,
    _left = (screen.width - _width) / 2;
    window.open("rank.html", "_blank", 'width=' + _width + ',height=' + _height + ',top=' + _top + ',left=' + _left + ',toolbar = 0, menubar = no, scrollbars = no, resizable = 1, location = no, status = 0');
});

$(".com").click(function () {

    chrome.tabs.executeScript(null, {file: "js/jquery-3.4.1.js"});
    chrome.tabs.executeScript(null, {file: "com.js"});
});
