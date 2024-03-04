FROM gradle:8-alpine AS build
COPY --chown=gradle:gradle . /home/gradle/src
WORKDIR /home/gradle/src
RUN gradle build --no-daemon


FROM amazoncorretto:17-alpine
WORKDIR /user_service
COPY entrypoint.sh .
COPY --from=build /home/gradle/src/build/libs/*.jar app.jar
EXPOSE 8080
RUN chmod +x entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]