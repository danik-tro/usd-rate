### USD-RATE 

# Getting started

## Running with docker-compose

When you're ready, start your application by running:
```bash
docker compose up --build
```

Your application will be available at http://localhost:8000.

## Docs
Project documentation is available at the link: http://localhost:8000/docs.

open.er-api.com was chosen as a third-party service for the possibility of using a free plan without an API token.

### Technologies Used
*	Go: Chosen for its performance and simplicity.
*	Fiber: A fast web framework for building applications.
*	GORM: An ORM library for Go, providing convenient access to databases and built-in migration support.
*	PostgreSQL: The primary database for this project.
*	Redis: Used for caching to improve performance and reduce load on the primary database.

### Sending emails

To send emails, a cron job setup library was used, which sends a request to the email sending endpoint, which happens asynchronously in a goroutine.

All the main information is stored in Postgres. Redis is used as a cache to save the last dollar exchange rate and to reduce requests for third-party services. The cache lives for 10 minutes, so a third-party service request will be made once every 10 minutes.
