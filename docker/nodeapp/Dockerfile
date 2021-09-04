# FROM 으로 BASE 이미지 로드
FROM node:7

# ADD 명령어로 이미지에 app.js 파일 추가
ADD app.js /app.js

# ENTRYPOINT 명령어로 node를 실행하고 매개변수로 app.js를 전달
# 실행되는 명령어 node app.js
ENTRYPOINT [ "node", "app.js" ]