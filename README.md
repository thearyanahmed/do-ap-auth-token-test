# test
```bash
curl "https://yourdomain.com" -H "Content-Type: application/json" \ -d '{}' -i -s -o /dev/null -D - | grep "token"
curl -X POST "https://yourdomain.com" -H "Content-Type: application/json" \ -d '{}' -i -s -o /dev/null -D - | grep "token"
```
