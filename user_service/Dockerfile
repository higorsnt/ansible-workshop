FROM amazoncorretto:17-alpine

WORKDIR /user_service

COPY . .

RUN ./gradlew build --no-daemon

COPY build/libs/*.jar app.jar

EXPOSE 8080

RUN chmod +x entrypoint.sh
ENTRYPOINT ["entrypoint.sh"]