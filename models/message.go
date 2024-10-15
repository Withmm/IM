package models

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	FromId   int64  // sender
	TargetId int64  // receiver
	Type     int    // msgtype(public char; private chat)
	Media    int    // msg(file, text; ...)
	Content  string // content
	Pic      string
	Url      string
	Desc     string
	Amount   int
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets mapset.Set[int64]
}

// 映射关系
var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁
var rwLocker sync.RWMutex

// chat : senderId receiverId mestype mescontent sendtype
func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	userId, _ := strconv.ParseInt(query.Get("userId"), 10, 64)
	fmt.Println("userId", userId)
	//msgType := query.Get("type")
	//targetId := query.Get("targetId")
	//context := query.Get("context")
	isvalid := true
	conn, err := (&websocket.Upgrader{
		// token 校验
		CheckOrigin: func(r *http.Request) bool {
			return isvalid
		},
	}).Upgrade(writer, request, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	// get conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: mapset.NewSet[int64](),
	}

	// user contact

	// userid <-> node
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	// send
	go sendProc(node)
	// receive
	go receiveProc(node)

	sendMsg(userId, []byte("welcome to IM!"))
}

func sendProc(node *Node) {
	for {
		// fmt.Println("[debug] sendProc")
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func receiveProc(node *Node) {
	for {
		// fmt.Println("[debug] receiveProc")
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("[ws] <<<<<< ", data)
	}
}

var udpsendChan chan []byte = make(chan []byte, 1024)

func broadMsg(data []byte) {
	udpsendChan <- data
}

func init() {
	go udpSendProc()
	go udpReceiveProc()
}

// complete udp
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()
	for {
		select {
		case data := <-udpsendChan:
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func udpReceiveProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()
	for {
		var buf [512]byte
		n, err := con.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		disPatch(buf[0:n])
	}
}

// 后端调度的逻辑
func disPatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1: // 私信
		sendMsg(msg.TargetId, data)
	case 2: // 群发
		sendGroupMsg()
	case 3: // 广播
		sendAllMsg()
	case 4:
	}
}

// 发送消息额给个人， userId：发送消息的对象 msg：消息
func sendMsg(userId int64, msg []byte) {
	rwLocker.RLock()
	node, ok := clientMap[userId]
	rwLocker.RUnlock()

	if ok {
		node.DataQueue <- msg
	}

}

func sendGroupMsg() {

}

func sendAllMsg() {

}
