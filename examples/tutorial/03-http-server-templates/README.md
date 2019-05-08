# HTTP Server Templating

Using https://golang.org/pkg/text/template/ update your endpoint created in `02` to return

```html
<html>
<head>
    <title>{insert_name}</title>
</head>
<body>
Hello, {insert_name}
</body>
```