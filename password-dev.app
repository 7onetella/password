[password]
heath  = /api/health
memory = 128

env.HTTP_PORT    = ${NOMAD_PORT_http}
env.STAGE        = devpass
env.DB_CONNSTR   = postgres://dev:dev114@tmt-vm11.7onetella.net:5432/devdb?sslmode=disable
env.CRYPTO_TOKEN = test
env.CREDENTIAL   = John.Smith@example.com:users91234