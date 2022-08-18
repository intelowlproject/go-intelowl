# Optional Parameters
For the sake of simplicity, we decided that for some endpoints we’ll be passing `Option Parameters` this is to facilitate easy access, configuration and automation so that you don’t need to pass in many parameters but just a simple struct that can be easily converted to and from JSON!

For example, let us look at the `TagParams` we use it as an argument for a method `Create` for `TagService`. From a glance, the `TagParams` look simple. They hold 2 fields: `Label`, and `Color` which can be passed seperatly to the method but imagine if you have many fields! (if you don’t believe see the [`ObservableAnalysisParams`](../../gointelowl/analysis.go))

For a practical implementation you can see the [example](./optionalParams.go)