FROM golang:1.11-alpine AS build
WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/demo
# 실행파일 만듬

FROM scratch
# 다시시작
# 위 내용이 변경이 잘 안생길 때 사용
# 위 네개의 라인을 하나의 블럭으로 인식하여 실행
COPY --from=build /bin/demo /bin/demo
# 컴파일에 생긴 불필요한 파일 없애고 컴파일된 파일만 복사해서 다시 시작
ENTRYPOINT ["/bin/demo"]
# /bin/demo를 실행