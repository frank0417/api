// Code generated by protoc-gen-grpc-js-client
// source: build.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function buildsDescribe(p, conf) {
    path = '/ci/v1beta1/jobs/' + p['job_name'] + '/builds/' + p['number']
    delete p['job_name']
    delete p['number']
    return xhr(path, 'GET', conf, p);
}

function buildsList(p, conf) {
    path = '/ci/v1beta1/jobs/' + p['job_name'] + '/builds'
    delete p['job_name']
    return xhr(path, 'GET', conf, p);
}

var services = {
    builds: {
        describe: buildsDescribe,
        list: buildsList
    }
};

module.exports = { ci: { v1beta1: services } };
