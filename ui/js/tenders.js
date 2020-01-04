// Tenders - @jbuchbinder
// Application logic

var config = {};

$(document).ready(function () {
    // Load local UI config before we go any further
    $.getJSON("/api/config", {}, function (response) {
        config = response;
        console.log(config);
    });

    //setInterval(showClock, 1000);
});

function showClock() {
    var d = new Date();
    d.setMinutes(d.getMinutes() - d.getTimezoneOffset());
    $('#clock').html(d.toISOString().slice(0, 19).replace('T', ' '));
}

