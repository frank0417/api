// Code generated by protoc-gen-grpc-js-client
// source: artifact.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function ArtifactsSearch(p, conf) {
	path = '/artifactory/v1beta1/search'
	return xhr(path, 'GET', conf, p);
}

function ArtifactsList(p, conf) {
	path = '/artifactory/v1beta1/artifacts/' + p['type']
	delete p['type']
	return xhr(path, 'GET', conf, p);
}

module.exports = {
  Artifacts: {
      Search: ArtifactsSearch,
      List: ArtifactsList
  }
};