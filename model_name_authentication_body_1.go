/*
 * EMQX API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 5.1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type NameAuthenticationBody1 struct {
    AuthnJwtJwks
    AuthnJwtPublicKey
    AuthnJwtHmac
    AuthnHttpPost
    AuthnHttpGet
    AuthnRedisSentinel
    AuthnRedisCluster
    AuthnRedisSingle
    AuthnMongoSharded
    AuthnMongoRs
    AuthnMongoSingle
    AuthnPostgresql
    AuthnMysql
    AuthnBuiltinDb
}
