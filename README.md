  <a href="https://www.python.org" target="_blank" rel="noreferrer"> 
        <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/python/python-original.svg" alt="python" width="40" height="40"/> 
    </a> 
    <br>
<h1>Python ile yazılmış olan geolocation service</h1>
Betik, IP adresine ait ülke, bölge, şehir, enlem ve boylam gibi konum bilgilerini sağlamak amacıyla ipstack API'sini kullanmaktadır.
Amaç

Betik, kullanıcıdan bir IP adresi alarak, bu IP adresine ait coğrafi konum bilgilerini otomatik olarak elde etmek ve kullanıcıya sunmak amacıyla tasarlanmıştır. Bu işlem, çeşitli nedenlerle IP adresi tabanlı konum tespiti yapması gereken kullanıcılar için kullanışlı olabilir.
Kullanılan Teknolojiler

    Python: Betiğin yazıldığı programlama dili.
    requests: Python'un HTTP kütüphanesi, API çağrılarını gerçekleştirmek için kullanılır.
    ipstack API: IP adresi tabanlı konum bilgisi sağlayan bir API hizmeti. 

Betiğin İşleyişi


    API Anahtarının Tanımlanması: ipstack API'sine erişim sağlamak için bir API anahtarı gereklidir. Kullanıcı, bu anahtarı betikte belirtilen yere eklemelidir.

    Kullanıcıdan IP Adresinin Alınması: Betik, kullanıcıdan konum bilgisi öğrenilmek istenen IP adresini girmesini ister.

    API Çağrısı: Kullanıcıdan alınan IP adresi ve API anahtarı kullanılarak, ipstack API'sine bir HTTP GET isteği gönderilir.

    Verilerin Alınması ve İşlenmesi:
        API'den gelen yanıt JSON formatında işlenir.
        HTTP yanıtının durumu kontrol edilir. Yanıt başarılıysa (status_code 200), IP adresine ait konum bilgileri JSON verisinden alınır.
        Eğer yanıt başarısızsa, hata bilgisi JSON verisinden çıkarılarak kullanıcıya bildirilir.

    Bilgilerin Kullanıcıya Sunulması: API yanıtı başarılı ise, elde edilen konum bilgileri kullanıcıya sunulur. Bilgiler arasında IP adresi, ülke, bölge, şehir, enlem ve boylam yer alır. Eğer bir hata oluşmuşsa, hata mesajı kullanıcıya iletilir.

Sonuç

Bu Python betiği, kullanıcıların bir IP adresine ait coğrafi konum bilgilerini hızlı ve kolay bir şekilde elde etmelerini sağlar. Betik, ipstack API'si aracılığıyla doğru ve güvenilir veriler sunar. Kullanıcıdan alınan IP adresine dayalı olarak ülke, bölge, şehir, enlem ve boylam bilgileri API yanıtı doğrultusunda kullanıcıya gösterilir. Hatalı durumlar için ise uygun hata mesajları iletilerek kullanıcı bilgilendirilir.
<br>
   <a href="https://golang.org/" target="_blank" rel="noreferrer"> 
        <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="golang" width="40" height="40"/> 
    </a>
<h1>Golang ile yazılmış olan geolocation service</h1>

Program, ipstack API'sini kullanarak bir IP adresine ait ülke, bölge, şehir, enlem ve boylam bilgilerini elde etmektedir.
Genel Bakış

Program, kullanıcıdan bir IP adresi girmesini istemekte ve ardından ipstack API'sini kullanarak bu IP adresine ait konum bilgilerini getirmektedir. Alınan bilgiler kullanıcıya gösterilmektedir. Programın temel bileşenleri ve işleyişi aşağıda detaylı olarak açıklanmıştır.
Programın Detayları

    Gerekli Kütüphanelerin Dahil Edilmesi
        encoding/json: JSON verilerini işlemek için kullanılır.
        fmt: Girdi ve çıktılar için formatlı I/O işlemleri sağlar.
        net/http: HTTP isteklerini yapmak için kullanılır.
        os: İşletim sistemi ile etkileşim kurmak için kullanılır (örneğin, hata durumunda programdan çıkmak).

    API Yanıtı İçin Struct Tanımlaması
        IpstackResponse adlı struct, ipstack API'sinden dönen JSON verilerini karşılamak için tanımlanmıştır.

    Konum Bilgisi Alma Fonksiyonu
        getLocation fonksiyonu, verilen bir IP adresi ve API anahtarı ile ipstack API'sine istek gönderir ve dönen yanıtı IpstackResponse struct'ına dönüştürür.
        Fonksiyon, HTTP isteği yapar, yanıtı JSON formatında okur ve IpstackResponse struct'ına unmarshale eder.
        Eğer bir hata oluşursa, hata mesajını döndürür.

    Ana Fonksiyon
        main fonksiyonu, kullanıcıdan bir IP adresi girmesini ister ve ardından getLocation fonksiyonunu çağırarak bu IP adresine ait konum bilgilerini alır.
        Alınan konum bilgileri kullanıcıya yazdırılır.
        Eğer bir hata oluşursa, hata mesajını yazdırır ve programdan çıkar.

Çalışma Şekli

    Kullanıcı programı çalıştırır ve IP adresini girer.
    Program, ipstack API'sine istek gönderir ve IP adresine ait coğrafi konum bilgilerini alır.
    Alınan bilgiler (IP adresi, ülke, bölge, şehir, enlem ve boylam) ekrana yazdırılır.
    Eğer API anahtarı geçersizse veya başka bir hata oluşursa, hata mesajı ekrana yazdırılır ve program sonlandırılır.

Sonuç


Bu program, kullanıcı tarafından girilen bir IP adresine ait coğrafi konum bilgilerini almayı ve bu bilgileri ekrana yazdırmayı amaçlamaktadır. Program, HTTP istekleri ve JSON işlemleri için Go'nun standart kütüphanelerini kullanmaktadır. IP geolocation hizmeti olarak ipstack API'si kullanılmıştır. Kullanıcıya basit ve anlaşılır bir arayüz sunar ve gerekli durumlarda hata mesajları ile bilgilendirir.











