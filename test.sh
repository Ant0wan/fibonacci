# Basic cases
curl 'http://localhost:8000/fib?n=1'
curl 'http://localhost:8000/fib?n=21'
curl 'http://localhost:8000/fib?n=321'
curl 'http://localhost:8000/fib?n=4321'
curl 'http://localhost:8000/fib?n=54321'
curl 'http://localhost:8000/fib?n=654321'
curl 'http://localhost:8000/fib?n=7654321'
curl 'http://localhost:8000/fib?n=87654321'
curl 'http://localhost:8000/fib?n=987654321'
curl 'http://localhost:8000/fib?n=1987654321'

#curl 'http://localhost:8000/fib?n=222323230983049820347032634926918623872197391273109876543456789876543212345678998776862117639817639186876329187632917863918'

# Malformed
#curl 'http://localhost:8000/fib?n=22232fib?n=99999'
#curl 'http://localhost:8000/fib?n=%3Cscript%3Ealert(1)%3C/script%3E'

# Big numbers



# Spamming
#while true; do curl 'http://localhost:8000/fib?n=notanumber'; done
#while true; do curl 'http://localhost:8000/fib?n=98765435678982320932873927392739797397397873292012730197301973019730197301987320918732910872301723017820731209731029719070'; done

