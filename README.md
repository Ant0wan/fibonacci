# A Fibonacci Number Service

## Comments

Calcul vectoriel pour le process en go which is the fastest model to compute it because it is the only approch offering a O(log n), other approches: reccursive O(2^n) or memoization O(n)
"Fast doubling" is an algorithm know to be faster but lack of documentation and mostly acknowledging it too late once development was done.

Use go for speed and enabeling secure and lightweight package thanks to static build. No lib nor bin available in image from scratch limiting surface attack.
This needs to be set for production purpose - to fix an abuse limit.
It could dynamically be set depending on the hardware it runs on.
Comment this for test purposes, if you do not mind wait for compute.
Estimate the time it will take to compute fibonacci on the number given before computing in order to reject it.
Could have added logs
Could have added health point and metrics use
Could have added ip ban or single request per user
Could have added a timeout on requests
Could have tested fuzzing
Could not handle bigger numbers than ~8 digit without taking 1min or more
Could have done apparmor/selinux profiles for running fibonacci binary.
Could decide to compute or not depending upon the length of its result (compute it approximatly before computing fibonacci)
Could have done network policies with this Ingress only on port 8000

Container runs on gvisor for limiting node disaster if container hijacked.

Algorithm in lib/fibonacci has been optimize using the benchmark.sh which time requests. The algorithm tested mainly 2 libraries to handle bignumbers GMP (gmp.log) and math/big (big.log). Benchmark showed that math/big was a little better and luckily is "go native" library. The implementation was improved with help of chatGPT.

Network policies with Ingress only and default-deny for the rest to avoid reaching other endpoints if pod has been hijacked, also preventing spamming other services

Deployment in readonlyrootfilesystem, runAs 1000, with gVisor used.

Use Trivy as static analysis to prevent a deployment with known security issues.

The service is so simple that it is not necessary to use Helm or Kustomize at the moment.
CI not implemented, build must be manual unfortunatly.

The use of HPA requires Metrics Server in order to work properly.

Documentation and comments in code written with help of chatGPT.

Sources:
https://robwilsondev.medium.com/bigo-and-beyond-how-to-compute-fibonacci-sequence-efficiently-with-matrix-exponentiation-d9924545fe54
https://www.nayuki.io/page/fast-fibonacci-algorithms
https://www.calculatorsoup.com/calculators/discretemathematics/fibonacci-calculator.php - for verification

## Try it with Docker
docker build --tag fibonacci:0 .

trivy image fibonacci:0

docker run --interactive --tty --publish 8000:8000 fibonacci:0

curl -s 'http://localhost:8000/fib?n=7654321'

## The Image Details

Go static build for lightweight image

```
REPOSITORY                    TAG       IMAGE ID       CREATED          SIZE
fibonacci                     0         8eab71ceddce   21 seconds ago   7.03MB
```

