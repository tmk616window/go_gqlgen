# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
# RUN mkdir /go/src
WORKDIR /go/api_test

# GoApp起動
# CMD ["go","run","main.go"]
ADD . /go/api_test
