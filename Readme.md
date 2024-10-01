to run actual implementation which has a database connection simulation

```
go run main.go
```

to run tests

```
go test -v
```

<br/>
I have tried to create a product repository interface and aproduct service struct which has a field of type product repository interface

on initilaizing the product service we pass an instance of implementation of the product repository in the product service

```
productService := service.ProductService{
		ProductRepository: &repository.ProductRepositoryImpl{}}
```

where we are basically doing the dependency injection

for testing purpose we are injecting a mock implementation of the Product Repository interface

```
cases := []struct {
		productRepository repository.ProductRepository
		expected          int
	}{
		{
			productRepository: &ProductRepositoryMockImplSuccess{},

			expected: 2,
		},
		{
			productRepository: &ProductRepositoryMockImplFailed{},

			expected: 2,
		},
	}
	for _, c := range cases {
		productService := service.ProductService{
			ProductRepository: c.productRepository}
		availableProducts, error := productService.GetAvaialbleProducts("Electronics")
		if error != nil {
			panic(error)
		}
		if len(availableProducts) != c.expected {
			t.Errorf("Expected %v but got %v", c.expected, len(availableProducts))
		}
	}
```
