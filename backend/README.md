# Backend Server for Doraku

## Overview

This directory is a source of backend server for Doraku.  
The code is written in Golang v1.11.5

## Run Server in local

```bash
go build -o doraku-server
./doraku-server
```

## Run Server with Docker

```bash
export RUN_PORT=80
docker run -d --name doraku --restart unless-stopped -p $RUN_PORT:8080 smiyoshi/dorakuserver
```
