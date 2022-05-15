package com.manabie.todo.configuration;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.manabie.todo.common.HttpResponseConverter;
import com.manabie.todo.utils.AppResponse;
import org.springframework.context.MessageSource;
import org.springframework.http.HttpOutputMessage;
import org.springframework.http.MediaType;
import org.springframework.http.converter.HttpMessageNotWritableException;
import org.springframework.http.converter.json.AbstractJackson2HttpMessageConverter;

import java.io.IOException;
import java.lang.reflect.Type;
import java.util.Locale;

@HttpResponseConverter
public class AppResponseConverter extends AbstractJackson2HttpMessageConverter {

  private final MessageSource messageSource;

  protected AppResponseConverter(ObjectMapper objectMapper, MessageSource messageSource) {
    super(objectMapper, MediaType.APPLICATION_JSON, new MediaType("application", "*+json"));
    this.messageSource = messageSource;
  }

  @Override
  public boolean canRead(Class<?> clazz, MediaType mediaType) {
    return false;
  }

  @Override
  public boolean canWrite(Class<?> clazz, MediaType mediaType) {
    return clazz.equals(AppResponse.class) && super.canWrite(clazz, mediaType);
  }

  @Override
  protected void writeInternal(Object object, Type type, HttpOutputMessage outputMessage)
      throws IOException, HttpMessageNotWritableException {
    AppResponse<?> response = bindingMessage((AppResponse<?>) object);
    super.writeInternal(response, type, outputMessage);
  }

  private AppResponse<?> bindingMessage(AppResponse<?> object) {
    object.setMessage(messageSource.getMessage(object.getMessage(), null, Locale.getDefault()));
    return object;
  }


}
