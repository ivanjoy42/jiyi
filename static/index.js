BackUrl = "";
BodyXHR = null;
Deep = 0;
Home = "learn/index";

$(function () {
    url = GetHash();
    Href(url, 0);

    $.ajaxSetup({
        cache: true,
    });
});

function UpdateVer() {
    ver = GetCookie("ver");

    $('link[rel="stylesheet"]').each(function () {
        latest = this.href.replace(/\?.*|$/, "?" + ver);
        if (this.href != latest) {
            this.href = latest;
        }
    });

    $("script").each(function () {
        if (this.src == "" || this.src.indexOf("jquery") >= 0) {
            return;
        }
        latest = this.src.replace(/\?.*|$/, "?" + ver);
        if (this.src != latest) {
            this.remove();
            $("<script>").attr("src", latest).appendTo("head");
        }
    });
}

function GetCookie(name) {
    var arr,
        reg = new RegExp("(^| )" + name + "=([^;]*)(;|$)");
    if ((arr = document.cookie.match(reg))) {
        return arr[2];
    } else {
        return null;
    }
}

function GetHash() {
    if (location.hash == "") {
        url = Home;
    } else {
        url = location.hash.slice(1);
    }
    return url;
}

function Href(url, direct) {
    if (!url.match(/.+/)) {
        return;
    }

    if (BodyXHR) {
        BodyXHR.abort();
    }

    if (direct < 0 && Deep > 0) {
        BackUrl = url;
        history.go(direct);
        return;
    } else if (direct == undefined && Deep > 0) {
        BackUrl = url;
        history.go(-Deep);
        Deep = 1;
        return;
    } else {
        BodyXHR = $.ajax({
            url: url,
            success: function (res) {
                UpdateHistory(url, direct);
                $("body").html(res);
                autofocus();
                scrollTo(0, 0);
                UpdateVer();
            },
        });
    }
}

function autofocus() {
    e = $("input[autofocus],textarea[autofocus]");
    tmp = e.val();
    e.val("").focus().val(tmp);
}

function UpdateHistory(url, direct) {
    if (direct == 1) {
        history.pushState(url, null, "#" + url);
        Deep++;
    } else if (direct == 0) {
        history.replaceState(url, null, "#" + url);
    }
}

onpopstate = function (event) {
    if (event.state == null) {
        url = GetHash();
        Deep++;
    } else if (BackUrl != "") {
        url = BackUrl;
        BackUrl = "";
        Deep--;
    } else {
        url = event.state;
        Deep--;
    }
    Href(url, 0);
};

function Submit(url, direct) {
    var data = {};
    $.each($("form").serializeArray(), function () {
        data[this.name] = this.value;
    });

    $.ajax({
        type: "POST",
        url: url,
        data: data,
        success: function (res) {
            if (direct < 0) {
                history.go(direct);
                Deep = Deep + direct + 1;
            } else {
                $("body").html(res);
                scrollTo(0, 0);
            }
        },
    });
    return false;
}

function setRadio(inputName, kindId, e) {
    $(e).siblings("button").attr("class", "white")
    $(e).attr("class", "green")
    $("input[name=" + inputName + "]").val(kindId)
}