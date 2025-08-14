# FULLSTACK-MULTICONTAINER

Seamless Multi-Container Power for Rapid Innovation

![last commit](https://img.shields.io/github/last-commit/DandyMaulana19/FullStack-MultiContainer?color=brightgreen)
![today](https://img.shields.io/badge/today-blue)
![blade](https://img.shields.io/badge/blade-49.1%25-blue)
![languages](https://img.shields.io/badge/languages-6-gray)

**Built with the tools and technologies:**

![JSON](https://img.shields.io/badge/-JSON-black?logo=json&logoColor=white)
![Markdown](https://img.shields.io/badge/-Markdown-black?logo=markdown&logoColor=white)
![npm](https://img.shields.io/badge/-npm-red?logo=npm&logoColor=white)
![Autoprefixer](https://img.shields.io/badge/-Autoprefixer-E34F26?logo=autoprefixer&logoColor=white)
![PostCSS](https://img.shields.io/badge/-PostCSS-DD3A0A?logo=postcss&logoColor=white)
![Composer](https://img.shields.io/badge/-Composer-885630?logo=composer&logoColor=white)
![JavaScript](https://img.shields.io/badge/-JavaScript-F7DF1E?logo=javascript&logoColor=black)
![Now](https://img.shields.io/badge/-NOW-000000?logo=vercel&logoColor=white)
![Go](https://img.shields.io/badge/-Go-00ADD8?logo=go&logoColor=white)
![React](https://img.shields.io/badge/-React-61DAFB?logo=react&logoColor=black)
![Gin](https://img.shields.io/badge/-Gin-00ADD8?logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/-MySQL-4479A1?logo=mysql&logoColor=white)
![Docker](https://img.shields.io/badge/-Docker-2496ED?logo=docker&logoColor=white)
![XML](https://img.shields.io/badge/-XML-FF6600?logo=xml&logoColor=white)
![PHP](https://img.shields.io/badge/-PHP-777BB4?logo=php&logoColor=white)
![Vite](https://img.shields.io/badge/-Vite-646CFF?logo=vite&logoColor=white)
![Axios](https://img.shields.io/badge/-Axios-5A29E4?logo=axios&logoColor=white)
![Bootstrap](https://img.shields.io/badge/-Bootstrap-7952B3?logo=bootstrap&logoColor=white)
![YAML](https://img.shields.io/badge/-YAML-CB171E?logo=yaml&logoColor=white)

---

## Table of Contents

- [Overview](#overview)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Testing](#testing)

---

## Overview

FullStack-MultiContainer is an integrated, containerized platform designed to streamline the development, deployment, and management of multi-service contact form applications.  
It combines a Laravel-based CMS, a Go API backend, and a Next.js client interface within a cohesive Docker environment.

### Why FullStack-MultiContainer?

This project simplifies complex multi-service architectures by providing a ready-to-deploy environment with robust tooling.  
The core features include:

- 🌊 **Container Orchestration:** Uses Docker Compose to manage services like database, API, frontend, and web server for reliable, consistent deployments.
- 🍃 **Modern Frontend:** Implements a Next.js client with Tailwind CSS for a responsive, user-friendly contact form experience.
- 🛠 **Modular Backend:** Combines Laravel for content management and a scalable Go API for contact data handling, ensuring flexibility and performance.
- ✅ **Developer-Friendly:** Includes comprehensive testing configurations, environment management, and standardized code styling for efficient development.
- 🔒 **Secure & Reliable:** Configures PHP, Nginx, and database services for secure, high-performance operation within a multi-container setup.

---

## Getting Started

### Prerequisites

This project requires the following dependencies:

- **Programming Language:** PHP
- **Package Manager:** Npm, Composer, Go modules
- **Container Runtime:** Docker

---

### Installation

Build FullStack-MultiContainer from the source and install dependencies:

1. **Clone the repository:**
   ```bash
   git clone https://github.com/DandyMaulana19/FullStack-MultiContainer
   ```

2. **Navigate to the project directory:**
   ```bash
   cd FullStack-MultiContainer
   ```

3. **Install the dependencies:**

   **Using [docker](https://docs.docker.com/):**
   ```bash
   docker build -t DandyMaulana19/FullStack-MultiContainer .
   ```

   **Using [npm](https://www.npmjs.com/):**
   ```bash
   npm install
   ```

   **Using [composer](https://getcomposer.org/):**
   ```bash
   composer install
   ```

   **Using [go modules](https://go.dev/doc/modules):**
   ```bash
   go build
   ```

---

### Usage

Run the project with:

**Using [docker](https://docs.docker.com/):**
```bash
docker run -it {image_name}
```

**Using [npm](https://www.npmjs.com/):**
```bash
npm start
```

**Using [composer](https://getcomposer.org/):**
```bash
php {entrypoint}
```

**Using [go modules](https://go.dev/doc/modules):**
```bash
go run {entrypoint}
```

---

### Testing

Fullstack-multicontainer uses the **{test_framework}** test framework.  
Run the test suite with:

**Using [docker](https://docs.docker.com/):**
```bash
echo 'INSERT-TEST-COMMAND-HERE'
```

**Using [npm](https://www.npmjs.com/):**
```bash
npm test
```

**Using [composer](https://getcomposer.org/):**
```bash
vendor/bin/phpunit
```

**Using [go modules](https://go.dev/doc/modules):**
```bash
go test ./...
```
