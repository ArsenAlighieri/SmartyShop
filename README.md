# SmartyShop

![SmartyShop Logo](frontend/src/assets/images/logo.png) <!-- Logunuzun yolu, eğer varsa -->

SmartyShop, kullanıcıların çeşitli e-ticaret sitelerinden ürünleri karşılaştırmasına ve Gemini API'si aracılığıyla yapay zeka destekli içgörüler elde etmesine olanak tanıyan akıllı bir alışveriş asistanıdır.

## Proje Genel Bakışı

Bu proje, Go ile yazılmış bir arka uç (backend) ve React ile geliştirilmiş bir ön uç (frontend) içeren bir web uygulamasıdır. Amacı, kullanıcılara bilinçli alışveriş kararları vermelerinde yardımcı olmaktır.

### Temel Özellikler:

*   **Ürün Kazıma (Scraping):** Trendyol, Teknosa, MediaMarkt ve Amazon gibi popüler e-ticaret sitelerinden ürün bilgilerini (ad, fiyat, puan, yorum sayısı vb.) toplar.
*   **Ürün Karşılaştırma:** Farklı satıcılardan gelen ürünlerin fiyatlarını ve özelliklerini karşılaştırma imkanı sunar.
*   **Yapay Zeka Destekli İçgörüler:** Gemini API'sini kullanarak ürünler hakkında akıllı özetler ve analizler sağlar.
*   **Modern Kullanıcı Arayüzü:** Temiz, modern ve koyu temalı bir arayüz ile sezgisel bir kullanıcı deneyimi sunar.

## Teknolojiler

### Backend (Go)

*   **Gin Web Framework:** Hızlı ve hafif bir web çerçevesi.
*   **Colly:** Web kazıma işlemleri için güçlü bir kütüphane.
*   **Go Modules:** Bağımlılık yönetimi için.
*   **Godotenv:** Ortam değişkenlerini `.env` dosyasından yüklemek için.

### Frontend (React)

*   **React.js:** Kullanıcı arayüzü oluşturmak için JavaScript kütüphanesi.
*   **CSS:** Modern ve duyarlı tasarım için.

## Kurulum ve Çalıştırma

Projeyi yerel makinenizde çalıştırmak için aşağıdaki adımları izleyin.

### Önkoşullar

*   [Go](https://golang.org/doc/install) (Backend için)
*   [Node.js](https://nodejs.org/en/download/) ve [npm](https://www.npmjs.com/get-npm) (Frontend için)

### 1. Backend Kurulumu

1.  Proje kök dizinine gidin:
    ```bash
    cd SmartyShop
    ```
2.  `backend` dizinine gidin:
    ```bash
    cd backend
    ```
3.  Bir `.env` dosyası oluşturun ve Gemini API anahtarınızı ekleyin:
    ```
    GEMINI_API_KEY=your_gemini_api_key_here
    ```
    *   Gemini API anahtarınızı [Google AI Studio](https://aistudio.google.com/app/apikey) adresinden alabilirsiniz.

4.  Backend sunucusunu çalıştırın:
    ```bash
    go run cmd/main.go
    ```
    Sunucu `http://localhost:8080` adresinde çalışmaya başlayacaktır.

### 2. Frontend Kurulumu

1.  Yeni bir terminal açın ve proje kök dizinine gidin:
    ```bash
    cd SmartyShop
    ```
2.  `frontend` dizinine gidin:
    ```bash
    cd frontend
    ```
3.  Gerekli bağımlılıkları yükleyin:
    ```bash
    npm install
    ```
4.  Frontend uygulamasını başlatın:
    ```bash
    npm start
    ```
    Uygulama genellikle `http://localhost:3000` adresinde açılacaktır.

## API Uç Noktaları

Backend, aşağıdaki API uç noktalarını sunar:

*   `GET /products?site=<site>&query=<query>`: Belirtilen siteden (trendyol, teknosa, mediamarkt, amazon) verilen sorguya göre ürünleri kazır ve döndürür.
    *   Örnek: `http://localhost:8080/products?site=trendyol&query=laptop`
*   `GET /products/top10?site=<site>&query=<query>`: Önbelleğe alınmış sonuçlardan en iyi 10 puanlı ürünü döndürür.
*   `POST /gemini/query`: Analiz için bir sorgu ve ürün listesini Gemini API'sine gönderir ve yapay zeka içgörülerini döndürür.
    *   İstek Gövdesi (JSON):
        ```json
        {
            "query": "ürünleri karşılaştır",
            "products": [
                { /* Ürün 1 */ },
                { /* Ürün 2 */ }
            ]
        }
        ```

## Geliştirme Kuralları

*   **Bağımlılık Yönetimi:** Go modülleri (backend) ve npm (frontend) kullanılır.
*   **Kod Yapısı:** Backend kodu `api`, `scrapers`, `gemini`, `internal`, `config` gibi paketlere ayrılmıştır.
*   **Hata Yönetimi:** Go'nun standart hata yönetimi mekanizmaları kullanılır.
*   **Loglama:** Standart `log` paketi kullanılır.

## Katkıda Bulunma

Katkılarınızı memnuniyetle karşılarız! Lütfen bir pull request göndermeden önce mevcut kod stilini ve kurallarını takip edin.

## Lisans

Bu proje açık kaynaklıdır ve [MIT Lisansı](LICENSE) altında lisanslanmıştır. <!-- Eğer bir lisans dosyanız varsa, burayı güncelleyin -->
