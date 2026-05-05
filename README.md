# 🌊 Tide API Go

[![Continuous Integration](https://github.com/casimiroarruda/go-tide-table-api/actions/workflows/ci.yml/badge.svg)](https://github.com/casimiroarruda/go-tide-table-api/actions/workflows/ci.yml)
[![Deploy to Cloud Run](https://github.com/casimiroarruda/go-tide-table-api/actions/workflows/deploy.yml/badge.svg)](https://github.com/casimiroarruda/go-tide-table-api/actions/workflows/deploy.yml)
![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=flat&logo=go)
![PostgreSQL](https://img.shields.io/badge/Database-PostgreSQL%20%2B%20PostGIS-336791?style=flat&logo=postgresql)
![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat)
![Status](https://img.shields.io/badge/Status-Active-green.svg?style=flat)
![Docs](https://img.shields.io/badge/Docs-Swagger-00ADD8?style=flat)
![Test Coverage](https://img.shields.io/badge/Test%20Coverage-89.22%25-green.svg?style=flat)

🇬🇧 [English Version](README.en.md)

🇧🇷 **Versão em Português do Brasil**

A **Tide API Go** é um serviço de alta performance para consulta de tábuas de marés, projetado com foco em precisão geográfica e eficiência de custos. Construída em **Go** e utilizando processamento geoespacial avançado via **PostGIS**, a API fornece dados precisos de marés (alta/baixa) baseados em coordenadas em tempo real ou busca por localidade.

## Motivação

A ideia surgiu da necessidade de ter uma API de marés que fosse rápida, precisa e que pudesse ser usada em qualquer lugar do mundo. 

A Tábua de Maré é uma ferramenta muito utilizada em praticamente todas as áreas litorâneas, como esportes náuticos, pesca, turismo, e etc. Minha família se utiliza muito da informação das tábuas para nossos passeios pela belíssima praia de Boa Viagem, em Recife - Pernambuco. 

Juntando o útil ao agradável: desenvolvi essa API com o objetivo de atender essa demanda e compartilhar as descobertas e aprendizados obtidos no processo. Também é parte de um trabalho de faculdade. Também é um desafio pessoal, pois há muito não colocava um sofware Open Source sozinho em produção - utilizando ferramentas que não eram de meu domínio, como a linguagem Go. 

## Funcionalidades

* **Busca por localidade:** Busca de marés por nome da localidade
* **Busca por coordenadas:** Busca de marés por coordenadas geográficas
* **Tabela de marés:** Exibição da tabela de marés por dia

---

## 🚀 Diferenciais Técnicos

* **Busca Geoespacial (KNN):** Localização instantânea da estação de maré mais próxima via latitude/longitude utilizando o operador `<->` e índices **GIST**, garantindo performance mesmo com grandes volumes de dados.
* **Queries Timezone-Aware:** Lógica de filtragem executada no motor do banco de dados (PostgreSQL), respeitando o fuso horário local de cada estação para definir o "início e fim do dia".
* **SARGable Queries:** Implementação de filtros de data que preservam a eficiência dos índices de banco de dados, otimizando o consumo de CPU no Google Cloud Run.
* **Arquitetura Serverless:** Deploy automatizado via **GitHub Actions** para o **Google Cloud Run**, com injeção de segredos via GCP Secret Manager.
* **Segurança & Precisão:** Autenticação via **JWT** e uso de tipos decimais de precisão fixa (`shopspring/decimal`) para evitar erros de arredondamento em cálculos de altura.

---

## 🛠️ Tech Stack

* **Linguagem:** Go (Golang)
* **Banco de Dados:** PostgreSQL + PostGIS
* **Infraestrutura:** Google Cloud Platform (Cloud Run, Artifact Registry)
* **CI/CD:** GitHub Actions
* **Documentação:** ApiDog / OpenAPI 3.0

---

## 📦 Como Rodar Localmente

### Pré-requisitos
* Docker & Docker Compose
* Go 1.26+

### Passo a Passo
1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/casimiroarruda/go-tide-table-api.git
    cd tide-api-go
    ```
2.  **Configuração de Ambiente:**
    Crie um arquivo `.env` baseado no exemplo:
    ```env
    DATABASE_URL=postgres://user:pass@localhost:5432/tide_db?sslmode=disable
    JWT_SECRET=sua_chave_secreta_aqui
    ```
3.  **Subir Ambiente:**
    ```bash
    docker-compose up -d
    ```
    *O banco de dados **não** executará as migrações automaticamente, incluindo a ativação do PostGIS. Você deve executá-las manualmente. Siga as instruções do arquivo `docs/DEPLOY.md`.*

---

## 📡 API Endpoints

| Método | Endpoint | Descrição |
| :--- | :--- | :--- |
| `POST` | `/api/auth/token` | Geração de token JWT. |
| `GET` | `/api/location` | Lista localidades ou busca por `?name=`. |
| `GET` | `/api/location?lat=&lon=` | Busca localidade mais próxima via coordenadas. |
| `GET` | `/api/location/{id}/tides/{day}` | Tábua de marés filtrada por dia e fuso local. |

---

## 🏗️ Arquitetura

O projeto utiliza uma abordagem inspirada em **Clean Architecture**, garantindo que a lógica de domínio (marés e geolocalização) seja independente de frameworks externos ou drivers de banco de dados.

O pipeline de CI/CD garante que cada alteração na branch `main` seja testada e disponibilizada em produção no GCP em poucos minutos, sem intervenção manual.

---

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para detalhes.

---
*Desenvolvido como parte do projeto TIPPE, para atender à necessidade de consulta de tábuas de marés em tempo real em qualquer lugar do mundo.

Também é parte do projeto TIPPE:
> Extrator de dados de tábuas de marés a partir da documentação da Marinha do Brasil:
https://github.com/casimiroarruda/chm-tide-reader

Este projeto foi desenvolvido com foco em`	 aprendizado e aprimoramento profissional*

*Feito com 💙 por [Anderson Casimiro](https://andr.pt)*

