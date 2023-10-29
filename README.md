[![Go Version][gover-img]][gover] [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![GoReport][rpt-img]][rpt]

# Fast and intuitive validation library for Go

This lib uses the `Is...` validation functions from the [govalidator](https://github.com/asaskevich/govalidator) project.

## Installation

```shell
go get github.com/tiendc/go-validator
```

## Usage

```go
    type Person struct {
        Name      string
        Birthdate time.Time

        Unemployed bool
        Salary     uint
        Rank       string
        WorkEmail  string
    }

    p := Person{
        Name:       "",
        Unemployed: false,
        Rank:       "Manager",
        Salary:     1000,
    }

    errs := Validate(
        Required(&p.Name),
        StrLen(&p.Name, 3, 100).OnError(
            SetField("name", nil),
            SetCustomKey("ERR_PERSON_NAME_REQUIRED"),
        ),
        Required(&p.Birthdate),
        When(!p.Unemployed).Then(
            Required(&p.Salary),

            // WorkEmail is required and must be valid
            Required(&p.WorkEmail),
            StrIsEmail(&p.WorkEmail),
            // OR use this to produce only one error when the validation fails
            Group(
                Required(&p.WorkEmail),
                StrIsEmail(&p.WorkEmail),
            ).OnError(...),

            StrIn(&p.Rank, "Employee", "Manager", "Director"),
            Case(
                When(p.Rank == "Manager").Then(NumGT(&p.Salary, 10000)),
                When(p.Rank == "Director").Then(NumGT(&p.Salary, 30000)),
            ).Default(
                NumLT(&p.Salary, 10000),
            ),
        ).Else(
            NumEQ(&p.Salary, 0),
            StrEQ(&p.WorkEmail, ""),
        ),
    )

    for _, e := range errs {
        detail, warnErr := e.BuildDetail()
        fmt.Printf("%+v\n", detail)
    }
```

## Usage

    TBD

## Contributing

- You are welcome to make pull requests for new functions and bug fixes.

## Authors

- Dao Cong Tien ([tiendc](https://github.com/tiendc))

## License

- [MIT License](LICENSE)

[doc-img]: https://pkg.go.dev/badge/github.com/tiendc/go-validator
[doc]: https://pkg.go.dev/github.com/tiendc/go-validator
[gover-img]: https://img.shields.io/badge/Go-%3E%3D%201.20-blue
[gover]: https://img.shields.io/badge/Go-%3E%3D%201.20-blue
[ci-img]: https://github.com/tiendc/go-validator/actions/workflows/go.yml/badge.svg
[ci]: https://github.com/tiendc/go-validator/actions/workflows/go.yml
[cov-img]: https://codecov.io/gh/tiendc/go-validator/branch/main/graph/badge.svg
[cov]: https://codecov.io/gh/tiendc/go-validator
[rpt-img]: https://goreportcard.com/badge/github.com/tiendc/go-validator
[rpt]: https://goreportcard.com/report/github.com/tiendc/go-validator