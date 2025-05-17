# TicketHub - Scalable Event Ticketing System 🎟️

TicketHub is a modern and scalable event ticketing system designed to provide seamless event management, ticket purchases, and payment processing. Built with **Golang** and **PostgreSQL** for the backend, and **Next.js** with **TypeScript** for the frontend, this system ensures high performance, secure authentication, and real-time features.

## Features 🌟
- **User Authentication:** Secure login and registration with JWT token management.
- **Event Management:** Organizers can create, edit, and manage events easily.
- **Ticket Purchase Flow:** Real-time ticket inventory updates and integrated payment system simulation using **Stripe.js**.
- **Mobile-First Design:** A responsive UI built with **Tailwind CSS** for seamless mobile and desktop experience.
- **Real-time Updates:** Live updates for event status and ticket availability to ensure accurate information is always displayed.
- **Admin Panel:** A dashboard for event organizers to track ticket sales, manage event details, and handle payment information.
- **Search and Filters:** Search for events by location, category, and date with an intuitive filtering system.

## Built With 🛠️
**Backend & DevOps:**

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://gin-gonic.com/)
[![gRPC](https://img.shields.io/badge/gRPC-4285F4?style=for-the-badge&logo=google&logoColor=white)](https://grpc.io/)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-316192?style=for-the-badge&logo=postgresql&logoColor=white)](https://www.postgresql.org/)
[![Redis](https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white)](https://redis.io/)
[![Docker](https://img.shields.io/badge/Docker-2CA5E0?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)
[![AWS](https://img.shields.io/badge/AWS-232F3E?style=for-the-badge&logo=amazonaws&logoColor=white)](https://aws.amazon.com/)

**Frontend & UI/UX:**

[![Next.js](https://img.shields.io/badge/Next.js-000000?style=for-the-badge&logo=next.js&logoColor=white)](https://nextjs.org/)
[![React](https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB)](https://react.dev/)
[![TypeScript](https://img.shields.io/badge/TypeScript-3178C6?style=for-the-badge&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-06B6D4?style=for-the-badge&logo=tailwind-css&logoColor=white)](https://tailwindcss.com/)
[![Stripe](https://img.shields.io/badge/Stripe-6772E5?style=for-the-badge&logo=stripe&logoColor=white)](https://stripe.com/)

## Getting Started 🚀

To set up the project locally, follow these steps:

### Prerequisites 🛠️
- **Backend:**
  - Go 1.21+
  - Docker & Docker Compose
  - PostgreSQL 15
  - Redis 7

- **Frontend:**
  - Node.js 20+
  - npm or yarn

### Installation Steps 🔧

1. **Clone the repository:**
   ```sh
   git clone https://github.com/shristirajpoot/eventura.git
   cd eventura
Set up environment variables:

sh
Copy
Edit
cp .env.example .env
Start dependencies:

sh
Copy
Edit
docker-compose up -d postgres redis
Run database migrations:

sh
Copy
Edit
make migrateup
Start the server:

sh
Copy
Edit
make server
Access the App:

Frontend: http://localhost:3000
Backend API: http://localhost:8080
Contributing 🤝
We welcome contributions to improve the system!

Fork the repository.
Create a new branch: git checkout -b feature/your-feature.
Commit your changes: git commit -m 'Add feature'.
Push your changes: git push origin feature/your-feature.
Submit a Pull Request.

## Deployment & Monitoring Tips

Use AWS EC2 or Render to deploy backend containers.

Integrate Grafana + Prometheus for live monitoring.

Use Logrus or Zap for structured logging.

## Learning & Takeaways

While building TicketHub, I deepened my understanding of:

Designing scalable backend systems using Golang.

Dockerizing microservices and managing multi-container environments.

Implementing real-time features and resilient infrastructure.

Setting up CI/CD pipelines and handling environment-based configurations.

Writing clean, testable, production-level code with modular architecture
