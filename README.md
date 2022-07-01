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

Set expire time for token

```json
{
  "role": "admin",
  // get epoch number from https://www.epochconverter.com/
  "exp": 1656656742
}
```

For more claims, see [https://datatracker.ietf.org/doc/html/rfc7519#section-4](https://datatracker.ietf.org/doc/html/rfc7519#section-4)

## RSA key pair

Generate private key: `openssl genrsa -out private-key.pem 2048`

Generate public key: `openssl rsa -in private-key.pem -pubout -out public-key.pem`

## Quick start

```
$ openssl genrsa -out private-key.pem 2048
$ openssl rsa -in private-key.pem -pubout -out public-key.pem

$ go build

$ ./genjwt create -c default/claims.json -k private-key.pem
eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.TAb8b67hoLZmKJ5w1DSX47yCpNnSVZItlepkGqWaNDmJIOK1p7ixqZkJFeFUNpo5ZFtP5fdkiZ9BUdlxwSnvuNJVG93mU8-RdQ2hIJGRtD0ugocO2JM3uXQdGZt5q0OHkmL9KN-704gYHkbCK_3kHJbdQ2HROrLxVkpX8rfn0wVeAQ57VKC__6aVVRZUKMxjjWMj_RtEqZ7f2QuNFw6M7n5uzzRhwEJy2DzxFbhCb7cZxkuIVPQx4qOu3SJQH9QnKyacoRWVkRAcVcqS3JCJZ9Qy_ZJqYEYR9qlzdu_amO7p21g2mEXbgeUtO1T_iZn1AnM1P4FCXK18TGzOUVYUfDhtywPvR5MYVEc5pFLGGLmOtKPSNHiNAE6og8LBvHWy-LVVyMZnf6UtodOyJkZeicMS1lMGEbxv3MjdcoEMMdPiVdB7Uytmd9J4CYoEI6ZOW-o_OZMo74ZHk6MlTSj8X_lXV59RuwBzYa3FNrtDRARMuS44_VbGWOD40TvDNFYHGXdtoypQgK_br7cAxHe-Dnzfx76Eyz77OyG7t-3mv_eCvvaYZbXrvijDBJ4Haz7iZ3wBCJ3PUFROsL-oUaCTxqJiQQLlVnMOgrq4AOdxqiud1l6mbbLzEqRqjThNwGzNmmg4E1ybudenrk5gRT-M3IPFoVJYviYmnAjvDuoVhQA

$ TOKEN=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJyb2xlIjoiYWRtaW4ifQ.TAb8b67hoLZmKJ5w1DSX47yCpNnSVZItlepkGqWaNDmJIOK1p7ixqZkJFeFUNpo5ZFtP5fdkiZ9BUdlxwSnvuNJVG93mU8-RdQ2hIJGRtD0ugocO2JM3uXQdGZt5q0OHkmL9KN-704gYHkbCK_3kHJbdQ2HROrLxVkpX8rfn0wVeAQ57VKC__6aVVRZUKMxjjWMj_RtEqZ7f2QuNFw6M7n5uzzRhwEJy2DzxFbhCb7cZxkuIVPQx4qOu3SJQH9QnKyacoRWVkRAcVcqS3JCJZ9Qy_ZJqYEYR9qlzdu_amO7p21g2mEXbgeUtO1T_iZn1AnM1P4FCXK18TGzOUVYUfDhtywPvR5MYVEc5pFLGGLmOtKPSNHiNAE6og8LBvHWy-LVVyMZnf6UtodOyJkZeicMS1lMGEbxv3MjdcoEMMdPiVdB7Uytmd9J4CYoEI6ZOW-o_OZMo74ZHk6MlTSj8X_lXV59RuwBzYa3FNrtDRARMuS44_VbGWOD40TvDNFYHGXdtoypQgK_br7cAxHe-Dnzfx76Eyz77OyG7t-3mv_eCvvaYZbXrvijDBJ4Haz7iZ3wBCJ3PUFROsL-oUaCTxqJiQQLlVnMOgrq4AOdxqiud1l6mbbLzEqRqjThNwGzNmmg4E1ybudenrk5gRT-M3IPFoVJYviYmnAjvDuoVhQA

$ ./genjwt verify -t $TOKEN -k public-key.pem
token is valid
===== token:
header:  map[alg:RS256 typ:JWT]
claims:  map[role:admin]
method:  &{RS256 SHA-256}
signature:  TAb8b67hoLZmKJ5w1DSX47yCpNnSVZItlepkGqWaNDmJIOK1p7ixqZkJFeFUNpo5ZFtP5fdkiZ9BUdlxwSnvuNJVG93mU8-RdQ2hIJGRtD0ugocO2JM3uXQdGZt5q0OHkmL9KN-704gYHkbCK_3kHJbdQ2HROrLxVkpX8rfn0wVeAQ57VKC__6aVVRZUKMxjjWMj_RtEqZ7f2QuNFw6M7n5uzzRhwEJy2DzxFbhCb7cZxkuIVPQx4qOu3SJQH9QnKyacoRWVkRAcVcqS3JCJZ9Qy_ZJqYEYR9qlzdu_amO7p21g2mEXbgeUtO1T_iZn1AnM1P4FCXK18TGzOUVYUfDhtywPvR5MYVEc5pFLGGLmOtKPSNHiNAE6og8LBvHWy-LVVyMZnf6UtodOyJkZeicMS1lMGEbxv3MjdcoEMMdPiVdB7Uytmd9J4CYoEI6ZOW-o_OZMo74ZHk6MlTSj8X_lXV59RuwBzYa3FNrtDRARMuS44_VbGWOD40TvDNFYHGXdtoypQgK_br7cAxHe-Dnzfx76Eyz77OyG7t-3mv_eCvvaYZbXrvijDBJ4Haz7iZ3wBCJ3PUFROsL-oUaCTxqJiQQLlVnMOgrq4AOdxqiud1l6mbbLzEqRqjThNwGzNmmg4E1ybudenrk5gRT-M3IPFoVJYviYmnAjvDuoVhQA
```
