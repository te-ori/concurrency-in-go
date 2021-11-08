# sync.Pool

* Eğer pool'da kullandığınız objeler homojen değilse, yani farklı `type`'larda ise *pooling* ile kazandığınız zamanı type checking ve type conversion'lar yüzünden kaybedeceksiniz. Bu nedenle objelerin homojen olmadığı durumlarda *pooling* istenilen performansı sağlamayabilir.
* `pool.Get()` ile havuzdan bir öğe aldığınızda bu öğenin durmu ile ilgili varsayımlarda bulunmayın. Örneğin pool'a `cnt` field'ı `1` olan bir öğe eklediğinizde ve `Get` ile bu öğeyi tekrar çağırdığınızda `cnt`'nin değerinin `1` olduğunu varsaymayın.
* Havuzdan aldığınız bir öğe ile işiniz bittiğinizde `Put` ile onu havuza iade etmeyi unutmayın. Aksi halde `Pool` kullanmanın bir avantajı olmaz.