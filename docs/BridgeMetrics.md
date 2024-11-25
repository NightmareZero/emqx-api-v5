# BridgeMetrics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dropped** | **int32** | Count of messages dropped. | [optional] [default to null]
**DroppedOther** | **int32** | Count of messages dropped due to other reasons. | [optional] [default to null]
**DroppedQueueFull** | **int32** | Count of messages dropped due to the queue is full. | [optional] [default to null]
**DroppedResourceNotFound** | **int32** | Count of messages dropped due to the resource is not found. | [optional] [default to null]
**DroppedResourceStopped** | **int32** | Count of messages dropped due to the resource is stopped. | [optional] [default to null]
**Matched** | **int32** | Count of this bridge is matched and queried. | [optional] [default to null]
**Queuing** | **int32** | Count of messages that are currently queuing. | [optional] [default to null]
**Retried** | **int32** | Times of retried. | [optional] [default to null]
**Failed** | **int32** | Count of messages that sent failed. | [optional] [default to null]
**Inflight** | **int32** | Count of messages that were sent asynchronously but ACKs are not yet received. | [optional] [default to null]
**Success** | **int32** | Count of messages that sent successfully. | [optional] [default to null]
**Rate** | **float64** | The rate of matched, times/second | [optional] [default to null]
**RateMax** | **float64** | The max rate of matched, times/second | [optional] [default to null]
**RateLast5m** | **float64** | The average rate of matched in the last 5 minutes, times/second | [optional] [default to null]
**Received** | **float64** | Count of messages that is received from the remote system. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

