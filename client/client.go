package client

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ananagame/rich-go/ipc"
)

type Handshake struct {
	V        string `json:"v"`
	ClientId string `json:"client_id"`
}

type Frame struct {
	Cmd   string `json:"cmd"`
	Args  Args   `json:"args"`
	Nonce string `json:"nonce"`
}

type Args struct {
	Pid      int       `json:"pid"`
	Activity *Activity `json:"activity"`
}

type Activity struct {
	Details    string     `json:"details"`
	State      string     `json:"state"`
	Timestamps Timestamps `json:"timestamps"`
	Assets     Assets     `json:"assets"`
}

type Timestamps struct {
	Start int64 `json:"start"`
	// End   int64 `json:"end"`
}

type Assets struct {
	LargeImage string `json:"large_image"`
	LargeText  string `json:"large_text"`
	SmallImage string `json:"small_image"`
	SmallText  string `json:"small_text"`
}

func Login(clientid string) {
	payload, err := json.Marshal(Handshake{"1", clientid})
	if err != nil {
		panic(err)
	}

	ipc.OpenSocket()
	fmt.Println(ipc.Send(0, string(payload)))
}

func SetActivity(activity *Activity) {
	payload, err := json.Marshal(Frame{
		"SET_ACTIVITY",
		Args{
			os.Getpid(),
			activity,
		},
		getNonce(),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(payload))

	fmt.Println(ipc.Send(1, string(payload)))
}

func getNonce() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	buf[6] = (buf[6] & 0x0f) | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}
