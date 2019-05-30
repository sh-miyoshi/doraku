# Test Data

## Overview

This directory includes data for test.

## Usage

1. Login

```bash
curl -X POST http://localhost:8080/api/v1/login -d "@login.json"
```

2. Get User Data

```bash
TOKEN=`curl -s -X POST http://localhost:8080/api/v1/login -d "@login.json" | jq .token | tr -d '"'`
curl -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/v1/user/test
```

3. Create New User

```bash
curl -X POST -d "@new_user.json" http://localhost:8080/api/v1/user
curl -X POST -d "@new_user.json" http://localhost:8080/api/v1/login
```

4. Delete User Data

```bash
TOKEN=`curl -s -X POST http://localhost:8080/api/v1/login -d "@new_user.json" | jq .token | tr -d '"'`
curl -X DELETE -H "Authorization: Bearer $TOKEN" http://localhost:8080/api/v1/user/new_user
```
