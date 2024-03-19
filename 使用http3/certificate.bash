function create(){
HOST=$1
HOST2=$2
DAY=$3

# 创建ca私钥，服务器私钥，ca证书
openssl genrsa -out ca.key 2048
openssl genrsa -out server.key 2048
openssl req -new -x509 -days $DAY -key ca.key  -subj "/CN=${HOST}" -out ca.crt

# 创建服务器证书签名请求
openssl req -new \
    -key server.key \
    -subj "/C=CN/OU=a/O=b/CN=c" \
    -reqexts SAN \
    -config <(cat /etc/ssl/openssl.cnf \
        <(printf "\n[SAN]\nsubjectAltName=DNS:${HOST},DNS:${HOST2}")) \
    -out server.csr

# 通过ca颁发服务器证书
openssl x509 -req -days $DAY \
    -in server.csr -out server.crt \
    -CA ca.crt -CAkey ca.key -CAcreateserial \
    -extensions SAN \
    -extfile <(cat /etc/ssl/openssl.cnf <(printf "[SAN]\nsubjectAltName=DNS:${HOST},DNS:${HOST2}"))
}

# 接收三个参数 地址1，地址2, 有效期天数
# 为多个地址颁发同一个证书，这是SAN的优势
create www.open1.com "*.open1.com" 3660