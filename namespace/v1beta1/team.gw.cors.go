// Code generated by protoc-gen-grpc-gateway-cors
// source: team.proto
// DO NOT EDIT!

/*
Package v1beta1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1beta1

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

// ExportTeamsCorsPatterns returns an array of grpc gatway mux patterns for
// Teams service to enable CORS.
func ExportTeamsCorsPatterns() []runtime.Pattern {
	return []runtime.Pattern{
		pattern_Teams_Create_0,
		pattern_Teams_Get_0,
		pattern_Teams_IsAvailable_0,
		pattern_Teams_Subscription_0,
	}
}
