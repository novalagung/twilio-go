/*
 * Twilio - Bulkexports
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.22.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// BulkexportsV1Job struct for BulkexportsV1Job
type BulkexportsV1Job struct {
	// The details of a job state which is an object that contains a `status` string, a day count integer, and list of days in the job
	Details *map[string]interface{} `json:"details,omitempty"`
	// The optional email to send the completion notification to
	Email *string `json:"email,omitempty"`
	// The end time for the export specified when creating the job
	EndDay *string `json:"end_day,omitempty"`
	// this is the time estimated until your job is complete. This is calculated each time you request the job list. The time is calculated based on the current rate of job completion (which may vary) and your job queue position
	EstimatedCompletionTime *string `json:"estimated_completion_time,omitempty"`
	// The friendly name specified when creating the job
	FriendlyName *string `json:"friendly_name,omitempty"`
	// This is the job position from the 1st in line. Your queue position will never increase. As jobs ahead of yours in the queue are processed, the queue position number will decrease
	JobQueuePosition *string `json:"job_queue_position,omitempty"`
	// The job_sid returned when the export was created
	JobSid *string `json:"job_sid,omitempty"`
	// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType *string `json:"resource_type,omitempty"`
	// The start time for the export specified when creating the job
	StartDay *string `json:"start_day,omitempty"`
	Url      *string `json:"url,omitempty"`
	// This is the method used to call the webhook
	WebhookMethod *string `json:"webhook_method,omitempty"`
	// The optional webhook url called on completion
	WebhookUrl *string `json:"webhook_url,omitempty"`
}
