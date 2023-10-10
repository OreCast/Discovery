# Discovery Service
Discovery service for OreCast.

### OreCast APIs

#### public APIs
- `/sites` get all sites
#### Example
```
# get all sites records
curl http://localhost:8320/sites
```

#### protected APIs
- `/sites` post new site data
- `/site/:site` delete site data

#### Example
```
# record.json
{
    "name":"Cornell",
    "description": "Cornell minerals site",
    "url": "http://127.0.0.1:8330",
    "access_key": "abc",
    "access_secret": "xyz",
    "use_ssl": false,
    "endpoint": "localhost:8330"
}

# inject new record
curl -v -X POST -H "Content-type: application/json" \
    -H "Authorization: Bearer $token" \
    -d@./record.json \
    http://localhost:8320/site
```
