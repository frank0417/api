// Code generated by protoc-gen-grpc-js-client
// source: loadbalancer.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function loadBalancersList(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/loadbalancers'
    delete p['cluster']
    return xhr(path, 'GET', conf, p);
}

function loadBalancersDescribe(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/loadbalancers/' + p['name']
    delete p['cluster']
    delete p['name']
    return xhr(path, 'GET', conf, p);
}

function loadBalancersCreate(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/loadbalancers'
    delete p['cluster']
    return xhr(path, 'POST', conf, null, p);
}

function loadBalancersUpdate(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/loadbalancers/' + p['name']
    delete p['cluster']
    delete p['name']
    return xhr(path, 'PUT', conf, null, p);
}

function loadBalancersDelete(p, conf) {
    path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/loadbalancers/' + p['name']
    delete p['cluster']
    delete p['name']
    return xhr(path, 'DELETE', conf, p);
}

var services = {
    loadBalancers: {
        list: loadBalancersList,
        describe: loadBalancersDescribe,
        create: loadBalancersCreate,
        update: loadBalancersUpdate,
        delete: loadBalancersDelete
    }
};

module.exports = { loadbalancer: { v1beta1: services } };
