package models

type Trigger struct {
	ID      int64       `db:"id"`
	Type    TriggerType `db:"type"`
	Trigger string      `db:"trigger"`
	Value   string      `db:"value"`
}

type TriggerType int8

const (
	TriggerTypeText TriggerType = iota
	TriggerTypeMedia
	TriggerTypeWarn
	TriggerTypeMute
	TriggerTypeBan
)
