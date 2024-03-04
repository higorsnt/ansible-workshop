FROM amazoncorretto:17-alpine

WORKDIR /middleware

RUN apk update && apk add gcompat --repository=https://pkgs.alpinelinux.org/package/edge/community/x86_64/grpc-java

COPY . .

RUN ./gradlew build --no-daemon

COPY build/libs/*.jar app.jar

EXPOSE 8081

CMD ["java", "-jar", "app.jar"]