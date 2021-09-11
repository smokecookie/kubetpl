FROM ubuntu:latest
RUN apt-get update; apt-get -y install fortune
ADD fortuneloop.sh /bin/fortuneloop.sh
RUN chmod 755 /bin/fortuneloop.sh  
# 실행권한 부여
ENTRYPOINT /bin/fortuneloop.sh  
# 명령어 실행