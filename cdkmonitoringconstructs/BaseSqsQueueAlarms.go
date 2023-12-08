package cdkmonitoringconstructs


// Experimental.
type BaseSqsQueueAlarms struct {
	// Experimental.
	AddQueueMaxIncomingMessagesAlarm *map[string]*MaxIncomingMessagesCountThreshold `field:"optional" json:"addQueueMaxIncomingMessagesAlarm" yaml:"addQueueMaxIncomingMessagesAlarm"`
	// Experimental.
	AddQueueMaxMessageAgeAlarm *map[string]*MaxMessageAgeThreshold `field:"optional" json:"addQueueMaxMessageAgeAlarm" yaml:"addQueueMaxMessageAgeAlarm"`
	// Experimental.
	AddQueueMaxSizeAlarm *map[string]*MaxMessageCountThreshold `field:"optional" json:"addQueueMaxSizeAlarm" yaml:"addQueueMaxSizeAlarm"`
	// Experimental.
	AddQueueMaxTimeToDrainMessagesAlarm *map[string]*MaxTimeToDrainThreshold `field:"optional" json:"addQueueMaxTimeToDrainMessagesAlarm" yaml:"addQueueMaxTimeToDrainMessagesAlarm"`
	// Experimental.
	AddQueueMinIncomingMessagesAlarm *map[string]*MinIncomingMessagesCountThreshold `field:"optional" json:"addQueueMinIncomingMessagesAlarm" yaml:"addQueueMinIncomingMessagesAlarm"`
	// Experimental.
	AddQueueMinSizeAlarm *map[string]*MinMessageCountThreshold `field:"optional" json:"addQueueMinSizeAlarm" yaml:"addQueueMinSizeAlarm"`
}

