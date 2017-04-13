const request = require('request');

module.exports = (options) => {
    if (typeof options !== 'object') {
        return Promise.reject('invalid options');
    }
    if (typeof options.url !== 'string') {
        return Promise.reject('invalid options.url');
    }
    if ('title' in options && typeof options.title !== 'string') {
        return Promise.reject('invalid options.title');
    }
    if ('keyword' in options && typeof options.keyword !== 'string') {
        return Promise.reject('invalid options.keyword');
    }
    if ('username' in options && typeof options.username !== 'string') {
        return Promise.reject('invalid options.username');
    }
    if ('password' in options && typeof options.password !== 'string') {
        return Promise.reject('invalid options.password');
    }
    let form = {
        action: 'shorturl',
        format: 'json',
        url: options.url
    };
    if (options.title) {
        form.title = options.title;
    }
    if (options.keyword) {
        form.keyword = options.keyword;
    }
    if (options.username && options.password) {
        form.username = options.username;
        form.password = options.password;
    }
    return new Promise((resolve, reject) => {
        const options = {
            url: 'https://u.nu/api.php',
            form: form
        };
        request.post(options, (err, httpResponse, body) => {
            if (err) {
                reject(err);
                return;
            }
            resolve(body);
        });
    });
};