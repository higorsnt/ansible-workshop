package br.edu.ufcg.user_service.handler;

import br.edu.ufcg.user_service.dto.error.ExceptionMessage;
import jakarta.persistence.EntityNotFoundException;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

@ControllerAdvice
public class ControllerExceptionHandler extends ResponseEntityExceptionHandler {

    @ExceptionHandler(EntityNotFoundException.class)
    public ResponseEntity<ExceptionMessage> handleEntityNotFoundException(EntityNotFoundException exception) {
        return new ResponseEntity<>(new ExceptionMessage(exception.getMessage()), HttpStatus.NOT_FOUND);
    }

}
