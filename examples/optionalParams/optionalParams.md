# Optional Parameters
For the sake of simplicity we decided that for some endpoints we'll be passing `Option Parameters` this is to faciliate easy access, configuration and automation so that you don't need to pass in many parameters but just a simple struct that can be easily converted to and from JSON!

For example lets look at the `TagParams` we use it as an argument for a method `Create` for `TagService`. From a glance the `TagParams` looks simple holds 2 fieds: `Label`, `Color` which can be passed seperatly to method but imagine if you have many fields! (if you don't believe see the [`ObservableAnalysisParams`](../../gointelowl/analysis.go))

Now if you've read my TED talk you can see the [example](./optionalParams.go)