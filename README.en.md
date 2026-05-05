# 🌊 Tide API Go

[![Continuous Integration](https://github.com/casimiroarruda/go-tide-table-api/actions/workflows/ci.yml/badge.svg)](https://github.com/casimiroarruda/go-tide-table-api/actions/workflows/ci.yml)
[![Deploy to Cloud Run](https://github.com/casimiroarruda/go-tide-table-api/actions/workflows/deploy.yml/badge.svg)](https://github.com/casimiroarruda/go-tide-table-api/actions/workflows/deploy.yml)
![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat&logo=go)
![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL%20%2B%20PostGIS-336791?style=flat&logo=postgresql)
![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat)
![Status](https://img.shields.io/badge/Status-Active-green.svg?style=flat)
![Docs](https://img.shields.io/badge/Docs-Swagger-00ADD8?style=flat)
![Test Coverage](https://img.shields.io/badge/Test%20Coverage-89.22%25-green.svg?style=flat)

🇧🇷 [Versão em Português](README.md)

🇬🇧 **English Version**

The **Tide API Go** is a high-performance service for querying tide tables, designed with a focus on geographic precision and cost efficiency. Built in **Go** and utilizing advanced geospatial processing via **PostGIS**, the API provides accurate tide data (high/low) based on real-time coordinates or location search.

## Motivation

The idea came from the need to have a tide API that was fast, accurate, and could be used anywhere in the world.

The Tide Table is a tool heavily used in practically all coastal areas, for water sports, fishing, tourism, and more. My family often relies on tide table information for our trips to the beautiful Boa Viagem beach in Recife, Pernambuco (Brazil).

Combining the useful with the pleasant: I developed this API to meet this demand and to share the discoveries and lessons learned during the process. It is also part of a college project. Additionally, it is a personal challenge, as it has been a long time since I single-handedly put an Open Source software into production using tools outside my comfort zone, such as the Go programming language.

## Features

* **Search by location:** Find tide data by location name
* **Search by coordinates:** Find tide data using geographic coordinates
* **Tide table:** Display the tide table by day

---

## 🚀 Technical Highlights

* **Geospatial Search (KNN):** Instant location of the nearest tide station via latitude/longitude using the `<->` operator and **GIST** indexes, ensuring performance even with large volumes of data.
* **Timezone-Aware Queries:** Filtering logic executed at the database engine level (PostgreSQL), respecting the local timezone of each station to define the "start and end of the day".
* **SARGable Queries:** Implementation of date filters that preserve the efficiency of database indexes, optimizing CPU consumption on Google Cloud Run.
* **Serverless Architecture:** Automated deployment via **GitHub Actions** to **Google Cloud Run**, with secrets injection via GCP Secret Manager.
* **Security & Precision:** Authentication via **JWT** and use of fixed-precision decimal types (`shopspring/decimal`) to avoid rounding errors in height calculations.

---

## 🛠️ Tech Stack

* **Language:** Go (Golang)
* **Database:** PostgreSQL + PostGIS
* **Infrastructure:** Google Cloud Platform (Cloud Run, Artifact Registry)
* **CI/CD:** GitHub Actions
* **Documentation:** ApiDog / OpenAPI 3.0

---

## 📦 How to Run Locally

### Prerequisites
* Docker & Docker Compose
* Go 1.26+

### Step by Step
1.  **Clone the repository:**
    ```bash
    git clone https://github.com/casimiroarruda/go-tide-table-api.git
    cd tide-api-go
    ```
2.  **Environment Setup:**
    Create a `.env` file based on the example:
    ```env
    DATABASE_URL=postgres://user:pass@localhost:5432/tide_db?sslmode=disable
    JWT_SECRET=your_secret_key_here
    ```
3.  **Start the Environment:**
    ```bash
    docker-compose up -d
    ```
    *The database will **not** run the migrations automatically, including enabling PostGIS. You must run them manually. Follow the instructions in the `docs/DEPLOY.md` file.*

---

## 📡 API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/api/auth/token` | JWT token generation. |
| `GET` | `/api/location` | Lists locations or search by `?name=`. |
| `GET` | `/api/location?lat=&lon=` | Search nearest location via coordinates. |
| `GET` | `/api/location/{id}/tides/{day}` | Tide table filtered by day and local timezone. |

---

## 🏗️ Architecture

The project uses an approach inspired by **Clean Architecture**, ensuring that the domain logic (tides and geolocation) is independent of external frameworks or database drivers.

The CI/CD pipeline ensures that every change in the `main` branch is tested and made available in production on GCP within a few minutes, without manual intervention.

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---
*Developed as part of the TIPPE project to meet the need for real-time tide table queries anywhere in the world.*

Also part of the TIPPE project:
> Tide table data extractor from the Brazilian Navy documentation:
https://github.com/casimiroarruda/chm-tide-reader

*This project was developed focusing on learning and professional improvement.*

*Made with 💙 by [Anderson Casimiro](https://andr.pt)*
