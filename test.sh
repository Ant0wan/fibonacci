# Basic cases
curl 'http://localhost:8000/fib?n=1'
curl 'http://localhost:8000/fib?n=21'
curl 'http://localhost:8000/fib?n=321'
curl 'http://localhost:8000/fib?n=4321'
curl 'http://localhost:8000/fib?n=54321'
curl 'http://localhost:8000/fib?n=654321'
curl 'http://localhost:8000/fib?n=7654321'
curl 'http://localhost:8000/fib?n=17654321'
curl 'http://localhost:8000/fib?n=101654321'

# Malformed
curl 'http://localhost:8000/fib?n=22232fib?n=99999'
curl 'http://localhost:8000/fib?n=%3Cscript%3Ealert(1)%3C/script%3E'

# Big numbers
#... could not succeed

# Spamming
while true; do curl 'http://localhost:8000/fib?n=notanumber'; done
while true; do curl 'http://localhost:8000/fib?n=98765435678982320932873927392739797397397873292012730197301973019730197301987320918732910872301723017820731209731029719070'; done

