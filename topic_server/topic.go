package topicserver

type topic struct {
	Title     string
	Content   string
	Creator   string
	Create_ts uint32
	Id        uint64
}

var map_topic = make(map[uint64]*topic)
var uniq_id = uint64(0)
