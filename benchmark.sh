#!/bin/bash
export TIME="%E"
> $1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=1' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=21' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=321' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=4321' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=54321' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=654321' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=7654321' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=9999999' 2>>$1 1>/dev/null
#/usr/bin/time curl -s 'http://localhost:8000/fib?n=27654321' 2>>$1 1>/dev/null
#/usr/bin/time curl -s 'http://localhost:8000/fib?n=99999999' 2>>$1 1>/dev/null
#/usr/bin/time curl -s 'http://localhost:8000/fib?n=112233445' 2>>$1 1>/dev/null
