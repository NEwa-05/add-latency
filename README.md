# Add Latency with Traefik

## Why this plugin

For specific test you may need to simulate latency.

This simple plugin let you add several seconds through Traefik processin when handling request up to your backend.

## Configuration

Add the latency you want:

```yaml
# Dynamic configuration
http:
  middlewares:
    addlatency:
      plugin:
        add-latency:
          addedLatency: 20
```

The following declaration (given here in YAML) defines a plugin:

```yaml
# Static configuration
experimental:
  plugins:
    example:
      moduleName: github.com/NEwa-05/add-latency
      version: v0.0.1
```
