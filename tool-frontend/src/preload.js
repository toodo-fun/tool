global.electron = require('electron')

const openDefaultBrowser = function (url) {
    var exec = require('child_process').exec;
    switch (process.platform) {
        case "darwin":
            exec('open ' + url);
            break;
        case "win32":
            exec('start ' + url);
            break;
        default:
            exec('xdg-open', [url]);
    }
}

window.openDefaultBrowser = openDefaultBrowser