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
    import (
        vld "github.com/tiendc/go-validator"
    )

    type Person struct {
        FirstName string
        LastName  string
        Birthdate time.Time

        Unemployed bool
        Salary     uint
        Rank       string
        WorkEmail  string
        Projects   []string
    }
    var p Person

    errs := vld.Validate(
        // Validate first and last names separately
        vld.StrLen(&p.FirstName, 3, 30).OnError(
            vld.SetField("first_name", nil),
            vld.SetCustomKey("ERR_VLD_PERSON_FIRST_NAME_INVALID"),
        ),
        vld.StrLen(&p.FirstName, 3, 30).OnError(
            vld.SetField("last_name", nil),
            vld.SetCustomKey("ERR_VLD_PERSON_LAST_NAME_INVALID"),
        ),

        // OR use this to produce only one error when one of them fails
        vld.Group(
            vld.StrLen(&p.FirstName, 3, 30),
            vld.StrLen(&p.LastName, 3, 30),
        ).OnError(
            vld.SetField("name", nil),
            vld.SetCustomKey("ERR_VLD_PERSON_NAME_INVALID"),
        ),

        // Birthdate is optional, but when it's present, it must be within 1950 and now
        vld.When(!p.Birthdate.IsZero()).Then(
            vld.TimeRange(p.Birthdate, <1950-01-01>, time.Now()).OnError(...),
        )

        vld.When(!p.Unemployed).Then(
            vld.Required(&p.Salary),
            // Work email must be valid
            vld.StrIsEmail(&p.WorkEmail),

            // Rank must be one of the constants
            vld.StrIn(&p.Rank, "Employee", "Manager", "Director"),
            vld.Case(
                vld.When(p.Rank == "Manager").Then(vld.NumGT(&p.Salary, 10000)),
                vld.When(p.Rank == "Director").Then(vld.NumGT(&p.Salary, 30000)),
            ).Default(
                vld.NumLT(&p.Salary, 10000),
            ),

            // Projects are optional, but when they are present, they must be unique and sorted
            vld.When(len(p.Projects) > 0).Then(
                vld.SliceUnique(p.Projects).OnError(...),
                vld.SliceSorted(p.Projects).OnError(...),
            )
        ).Else(
            // When person is unemployed
            vld.NumEQ(&p.Salary, 0),
            vld.StrEQ(&p.WorkEmail, ""),
        ),

        // Validate slice elements
        vld.Slice(p.Projects).ForEach(func(elem int, index int, validator ItemValidator) {
            validator.Validate(
                vld.StrLen(&elem, 10, 30).OnError(
                    vld.SetField(fmt.Sprintf("projects[%d]", index), nil),
                    vld.SetCustomKey("ERR_VLD_PROJECT_NAME_INVALID"),
                ),
            )
        }),

        // OTHER FUNCTIONS
        // Pass if at least one of the validations passes
        vld.OneOf(
            // List of validations
        ),

        // Pass if exact one of the validations passes
        vld.ExactOneOf(
            // List of validations
        ),

        // Pass if none of the validations passes
        vld.NotOf(
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
    // ERR_VLD_EMPLOYEE_AGE_TOO_BIG = "Employee {{.EmployeeName}} has age bigger than {{.Max}}"
    // In `error_messages.vi`:
    // ERR_VLD_EMPLOYEE_AGE_TOO_BIG = "Nhân viên {{.EmployeeName}} có tuổi lớn hơn {{.Max}}"

    errs := Validate(
        NumLTE(&p.Age, 40).OnError(
            // Custom param (the default template doesn't have this one)
            SetParam("EmployeeName", p.Name),
            // Custom key to define custom template to use
            SetCustomKey("ERR_VLD_EMPLOYEE_AGE_TOO_BIG"),
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
        NumLT(&budget, 1000000).OnError(
            SetField("Budget", nil),
        ),
    )

    // e.BuildDetail() may produce message `Budget must be less than 1000000`,
    // but you may want a message like: `Budget must be less than 1,000,000`.
    // Let's use a custom formatter

    errs := Validate(
        NumLT(&budget, 1000000).OnError(
            SetField("Budget", nil),
            SetNumParamFormatter(NewDecimalFormatFunc('.', ',', "%f")),
        ),
    )
```

## Contributing

- You are welcome to make pull requests for new functions and bug fixes.

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