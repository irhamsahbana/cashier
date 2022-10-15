package helper

import customtype "lucy/cashier/lib/custom_type"

func AddMessage(key string, value string, msg customtype.Message) customtype.Message {
	if _, ok := msg[key]; ok {
		msg[key] = append(msg[key], value)
	} else {
		msg[key] = []string{value}
	}

	return msg
}
