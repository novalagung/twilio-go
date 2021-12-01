/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListRoomParticipantSubscribedTrackResponse struct for ListRoomParticipantSubscribedTrackResponse
type ListRoomParticipantSubscribedTrackResponse struct {
	Meta             ListCompositionHookResponseMeta         `json:"meta,omitempty"`
	SubscribedTracks []VideoV1RoomParticipantSubscribedTrack `json:"subscribed_tracks,omitempty"`
}
