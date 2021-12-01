# Run locally

- [ ] clone repo
- [ ] open terminal in project directory:  
    - `go run main.go`
- [ ] in new terminal tab
    - `curl -H 'Authorization: Bearer {TOKEN}' -X GET 'http://localhost:8000/clicks'`
- [ ] with optional query param 'groupGuid'
    - `curl http://localhost:8000/clicks?groupGuid={GUID} -H "Accept: application/json" -H "Authorization: {TOKEN}" | jq`
    - Falls back to default guid if query param is not formatted correctly
