package models

import "time"

type Discussion struct {
	id        int64
	topic     string
	content   string
	author    int
	createdAt time.Time
}
