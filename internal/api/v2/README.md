# api/v2

## Foo

### GET /foo

```bash
curl http://localhost:8080/api/v2/foo
```

### POST /foo

```bash
curl -X POST http://localhost:8080/api/v2/foo \
    -H "Content-Type: application/json; charset=utf-8" \
    -d '{"foo":"foo"}'
```

### GET /foo/{id}

```bash
curl http://localhost:8080/api/v2/foo/1
```

## Bar

### GET /bar

```bash
curl http://localhost:8080/api/v2/bar
```

### POST /bar

```bash
curl -X POST http://localhost:8080/api/v2/bar \
    -H "Content-Type: application/json; charset=utf-8" \
    -d '{"bar":"bar"}'
```

### GET /bar/{id}

```bash
curl http://localhost:8080/api/v2/bar/1
```

## Baz

### GET /baz

```bash
curl http://localhost:8080/api/v2/baz
```

### POST /baz

```bash
curl -X POST http://localhost:8080/api/v2/baz \
    -H "Content-Type: application/json; charset=utf-8" \
    -d '{"baz":"baz"}'
```

### GET /baz/{id}

```bash
curl http://localhost:8080/api/v2/baz/1
```