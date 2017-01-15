# worker

Workers works with data in separate goroutines.

## How to use

1. Create bunch of workers which can work with data and put them into
   `chan Worker` with appropriate buffer
2. Create `chan Data` with no buffer
3. Run goroutine:
  - generates data and send them them to `chan Data`
  - at the end: close `chan Data` (**only when `chan Data` have no buffer**)
4. Run function `DoWork`

When `DoWork` function finish execution you have `chan Data` closed,
`chan Data` with the same set of workers (possible with other order) and
no additional goroutines.

When `chan Data` have buffer you should close can after `DoWork` end
execution (This is not recommended approach).

## Description

- `Data` - anything (`interface {}`)
- `Worker` - function which takes `Data` object
- `DoWork` - function which takes `chan Data` and `chan Worker`

Function `DoWork` take data from appropriate channel and pass them to available worker. `Worker` functions are running in separate goroutines
which ends shortly after `Worker` end execution.
