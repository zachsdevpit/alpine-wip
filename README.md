# alpine-wip

---

First Look At
[Alpine.js](https://alpinejs.dev)

## Tailwinds

---
Using The Cli For CSS Generation

```bash
npx @tailwindcss/cli -i resources/css/input.css -o resources/css/output.css --watch
```

## Bundle JS Dependencies

---
Opted to not use the CDN and instead bundled js concerns using [Laravel-Mix](https://laravel-mix.com)

```bash
npx mix 
```

Should generate the necessary /dist/app.js
