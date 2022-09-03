## How unit tests were written
The unit tests were written as a combination of [table driven tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests) and the approach used by [go-github](https://github.com/google/go-github)

Firstly we use a `TestData` struct that has the following fields:
1. `Input` - this is an `interface` as it is to be used as the input required for an endpoint
2. `Data` - this is a `string` as it'll be the `JSON` string that the endpoint is expected to return
2. `StatusCode` - this is an `int` as it is meant to be used as the expected response returned by the endpoint
3. `Want` - the expected struct that the method will return

Now the reason we made this was that these fields were needed for every endpoint hence combining them into a single struct provided us reusability and flexibility.

Now the testing suite used go's `httptest` library where we use `httptest.Server` as this setups a test server so that we can easily mock it. We also use `http.ServerMux` to mock our endpoints response.

## How to add a new test for an endpoint
Lets say IntelOwl added a new endpoint called **supercool** in `Tag`. Now you've implemented the endpoint as a method of `TagService` and now you want to add its unit tests.

First go to `tagService_test.go` in the `tests` directory and add

```Go
func TestSuperCoolEndPoint(t *testing.T) {
	testCases := make(map[string]TestData)
	testCases["simple"] = TestData{
		Input:      nil,
		Data:       `{ "supercool": "you're a great developer :)"}`,
		StatusCode: http.StatusOK,
		Want: "you're a great developer :)",
	}
	for name, testCase := range testCases {
		// subtest
		t.Run(name, func(t *testing.T) {
			// setup will give you the client, mux/router, closeServer
			client, apiHandler, closeServer := setup()
			defer closeServer()
			ctx := context.Background()
			// now you can use apiHandler to mock how the server will handle this endpoints request
			// you can use mux/router's Handle method or HandleFunc
			apiHandler.Handle("/api/tag/supercool", func(w http.ResponseWriter, r *http.Request) {
				// this is a helper test to check if it is the expected request sent by the client
				testMethod(t, r, "GET")
				w.Write([]byte(testCase.Data))
			})
			expectedRespone, err := client.TagService.SuperCool(ctx)
			if err != nil {
				testError(t, testCase, err)
			} else {
				testWantData(t, testCase.Want, expectedRespone)
			}
		})
	}
}

```

Great! Now you've added your own unit tests.

