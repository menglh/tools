package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//url 格式 amqp://账号:密码@rabbitmq服务器地址:端口号/Virtual Host
const MqUrl = "amqp://guest:guest@192.168.13.88:5672/"

type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string //队列名称
	Exchange  string //交换机
	Key       string //key
	MqUrl     string //链接信息
}

//创建RabbitMQ结构体实例
func newRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, MqUrl: MqUrl}
	var err error
	//创建rabbitmq连接

	//通过amqp.Dial()方法去链接rabbitmq服务端
	rabbitmq.conn, err = amqp.Dial(rabbitmq.MqUrl)

	//调用我们自定义的failOnErr()方法去处理异常错误信息
	rabbitmq.failOnErr(err, "创建连接错误!")

	//链接上rabbitmq之后通过rabbitmq.conn.Channel()去设置channel信道
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败!")
	return rabbitmq
}

//断开channel和connection
func (r *RabbitMQ) Destory() {
	//关闭信道资源
	r.channel.Close()

	//关闭链接资源
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

//简单模式step：1.创建简单模式下的rabbitmq实例
/**
simple模式下交换机为空   因为会默认使用rabbitmq默认的default交换机而不是真的没有
bindkey绑定建key也是为空的

特别注意：simple模式是最简单的rabbitmq的一种模式 他只需要传递queue队列名称过去即可
exchange交换机会默认使用default交换机  绑定建key的会不必要传
*/

func NewSimple(queueName string) *RabbitMQ {
	return newRabbitMQ(queueName, "", "")
}

//简单模式step:2.简单模式下生产者
func (r *RabbitMQ) PublishSimple(message string) {
	//1. 申请队列,如果队列不存在，则会自动创建，如果队列存在则跳过创建直接使用  这样的好处保障队列存在，消息能发送到队列当中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
		false,
		//是否为自动删除  意思是最后一个消费者断开链接以后是否将消息从队列当中删除  默认设置为false不自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞 发送消息以后是否要等待消费者的响应 消费了下一个才进来 就跟golang里面的无缓冲channle一个道理 默认为非阻塞即可设置为false
		false,
		//其他的属性，没有则直接诶传入空即可 nil  nil,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//2.发送消息到队列当中
	err = r.channel.Publish(
		//交换机 simple模式下默认为空 我们在上边已经赋值为空了  虽然为空 但其实也是在用的rabbitmq当中的default交换机运行
		r.Exchange,
		//队列的名称
		r.QueueName,
		//如果为true 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返还给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者则会把消息返还给发送者
		false,
		//要发送的消息
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

	if err != nil {
		fmt.Println(err)
	}
}

//简单模式step:3.简单模式下消费者代码
func (r *RabbitMQ) ReceiveSimple() {
	//1.申请队列,如果队列不存在，则会自动创建，如果队列存在则跳过创建直接使用  这样的好处保障队列存在，消息能发送到队列当中
	_, err := r.channel.QueueDeclare(
		//队列名称
		r.QueueName,
		//进入的消息是否持久化 进入队列如果不消费那么消息就在队列里面 如果重启服务器那么这个消息就没啦 通常设置为false
		true,
		//是否为自动删除  意思是最后一个消费者断开链接以后是否将消息从队列当中删除  默认设置为false不自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞 发送消息以后是否要等待消费者的响应 消费了下一个才进来 就跟golang里面的无缓冲channle一个道理 默认为非阻塞即可设置为false
		false,
		nil, //其他的属性，没有则直接诶传入空即可 nil  nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//2.接收消息
	//建立了链接 就跟socket一样 一直在监听 从未被终止  这也就保证了下边的子协程当中程序的无线循环的成立
	msgs, err := r.channel.Consume(
		//队列名称
		r.QueueName,
		//用来区分多个消费者
		"",
		//是否自动应答 意思就是收到一个消息已经被消费者消费完了是否主动告诉rabbitmq服务器我已经消费完了你可以去删除这个消息啦 默认是true
		true,
		//是否具有排他性
		false,
		//如果设置为true表示不能将同一个connection中发送的消息传递给同个connectio中的消费者
		false,
		//队列消费是否阻塞 fase表示是阻塞 true表示是不阻塞
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//3.消费消息
	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		//子协程不会结束  因为msgs有监听 会不断的有值进来 就算没值也在监听  就跟socket服务一样一直在监听从未被中断！
		for d := range msgs {
			//实现我们要处理的逻辑函数
			log.Printf("Recieved a message : %s", d.Body)
			//fmt.Println(d.Body)
		}
	}()
	log.Printf("[*] Waiting for message,To exit press CTRL+C")
	//最后我们来阻塞一下 这样主程序就不会死掉
	<-forever
}

//订阅模式step1:创建rabbitmq实例
func NewSubscribe(exchangeName string) *RabbitMQ {
	return newRabbitMQ("", exchangeName, "")
}

//订阅模式step2:生产者
func (r *RabbitMQ) PublishSubscribe(message string) {
	//1.尝试创建交换机exchange 如果交换机存在就不用管他，如果不存在则会创建交换机
	err := r.channel.ExchangeDeclare(
		//交换机名称
		r.Exchange,
		//广播类型  订阅模式下我们需要将类型设置为广播类型
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an excha "+"nge")

	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		"",
		//如果为true 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返还给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message), //发送的内容一定要转换成字节的形式
		})
}

//订阅模式step3:消费者
func (r *RabbitMQ) ReceiveSubscribe() {
	//1.尝试创建交换机exchange 如果交换机存在就不用管他，如果不存在则会创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an excha "+"nge")
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		//随机生产队列名称 这个地方一定要留空
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	//3.绑定队列到exchange中去
	err = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这里的key要为空
		"",
		r.Exchange,
		false,
		nil,
	)
	//4.消费代码
	//4.1接收队列消息
	message, err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//4.2真正开始消费消息
	forever := make(chan bool)
	go func() {
		for d := range message {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	fmt.Println("退出请按 ctrl+c")
	<-forever
}

//路由模式step1:创建RabbitMQ实例
func NewRouting(exchangeName string, routingKey string) *RabbitMQ {
	return newRabbitMQ("", exchangeName, routingKey)
}

//路由模式step2:发送消息
func (r *RabbitMQ) PublishRouting(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//类型  路由模式下我们需要将类型设置为direct这个和在订阅模式下是不一样的
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an excha "+"nge")
	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		//除了设置交换机这也要设置绑定的key值
		r.Key,
		//如果为true 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返还给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message), //发送的内容一定要转换成字节的形式
		})
}

//路由模式step3：消费者
func (r *RabbitMQ) ReceiveRouting() {
	//1.尝试创建交换机exchange 如果交换机存在就不用管他，如果不存在则会创建交换机
	err := r.channel.ExchangeDeclare(
		//交换机名称
		r.Exchange,
		//类型  路由模式下我们需要将类型设置为direct类型
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an excha "+"nge")
	//2.试探性创建队列，这里注意队列名称不要写哦
	q, err := r.channel.QueueDeclare(
		//随机生产队列名称 这个地方一定要留空
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")
	//3.绑定队列到exchange中去
	err = r.channel.QueueBind(
		q.Name, //队列的名称  通过key去找绑定好的队列
		//在路由模式下，这里的key要填写
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	//4.消费代码
	//4.1接收队列消息
	message, err := r.channel.Consume(
		q.Name,
		//用来区分多个消费者
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	//4.2真正开始消费消息
	forever := make(chan bool)
	go func() {
		for d := range message {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	fmt.Println("退出请按 ctrl+c")
	<-forever
}

//topic主题模式step1:创建RabbitMQ实例
func NewTopic(exchange string, routingkey string) *RabbitMQ {
	return newRabbitMQ("", exchange, routingkey)
}

//topic主题模式step2:发送消息
func (r *RabbitMQ) PublishTopic(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//类型 topic主题模式下我们需要将类型设置为topic
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an excha "+"nge")
	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		//除了设置交换机这也要设置绑定的key值
		r.Key,
		//如果为true 会根据exchange类型和routkey规则，如果无法找到符合条件的队列那么会把发送的消息返还给发送者
		false,
		//如果为true,当exchange发送消息到队列后发现队列上没有绑定消费者则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//topic主题模式step2:消费者
//要注意key 规则
//其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
//匹配 huxiaobai.* 表示匹配 huxiaobai.hello 但是huxiaobai.one.two 需要用huxiaobai.# 才能匹配到
func (r *RabbitMQ) ReceiveTopic() {
	//1.尝试创建交换机exchange 如果交换机存在就不用管他，如果不存在则会创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		//类型 topic主题模式下我们需要将类型设置为topic
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an excha "+"nge")
	//2.试探性创建队列，这里注意队列名称不要写哦
	q, err := r.channel.QueueDeclare(
		//随机生产队列名称 这个地方一定要留空
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")
	//3.绑定队列到exchange中去
	err = r.channel.QueueBind(
		q.Name, //队列的名称  通过key去找绑定好的队列
		//在路由模式下，这里的key要填写
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	//4.消费代码
	message, err := r.channel.Consume(
		//队列名称
		q.Name,
		//用来区分多个消费者
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range message {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	fmt.Println("退出请按 ctrl+c")
	<-forever
}
