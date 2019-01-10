# flat
flattens/unflattens map or JSON string

```
js := `{
  "name": {
    "first_name": "John",
    "last_name": "Doe"
  }
}`
js, err := flat.FlattenJSON(js)
// output: {"name.first_name": "John", "name.last_name": "Doe"}

js, err = flat.UnflattenJSON(js)
```

*NOTE:* Use at your own risk ;-)
