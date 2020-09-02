Datatypes for Miller records and context-variables.

* `containers.Lrec` is the sequence of key-value pairs which represents a Miller record. The key-lookup mechanism is optimized for Miller read/write usage patterns -- please see `lrec.go` for more details.
* `containers.Context` supports AWK-like variables such as `FILENAME`, `NF`, `NR`, and so on.