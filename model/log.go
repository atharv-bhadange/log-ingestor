package model

import "time"

// {
// 	"level": "error",
// 	"message": "Failed to connect to DB",
//     "resourceId": "server-1234",
// 	"timestamp": "2023-09-15T08:00:00Z",
// 	"traceId": "abc-xyz-123",
//     "spanId": "span-456",
//     "commit": "5e5342f",
//     "metadata": {
//         "parentResourceId": "server-0987"
//     }
// }

type Log struct {
	ID         uint      `json:"id" gorm:"column:id"`
	Level      string    `json:"level" gorm:"column:level"`
	Message    string    `json:"message" gorm:"column:message"`
	ResourceId string    `json:"resourceId" gorm:"column:resourceId"`
	Timestamp  time.Time `json:"timestamp" gorm:"column:timestamp"`
	TraceId    string    `json:"traceId" gorm:"column:traceId"`
	SpanId     string    `json:"spanId" gorm:"column:spanId"`
	Commit     string    `json:"commit" gorm:"column:commit"`
	Metadata   string    `json:"metadata" gorm:"column:metadata"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
