# Stress Test Application 📈

Welcome to the **Stress Test Application**! This tool allows you to run HTTP stress tests with high flexibility and speed.

## Introduction 🚀
The Stress Test Application is designed to help you test the robustness and performance of your HTTP servers. By sending a specified number of requests with a defined level of concurrency, you can gauge how well your server handles high load.

## Installation 🔧
To install the application, follow these steps:

1. Ensure you have Docker installed on your system. If not, download it from [docker.com](https://www.docker.com/).
2. Run ``docker build -t myapp:latest .`` to build the image
3. Run ``docker run --rm myapp:latest --requests=100 --url=http://www.google.com --concurrency=2`` to run the app


## Flags 🚩
The application supports the following flags:

```--url (string)```: The target URL for the stress test.

```--requests (uint64)```: The total number of requests to send.

```--concurrency (uint64)```: The number of concurrent requests to send.
