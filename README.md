<div id="top"></div>

[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h3 align="center">TOGO Golang Implementation Sample</h3>

  <p align="center">
    TOGO application implemented using Golang, PostgreSQL, Docker
    <br />
    <a href="#">View Demo</a>
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

### Project Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.

### Implementation Brainstorm

- According to the project requirement, I would love to design this project as a subscription-based TODO application, since different users can have different limited daily TODO tasks creation (e.g. Freemium tier user can create only 1 task/day, Silver tier has 10 tasks/day creation and Gold tier has 100 tasks/day).

- Initial flow:
  - User creates new account, then login to interact with the application through JWT token, then user registers which plan he/she wants to use Freemium/Silver/Gold. If he/she didn't choose any plan yet, Freemium plan is applied by default.
  - Possible endpoints:
    - POST /auth/register: User's account registration.
    - POST /auth/login: User login endpoint.
    - POST /auth/logout: User logout endpoint (Due to lack of time and scope of the challenge, I might not add this endpoint and doing invalidate token stuff with Redis integration)
    - GET /plans: Retrieves all current provided plans.
    - POST /subscribe: User subscribes to chosen plan.
    - GET /tasks: Retrieves all tasks for the current user.
    - GET /tasks/{id}: Retrieves specific task.
    - POST /tasks: Creates new tasks.
    - PUT /tasks/{id}: Modifies specific task.
    - DELETE /tasks/{id}: Soft deletes specific task.

### Built With

- [Golang 1.18](https://go.dev/)
- [PostgreSQL 14.2](https://www.postgresql.org/)
- [Docker 4.6.1](https://www.docker.com/)

<!-- GETTING STARTED -->

## Getting Started

### Prerequisites

- Make sure that you have installed [Docker](https://www.docker.com/) before running the application locally, you can download [Docker Desktop](https://www.docker.com/products/docker-desktop/) so docker cli can be installed accordingly.

  - Check Docker version after installing:

    ```sh
    docker version
    ```

- Make sure to create `.env` file, you can use the most powerful shortcut in developer's world (`Ctrl + C` and `Ctrl + V`) to copy and paste [.env.example](./.env.example) file, rename it to `.env` and change the variables inside `.env` file.

<!-- INSTALLATION -->

### Installation

1. Clone this repository

```sh
git clone https://github.com/TrinhTrungDung/togo.git
```

2. Download and install [Docker Desktop](https://www.docker.com/products/docker-desktop/)

3. Go to the project directory and use this following command to run the application:

```sh
docker-compose up -d
```

4. Migrate database changes:

```sh
go run cmd/migration/main.go
```

<!-- USAGE -->

## Usage

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<!-- CONTACT -->

## Contact

Trinh Trung Dung - [@dungtrungtrinh](https://twitter.com/dungtrungtrinh) - dungtrungtrinh@gmail.com

Project Link: [https://github.com/TrinhTrungDung/togo](https://github.com/TrinhTrungDung/togo)

<!-- LINKS -->

[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/TrinhTrungDung/togo/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/trinhtrungdung/
