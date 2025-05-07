package dto

import "task-tracker/internal/model"

type LogListItem struct {
	IPAddress string `json:"ip_address"`
	CreatedAt string `json:"created_at"`
}

func ToLogListItem(l model.Log) LogListItem {
	return LogListItem{
		IPAddress: l.IPAddress,
		CreatedAt: l.CreatedAt.Format("2006-01-02 15:04:05.999"),
	}
}

func ToLogList(t []model.Log) []LogListItem {
	var list []LogListItem
	for _, task := range t {
		list = append(list, ToLogListItem(task))
	}
	return list
}
