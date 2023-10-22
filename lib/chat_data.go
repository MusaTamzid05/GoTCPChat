package lib


type ChatData struct {
    Name string
    Message string
}

func MakeChatData(name, message string) ChatData {
    return ChatData{Name: name, Message: message}

}

func (c ChatData) String() string {
    return c.Name + " : " + c.Message 
}
