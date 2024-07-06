[![Go Version][gover-img]][gover] [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![GoReport][rpt-img]][rpt]

# Fast and intuitive validation library for Go

This lib uses the `Is...` validation functions from the [govalidator](https://github.com/asaskevich/govalidator) project.

## Installation

```shell
go get github.com/tiendc/go-validator
```

## Usage

#### General usage
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

        // OTHERS
        // Pass if at least one of the validations passes
        OneOf(
            // List of validations
        ),

        // Pass if exact one of the validations passes
        ExactOneOf(
            // List of validations
        ),

        // Pass if none of the validations passes
        NotOf(
            // List of validations
        ),
    )

    for _, e := range errs {
        detail, warnErr := e.BuildDetail()
        fmt.Printf("%+v\n", detail)
    }
```

#### Error message localization

- Method 1: inline localization (not recommended)
```go
    errs := Validate(
        NumLTE(&p.Age, 40).OnError(
            // Override the default template in english
            SetTemplate("Tuổi nhân viên phải nhỏ hơn hoặc bằng {{.Max}}"),
        ),
    )

    for _, e := range errs {
        detail, warnErr := e.BuildDetail()
        fmt.Printf("%+v\n", detail)
    }
```

- Method 2: using another localization lib (recommended)
```go
    // Supposed you have 2 files defining error messages
    // In `error_messages.en`:
    // ERR_EMPLOYEE_AGE_TOO_BIG = "Employee {{.EmployeeName}} has age bigger than {{.Max}}"
    // In `error_messages.vi`:
    // ERR_EMPLOYEE_AGE_TOO_BIG = "Nhân viên {{.EmployeeName}} có tuổi lớn hơn {{.Max}}"

    errs := Validate(
        NumLTE(&p.Age, 40).OnError(
            // Custom param (the default template doesn't have this one)
            SetParam("EmployeeName", p.Name),
            // Custom key to define custom template to use
            SetCustomKey("ERR_EMPLOYEE_AGE_TOO_BIG"),
        ),
    )

    for _, e := range errs {
        errKey := e.CustomKey()
        errParams : = e.Params() // or e.ParamsWithFormatter()
        errorMsg := translationFunction(errKey, errParams) // You need to provide this function
        fmt.Printf("%+v\n", errorMsg)
    }
```

#### Custom error param formatter

```go
    errs := Validate(
        NumLT(&p.Budget, 1000000).OnError(
            SetField("Budget", nil),
        ),
    )

    // e.BuildDetail() may produce message `Budget must be less than 1000000`,
    // but you may want a message like: `Budget must be less than 1,000,000`.
    // Let's use a custom formatter

    errs := Validate(
        NumLT(&p.Budget, 1000000).OnError(
            SetField("Budget", nil),
            SetNumParamFormatter(NewDecimalFormatFunc('.', ',', "%f")),
        ),
    )
```

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