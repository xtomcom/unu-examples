#!/usr/bin/env node

const unu = require('./unu');

const options = {
    url: 'https://xtom.com',
    title: 'xTom',
    keyword: 'xTom'
};

unu(options)
    .then(response => {
        console.log(response);
    })
    .catch(err => {
        console.log(err);
    });