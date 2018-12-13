REST API endpoint that accepts five parameters : two strings (**string1** and **string2**), and three integers (**int1**, **int2** and **limit**), and returns a JSON

It must return a list of strings with numbers from 1 to **limit**, where:
* all multiples of **int1** are replaced by **string1**,
* all multiples of **int2** are replaced by **string2**,
* all multiples of **int1**  and **int2** are replaced by **string1string2**

## How to run server

```bash
go run main.go
```

Then you can try in your browser : http://localhost:8080/artem?string1=fizz&string2=buzz&int1=3&int2=5&limit=15

## Tests

To run all unit-tests :

```bash
go test ./...
```