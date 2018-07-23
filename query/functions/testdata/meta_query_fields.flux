from(db: "test")
    |> range(start:-5m)
    |> filter(fn: (r) => r._measurement == "cpu")
    |> group(by: ["_field"])
    |> distinct(column: "_field")
    |> group(none:true)