module github.com/bernie-pham/simplebank

go 1.18

require (
	github.com/lib/pq v1.10.5
	github.com/stretchr/testify v1.7.1
	golang.org/x/crypto v0.0.0-20220513210258-46612604a0f9
)

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/bernie-pham/simplebank/util => ../../util
