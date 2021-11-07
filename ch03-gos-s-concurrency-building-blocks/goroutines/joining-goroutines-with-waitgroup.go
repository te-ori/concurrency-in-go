package main

import (
	"fmt"
	"sync"
)

func main() {
	// `WaitGroup` basitçe değeri atomik bir şekilde artan ya da azalan
	// bir sayaç tutmakta. Bu sayacın değeri `Add` metodu ile, verilen
	// argüman kadar artar. Yalnızca ilgili `WaitGroup` öğesinin `Done`
	// fonksiyonu çağrıldıkça **bir** azalır. `wg`'nin `Wait()` metodunun
	// çağrıldığı andan itibaren uygulamanın  akışı durur ve içindeki sayısın
	// 0 olmasını bekler. Ancak sayaç 0 olduğunda ana akışın devam etmesine
	// izin veirlir.
	var wg sync.WaitGroup

	sayHello := func() {
		// fonksiyonun çalışması tamamşandığında `wg`'nin `Done()` fonksiyonunun
		// çağrılması sağlanıyor. `Done()` metodu her çağrıldığında `wg'nin sayaç
		// değerini bir azaltacak
		defer wg.Done()
		fmt.Println("Hello")
	}

	// Uygulama içerisinde `wg.Add()` bir defa çağrılmış ve 1 değeri eklenmiş. Yani
	// `wg.Done()`'un bir defa çağrılması ana kaışın devamı için yeterli olcaktır.
	wg.Add(1)
	go sayHello()

	// Eğer `Wait()` metodu burada çağrılmış olsaydı `wg`'nin sayacının sıfır olması
	// burada beklecenkti. Yani ancak `goroutine` imizin çalışması tamamlandıktan ve
	// `defer` edilmiş `wg.Done()`'un sayacı bir azaltmasından sonra çalışmaya devam
	// edecekti. Böylece `"Hello"`, `"exiting app"`'ten önce yazdırılması granti
	// altına alınmış olacaktı
	// wg.Wait()

	fmt.Println("exiting app")

	// `WaitGroup`'un tamamlanması burada beklenirse henüz `goroutine`imiz çalışmasını
	// tamamlamadığı için ana akış buraya kadar gelecektir. Bu nednele de `"exiting
	// app"` yazısının ekrana yazdıırlması çok çok büyük bir ihtimalle `"Hello"`'dan
	// önce olacaktır.
	wg.Wait()
}
