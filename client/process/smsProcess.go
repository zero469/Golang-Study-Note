package process

import (
	"encoding/json"
	"fmt"
	"go_code/chapter18/project3/client/model"
	"go_code/chapter18/project3/client/utils"
	"go_code/chapter18/project3/common/message"
)

type SmsProcess struct {
}

func (smsp *SmsProcess) SendGroupMes(content string) (err error) {
	var mes message.Message
	mes.Type = message.SmsMesType

	var smsMes message.SmsMes
	smsMes.SenderID = model.CurUser.UserID
	smsMes.Content = content

	data, err := json.Marshal(smsMes)
	if err != nil {
		return fmt.Errorf("sendGroupMes failed : %v", err)
	}

	mes.Data = string(data)

	data, err = json.Marshal(mes)
	if err != nil {
		return fmt.Errorf("sendGroupMes failed : %v", err)
	}

	tf := utils.Transfer{
		Conn: model.CurUser.Conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		return fmt.Errorf("sendGroupMes failed : %v", err)
	}
	return nil
}