# Basic cases
export TIME="%E"
> $1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=1' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=21' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=321' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=4321' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=54321' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=654321' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=7654321' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=27654321' 2>>$1 1>/dev/null
/usr/bin/time curl -s 'http://localhost:8000/fib?n=112233445' 2>>$1 1>/dev/null

# Malformed
#curl 'http://localhost:8000/fib?n=22232fib?n=99999'
#curl 'http://localhost:8000/fib?n=%3Cscript%3Ealert(1)%3C/script%3E'

# Big numbers

# Spamming
#while true; do curl 'http://localhost:8000/fib?n=notanumber'; done
#while true; do curl 'http://localhost:8000/fib?n=98765435678982320932873927392739797397397873292012730197301973019730197301987320918732910872301723017820731209731029719070'; done

