package main

import (
	"fmt"
)

func main() {
	sayHello := func() {
		fmt.Println("hello")
	}

	// `go sayHello()` ile `sayHello`'nun paralel çalışması sağlanıyor. Böylece
	// `sayHello` direkt çalışmak yerine `go`'nun runtine'nın sağladığı kuyruğa
	// ekleniyor. runtime'ın durumuna göre ileri -ve de tam olarak öngörülemeyen-
	// bir zamanda bu fonksiton çalıştırılacak. Bu arada ana akış da devam etmekte.
	// Bu ayrılmadan sonra eğer kod içerisinde açık olarak belirtilmezse ana akış
	// `goroutine`'inin tamamlanmasını beklemeyecek. Sonuç olarak -çok büyük ihtimalle-
	// `sayHello`'nun çalışması tamamlanmadan ana akış tamamlalanca ve fonksiyonun
	// sonucu görülmeden uygulama kapancak.
	go sayHello()
	fmt.Println("exiting app")
}
