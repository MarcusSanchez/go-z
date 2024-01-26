usage:

```go

type Nested struct {
	Street  string `z:"street"`
	Address string `z:"address"`
}

type Payload struct {
	Name   string  `z:"name"`
	Email  string  `z:"email"`
	Age    int     `z:"age"`
	Height float64 `z:"height"`
	Nested Nested  `z:"nested"`
}

Schema := z.Struct{
	"name": z.String().Min(3, "Name must be at least 3 character long"),
	"email": z.String().
		Email("email must be valid").
		Min(5, "email must be at least 5 characters long").
		Max(50, "email must be less than 50 characters long"),
	"age": z.Int().
		GT(0, "age must be greater than 0").
		LT(120, "age must be less than 120"),
	"height": z.Float64().
		GT(1.3, "height must be greater than 1.3 meters").
		LT(2.1, "height must be less than 2.1 meters"),
	"nested": z.Struct{
		"street":  z.String().Min(5),
		"address": z.String().Min(5),
	},
}

payload := &Payload{
	Name:   "John Doe",
	Email:  "jdoe1@gmail.com",
	Age:    30,
	Height: 1.8,
	Nested: Nested{
		Street:  "123 Apple St",
		Address: "New York",
	},
}

if err := Schema.Validate(payload); err != nil {
	for _, e := range z.ErrSlice(err) {
		fmt.Println(e)
	}
}

```