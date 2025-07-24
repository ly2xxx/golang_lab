#!/bin/bash
# Test script for the REST API

API_BASE="http://localhost:8080/api"

echo "=== Testing REST API ==="
echo

# Test health endpoint
echo "1. Testing health endpoint:"
curl -s "$API_BASE/health" | python3 -m json.tool
echo
echo

# Test getting all users
echo "2. Getting all users:"
curl -s "$API_BASE/users" | python3 -m json.tool
echo
echo

# Test getting specific user
echo "3. Getting user with ID 1:"
curl -s "$API_BASE/users/1" | python3 -m json.tool
echo
echo

# Test creating a new user
echo "4. Creating a new user:"
curl -s -X POST \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","email":"alice@example.com","age":28}' \
  "$API_BASE/users" | python3 -m json.tool
echo
echo

# Test updating a user
echo "5. Updating user with ID 1:"
curl -s -X PUT \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Name","age":26}' \
  "$API_BASE/users/1" | python3 -m json.tool
echo
echo

# Test validation error
echo "6. Testing validation (invalid email):"
curl -s -X POST \
  -H "Content-Type: application/json" \
  -d '{"name":"","email":"invalid-email","age":200}' \
  "$API_BASE/users" | python3 -m json.tool
echo
echo

# Test 404 error
echo "7. Testing 404 error (user not found):"
curl -s "$API_BASE/users/999" | python3 -m json.tool
echo
echo

echo "API testing completed!"