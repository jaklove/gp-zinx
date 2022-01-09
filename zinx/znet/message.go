package znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

//创建一个Message消息包
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

//获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

func (msg *Message) GetData() []byte {
	return msg.Data
}

func (msg *Message) SetMsgId(id uint32) {
	msg.Id = id
}

func (msg *Message) SetData(data []byte) {
	msg.Data = data
}

func (msg *Message) SetDataLen(datalen uint32) {
	msg.DataLen = datalen
}
