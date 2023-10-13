package main

import (
	"io/fs"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

const (
	network string      = "tcp4"
	address string      = ":12345"
	flag    int         = os.O_RDONLY
	perm    fs.FileMode = 0777
)

func main() {
	//map с Go-поговорками
	proverbs := map[int]string{
		1:  "Don't communicate by sharing memory, share memory by communicating.",
		2:  "Concurrency is not parallelism.",
		3:  "Channels orchestrate; mutexes serialize.",
		4:  "The bigger the interface, the weaker the abstraction.",
		5:  "Make the zero value useful.",
		6:  "interface{} says nothing.",
		7:  "Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
		8:  "A little copying is better than a little dependency.",
		9:  "Syscall must always be guarded with build tags.",
		10: "Cgo must always be guarded with build tags.",
		11: "Cgo is not Go.",
		12: "With the unsafe package there are no guarantees.",
		13: "Clear is better than clever.",
		14: "Reflection is never clear.",
		15: "Errors are values.",
		16: "Don't just check errors, handle them gracefully.",
		17: "Design the architecture, name the components, document the details.",
		18: "Documentation is for users.",
		19: "Don't panic.",
	}
	//Сетевая служба по протоколу TCP
	srv, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Close()
	//запускаем бесконечный цикл, чтобы сервер не завершал работу
	for {
		conn, err := srv.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleFunc(conn, proverbs)
	}
}

// Вспомогательная функция для вывода клиенту Го-поговорок раз в 3 секунды.
func handleFunc(conn net.Conn, proverbs map[int]string) {
	defer conn.Close()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		r := rand.Intn(19) + 1
		val, ok := proverbs[r]
		if !ok {
			log.Printf("Запрошенный ключ %d находится вне существующих значений в map", r)
		}
		conn.Write([]byte(val + " "))
		time.Sleep(3 * time.Second)
	}
}
