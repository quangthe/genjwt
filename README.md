# genjwt

Generate JWT token for application authentication.

Support:

- Custom claims
- RSA private key PEM format

## Custom claims

Custom payload in token. Example `default/claims.json`

```json
{
  "role": "admin"
}
```

## RSA key pair

Generate private key: `openssl genrsa -out private-key.pem 2048`

Generate public key: `openssl rsa -in private-key.pem -pubout -out public-key.pem`
