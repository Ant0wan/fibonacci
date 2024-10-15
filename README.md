
# Fibonacci Number Service

This service computes Fibonacci numbers using an optimized algorithm designed for performance and security. Written in Go, it leverages the language's speed and offers a minimal attack surface with a statically built, lightweight container.

## Key Features

- **Optimized Fibonacci Calculation**: Implements an O(log n) algorithm for efficient Fibonacci computation. Alternative approaches like recursion (O(2^n)) or memoization (O(n)) were considered but are slower.
- **Minimalistic Container**: Runs in a highly secure, statically built Go environment with no external libraries or binaries. The image is built from scratch to limit the attack surface.
- **Security Hardened**: Deployed with gVisor, network policies, and other security measures to mitigate potential container hijacking or abuse.

## Table of Contents
1. [Algorithm Overview](#algorithm-overview)
2. [Security Features](#security-features)
3. [Performance and Optimization](#performance-and-optimization)
4. [Deployment](#deployment)
5. [Try It With Docker](#try-it-with-docker)
6. [Try It With Kubernetes](#try-it-with-kubernetes)
7. [References](#references)

## Algorithm Overview

This service uses the **fast doubling** algorithm to compute Fibonacci numbers with a time complexity of O(log n). This approach is faster than traditional recursive (O(2^n)) or memoization (O(n)) methods but lacks comprehensive documentation. Development was completed before fully exploring this algorithm, though it remains an efficient solution.

- **Go for Speed**: Go was chosen for its ability to produce fast, statically compiled binaries. It enables us to create lightweight and secure containers ideal for production environments.
- **Big Number Libraries**: The implementation initially tested two libraries for handling large numbers: `GMP` and Go's `math/big`. Benchmarks showed that `math/big` performed slightly better and had the added advantage of being a Go-native library.

### Limitations

- Calculations for very large Fibonacci numbers (greater than 8 digits) may take over a minute to compute.
- Future improvements could include estimating the compute time and rejecting numbers that would exceed a certain threshold before starting the calculation.

## Security Features

This service has been built with a security-first approach:

- **Container Isolation**: The service runs in a container on **gVisor**, an additional sandbox layer to protect the host from potential container escape attacks.
- **Network Policies**: Configured with Ingress-only network policies and a default-deny rule to prevent communication with other services if the pod is compromised.
- **Read-Only Filesystem**: The container runs with a **read-only root filesystem** to limit potential write operations from inside the container.
- **User Privilege Management**: The service is executed by a non-root user (`runAsUser 1000`).
- **Static Analysis**: Trivy is used for static analysis to prevent deployments with known vulnerabilities.

Additional hardening ideas that were considered but not yet implemented:
- IP banning or limiting the number of requests per user.
- Setting a timeout for request handling to avoid abuse.
- Enforcing AppArmor/SELinux profiles for more granular security.
- Fuzzing tests to explore edge cases in the algorithm.
- Logging errors.

## Performance and Optimization

The Fibonacci computation is optimized through:

- **Benchmarking**: Extensive benchmarks using a custom `benchmark.sh` script measured the performance of different big number libraries (`GMP` vs `math/big`). Go's `math/big` library was chosen based on performance.
- **Static Build**: The Go binary is statically compiled, which keeps the image lightweight (~7MB) and minimizes dependencies.

## Deployment

### Minimal Docker Image

The service is packaged in a minimal Docker container for efficient deployment.

- **Build**: The container is built from scratch with no external libraries or binaries, providing a minimal attack surface.
- **Trivy Scan**: Before deployment, the container is scanned for vulnerabilities using Trivy.
- **Horizontal Pod Autoscaler (HPA)**: For dynamic scaling, HPA can be enabled, although it requires a Metrics Server.

### Image Details

The final Docker image is lightweight:

```
REPOSITORY   TAG       IMAGE ID       CREATED          SIZE
fibonacci    0         8eab71ceddce   21 seconds ago   7.03MB
```

## Try It With Docker

You can try out the Fibonacci service locally using Docker. Follow these steps:

1. **Build the Docker Image**:

    ```bash
    docker build --tag fibonacci:0 .
    ```

2. **Scan the Image for Vulnerabilities**:

    ```bash
    trivy image fibonacci:0
    ```

3. **Run the Service**:

    ```bash
    docker run --interactive --tty --publish 8000:8000 fibonacci:0
    ```

4. **Make a Request**:

    To compute the Fibonacci number for `n=7654321`, run:

    ```bash
    curl -s 'http://localhost:8000/fib?n=7654321'
    ```

## Try It With Kubernetes

You can try out the full Fibonacci deployment on Kubernetes. Just one command away:

   **Apply the Manifest**:

    ```bash
    kubectl apply -f fib.yaml
    ```


## References

- [Fast Fibonacci Algorithms by Nayuki](https://www.nayuki.io/page/fast-fibonacci-algorithms)
- [Fibonacci Matrix Exponentiation Algorithm](https://robwilsondev.medium.com/bigo-and-beyond-how-to-compute-fibonacci-sequence-efficiently-with-matrix-exponentiation-d9924545fe54)
- [Fibonacci Calculator for Verification](https://www.calculatorsoup.com/calculators/discretemathematics/fibonacci-calculator.php)
