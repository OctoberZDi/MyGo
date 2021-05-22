package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"time"
)

// broker
var broker = "tcp://localhost:1883"

// clientId
var clientId = "testClient_Go"

// topic
var subTopic = "goTest"
var pubTopic = "goTest2"

func main() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)

	opts := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientId)

	// SetKeepAlive将设置客户端向代理发送PING请求之前应等待的时间（以秒为单位）。 这将使客户端知道与服务器的连接尚未丢失
	opts.SetKeepAlive(60 * time.Second)

	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(f)
	// SetPingTimeout将设置在向代理发送PING请求之后客户端确定连接丢失之前等待的时间（以秒为单位）。 默认值为10秒
	opts.SetPingTimeout(1 * time.Second)

	// NewClient将使用提供的ClientOptions中指定的所有选项创建一个MQTT v3.1.1客户端。
	// 客户端必须先调用Connect方法，然后才能使用它。 这是为了确保在实际准备好应用程序之前创建资源（例如网络连接）
	client := mqtt.NewClient(opts)

	// 创建连接
	connectToken := client.Connect()
	if connectToken.Wait() && connectToken.Error() != nil {
		panic(connectToken.Error())
	}

	// 订阅主题
	subToken := client.Subscribe(subTopic, 0, f)
	if subToken.Wait() && subToken.Error() != nil {
		fmt.Println(subToken.Error())
		os.Exit(1)
	}

	// 发布消息
	pubToken := client.Publish(pubTopic, 0, false, "Hello world ,hello go...")
	pubToken.Wait()
	time.Sleep(6 * time.Second)

	// 取消订阅
	if unsubToken := client.Unsubscribe(pubTopic); unsubToken.Wait() && unsubToken.Error() != nil {
		fmt.Println(unsubToken.Error())
		os.Exit(1)
	}

	// 断开连接
	client.Disconnect(250)
	time.Sleep(1 * time.Second)
}

var f mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("TOPIC****: %s\n", message.Topic())
	fmt.Printf("MSG****: %s\n", message.Payload())
}
