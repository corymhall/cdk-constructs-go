package cdkmonitoringconstructs


// Experimental.
type BaseDlqAlarms struct {
	// Alarm on the number of messages added to a queue.
	//
	// Note that this corresponds with the NumberOfMessagesSent metric, which does not capture messages sent to the DLQ
	// as a result of a failed processing attempt.
	// See: https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html#sqs-dlq-number-of-messages
	//
	// Experimental.
	AddDeadLetterQueueMaxIncomingMessagesAlarm *map[string]*MaxIncomingMessagesCountThreshold `field:"optional" json:"addDeadLetterQueueMaxIncomingMessagesAlarm" yaml:"addDeadLetterQueueMaxIncomingMessagesAlarm"`
	// Experimental.
	AddDeadLetterQueueMaxMessageAgeAlarm *map[string]*MaxMessageAgeThreshold `field:"optional" json:"addDeadLetterQueueMaxMessageAgeAlarm" yaml:"addDeadLetterQueueMaxMessageAgeAlarm"`
	// Experimental.
	AddDeadLetterQueueMaxSizeAlarm *map[string]*MaxMessageCountThreshold `field:"optional" json:"addDeadLetterQueueMaxSizeAlarm" yaml:"addDeadLetterQueueMaxSizeAlarm"`
}

