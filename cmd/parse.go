package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/pkg/types"
)

func main() {
	str := "{\"event_id\":\"46b56e4e-1d5d-4353-b0af-276ff6714e62\",\"timestamp\":1559987446596,\"device\":{\"name\":\"led-light-instance-01\",\"state\":\"online\",\"last_online\":\"2019-06-08 17:50:46\"}}"
	twinUpdate := &types.DeviceTwinUpdate{}
	json.Unmarshal([]byte(str), twinUpdate)
	fmt.Println(twinUpdate)

}
