// Пример: что возвращает listener.Accept() и какие данные даёт conn.
// Запуск: go run .   Затем в другом терминале: echo "привет" | nc localhost 8080
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Слушаем :8080. Подключитесь: nc localhost 8080")

	for {
		// Accept() возвращает два значения:
		//   conn — реализация net.Conn (чтение/запись с клиентом)
		//   err  — ошибка (например, если листенер закрыт)
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Accept:", err)
			continue
		}

		// Примеры получаемых значений из conn:
		fmt.Println("--- Новое подключение ---")
		fmt.Println("  Тип conn:     ", fmt.Sprintf("%T", conn))           // *net.TCPConn
		fmt.Println("  Адрес клиента:", conn.RemoteAddr().String())        // например 127.0.0.1:54321
		fmt.Println("  Наш адрес:    ", conn.LocalAddr().String())         // 127.0.0.1:8080
		fmt.Println("  RemoteAddr()  ", conn.RemoteAddr())                  // полный net.Addr
		fmt.Println("  LocalAddr()   ", conn.LocalAddr())

		// Закрываем соединение после показа (в реальном сервере здесь была бы go handle(conn))
		_ = conn.Close()
	}
}
