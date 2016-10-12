// Code generated by protoc-gen-grpc-js-client
// source: ssh.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function sSHGet(p, conf) {
    path = '/ssh/v1beta1/ssh'
    return xhr(path, 'GET', conf, p);
}

var services = {
    sSH: {
        get: sSHGet
    }
};

module.exports = { ssh: { v1beta1: services } };
