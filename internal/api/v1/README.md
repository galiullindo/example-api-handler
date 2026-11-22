# api/v1

## Foo

### GET /foo

```bash
curl http://localhost:8080/api/v1/foo
```

### POST /foo

```bash
curl -X POST http://localhost:8080/api/v1/foo \
    -H "Content-Type: application/json; charset=utf-8" \
    -d '{"foo":"foo"}'
```

### GET /foo/{id}

```bash
curl http://localhost:8080/api/v1/foo/1
```

## Bar

### GET /bar

```bash
curl http://localhost:8080/api/v1/bar
```

### POST /bar

```bash
curl -X POST http://localhost:8080/api/v1/bar \
    -H "Content-Type: application/json; charset=utf-8" \
    -d '{"bar":"bar"}'
```

### GET /bar/{id}

```bash
curl http://localhost:8080/api/v1/bar/1
```

## Baz

### GET /baz

```bash
curl http://localhost:8080/api/v1/baz
```

### POST /baz

```bash
curl -X POST http://localhost:8080/api/v1/baz \
    -H "Content-Type: application/json; charset=utf-8" \
    -d '{"baz":"baz"}'
```

### GET /baz/{id}

```bash
curl http://localhost:8080/api/v1/baz/1
```