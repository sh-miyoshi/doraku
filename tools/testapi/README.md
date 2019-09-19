# Test tool for backend API

## Overview

This directory includes a command line tool for testing backend API.

## System Requirement

golang v1.11 ~

## Usage

```bash
<command> login --name <user_name> --password <password>

<command> user add --name <user_name>
<command> user delete --name <user_name>
<command> user get --name <user_name>

<command> hobby add --user <user_name> --hobby <hobby_id or hobby_name>
<command> hobby delete --user <user_name> --hobby <hobby_id or hobby_name>
```
