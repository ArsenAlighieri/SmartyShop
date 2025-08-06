# SmartyShop: Akıllı Alışveriş Asistanı

![SmartyShop Logo](frontend/src/assets/images/logo.png)

**SmartyShop**, kullanıcıların çeşitli e-ticaret sitelerinden ürünleri anında karşılaştırmasına ve Google'ın güçlü Gemini API'si aracılığıyla yapay zeka destekli derinlemesine ürün analizleri elde etmesine olanak tanıyan modern bir web uygulamasıdır.

[![React](https://img.shields.io/badge/Frontend-React-blue?style=for-the-badge&logo=react)](https://reactjs.org/) [![Go](https://img.shields.io/badge/Backend-Go-blue?style=for-the-badge&logo=go)](https://golang.org/) [![Docker](https://img.shields.io/badge/Container-Docker-blue?style=for-the-badge&logo=docker)](https://www.docker.com/)

---

## ✨ Proje Demosu

*Uygulamanın genel görünümü ve çalışma mantığı.*

![SmartyShop Arayüzü](frontend/src/assets/images/Adsız.png)

---

## 🚀 Temel Özellikler

*   **🛍️ Çoklu Site Desteği:** Trendyol, Teknosa, MediaMarkt ve Amazon gibi popüler e-ticaret sitelerinden canlı ürün verilerini çeker.
*   **📊 Anlık Fiyat Karşılaştırma:** Aynı ürünün farklı satıcılardaki fiyatlarını ve özelliklerini tek bir ekranda karşılaştırır.
*   **🧠 Yapay Zeka Destekli Analiz:** Gemini API'sini kullanarak ürünlerin avantajları, dezavantajları ve kullanıcı yorumları hakkında akıllı özetler ve analizler sunar.
*   **🏆 En İyi Ürünleri Keşfet:** Belirli bir kategoride en yüksek puana sahip ilk 10 ürünü listeler.
*   **🌑 Modern ve Sezgisel Arayüz:** Temiz, modern ve koyu temalı bir arayüz ile kullanıcı dostu bir deneyim sağlar.

---

## 🛠️ Teknoloji Mimarisi

### Backend (Go)

*   **Gin Web Framework:** Yüksek performanslı ve hafif bir web çerçevesi.
*   **Colly:** Web kazıma (scraping) işlemleri için güçlü ve esnek bir kütüphane.
*   **Go Modules:** Bağımlılık yönetimi.
*   **Godotenv:** Ortam değişkenlerini `.env` dosyasından güvenli bir şekilde yüklemek için.

### Frontend (React)

*   **React.js:** Dinamik ve etkileşimli kullanıcı arayüzleri oluşturmak için.
*   **CSS Modules:** Stil çakışmalarını önlemek ve bileşen bazlı stil yönetimi için.
*   **Axios:** Backend API'si ile iletişim kurmak için modern bir HTTP istemcisi.

---

## 🐳 Docker ile Hızlı Başlangıç

Projeyi en hızlı şekilde ayağa kaldırmak için Docker ve Docker Compose kullanabilirsiniz.

1.  Proje kök dizininde `backend/` klasörü içine bir `.env` dosyası oluşturun ve Gemini API anahtarınızı ekleyin:
    ```
    GEMINI_API_KEY=your_gemini_api_key_here
    ```
    *   Gemini API anahtarınızı [Google AI Studio](https://aistudio.google.com/app/apikey) adresinden alabilirsiniz.

2.  Proje kök dizininde aşağıdaki komutu çalıştırın:
    ```bash
    docker-compose up --build
    ```
    Bu komut, hem backend hem de frontend için gerekli imajları oluşturacak ve container'ları başlatacaktır.
    *   **Backend:** `http://localhost:8080`
    *   **Frontend:** `http://localhost:3000`

---

## 💻 Yerel Kurulum (Docker Olmadan)

Projeyi yerel makinenizde manuel olarak çalıştırmak için aşağıdaki adımları izleyin.

### Önkoşullar

*   [Go (1.18+)](https://golang.org/doc/install)
*   [Node.js (16+)](https://nodejs.org/en/download/) ve npm

### 1. Backend Kurulumu

1.  `backend` dizinine gidin: `cd backend`
2.  Bir `.env` dosyası oluşturun ve `GEMINI_API_KEY`'inizi ekleyin (yukarıdaki gibi).
3.  Gerekli Go modüllerini indirin:
    ```bash
    go mod tidy
    ```
4.  Backend sunucusunu çalıştırın:
    ```bash
    go run cmd/main.go
    ```
    Sunucu `http://localhost:8080` adresinde çalışmaya başlayacaktır.

### 2. Frontend Kurulumu

1.  Yeni bir terminalde `frontend` dizinine gidin: `cd frontend`
2.  Gerekli bağımlılıkları yükleyin:
    ```bash
    npm install
    ```
3.  Frontend uygulamasını başlatın:
    ```bash
    npm start
    ```
    Uygulama `http://localhost:3000` adresinde açılacaktır.

---

## 🔌 API Uç Noktaları (Endpoints)

Backend, aşağıdaki API uç noktalarını sunar:

| Metot | Uç Nokta                                    | Açıklama                                                              |
| :---- | :------------------------------------------ | :-------------------------------------------------------------------- |
| `GET` | `/products?site=<site>&query=<query>`       | Belirtilen siteden ürünleri kazır. Ör: `/products?site=trendyol&query=laptop` |
| `GET` | `/products/top10?site=<site>&query=<query>` | Önbellekten en iyi 10 puanlı ürünü döndürür.                          |
| `POST`| `/gemini/query`                             | Ürün listesini analiz için Gemini API'sine gönderir.                  |

**POST `/gemini/query` Örnek İstek Gövdesi:**
```json
{
    "query": "Bu iki laptop modelini oyun performansı açısından karşılaştır.",
    "products": [
        {
            "name": "Laptop Model A",
            "price": "35000 TL",
            "rating": "4.8",
            "reviews_count": "250"
        },
        {
            "name": "Laptop Model B",
            "price": "38000 TL",
            "rating": "4.9",
            "reviews_count": "180"
        }
    ]
}
```

---

## 🤝 Katkıda Bulunma

Katkılarınız projenin gelişimi için çok değerlidir! Lütfen bir pull request göndermeden önce mevcut kod stilini ve kurallarını inceleyin.

1.  Projeyi fork'layın.
2.  Yeni bir özellik dalı oluşturun (`git checkout -b feature/yeni-ozellik`).
3.  Değişikliklerinizi commit'leyin (`git commit -m 'Yeni özellik eklendi'`).
4.  Dalınızı push'layın (`git push origin feature/yeni-ozellik`).
5.  Bir Pull Request açın.

---

## 📄 Lisans

Bu proje [MIT Lisansı](LICENSE) altında lisanslanmıştır.