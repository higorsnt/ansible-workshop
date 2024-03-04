FROM gradle:8-alpine AS build
COPY --chown=gradle:gradle . /home/gradle/src
WORKDIR /home/gradle/src
RUN apk update && apk add gcompat --repository=https://pkgs.alpinelinux.org/package/edge/community/x86_64/grpc-java
RUN gradle build --no-daemon


FROM amazoncorretto:17-alpine
WORKDIR /middleware
COPY --from=build /home/gradle/src/build/libs/*.jar app.jar
EXPOSE 8081
CMD ["java", "-jar", "app.jar"]