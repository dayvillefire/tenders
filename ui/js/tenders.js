// Tenders - @jbuchbinder
// Application logic

//var config = {};
var NOTIFICATION_TIMEOUT = 3;

/*
$(document).ready(function () {
    // Load local UI config before we go any further
    $.getJSON("/api/config", {}, function (response) {
        config = response;
        console.log(config);
    });

    //setInterval(showClock, 1000);
});
*/

function urlParam(name) {
    var results = new RegExp('[\?&]' + name + '=([^&#]*)').exec(window.location.href);
    return results[1] || 0;
}

function showClock() {
    var d = new Date();
    d.setMinutes(d.getMinutes() - d.getTimezoneOffset());
    $('#clock').html(d.toISOString().slice(0, 19).replace('T', ' '));
}