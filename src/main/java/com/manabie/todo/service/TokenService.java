package com.manabie.todo.service;

import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;
import com.manabie.todo.configuration.AppProperties;
import com.manabie.todo.domain.User;
import lombok.AllArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.Optional;

import static com.manabie.todo.utils.Constants.*;

@Service
@AllArgsConstructor
public class TokenService {

  private final AppProperties appProperties;

  public String generateToken(User user) {
    var algorithm = Algorithm.HMAC256(appProperties.getJwtSecretKey());
    var current = new Date();
    return JWT.create()
        .withIssuer(appProperties.getJwtIssuer())
        .withClaim(USER_ID_CLAIM, user.getId())
        .withClaim(TASK_QUOTE_CLAIM, user.getTaskQuote())
        .withClaim(USERNAME_CLAIM, user.getUsername())
        .withIssuedAt(current)
        .withExpiresAt(
            Date.from(current.toInstant().plusSeconds(appProperties.getJwtExpireInSeconds())))
        .sign(algorithm);
  }

  public Optional<User> validate(String token) {
    var algorithm = Algorithm.HMAC256(appProperties.getJwtSecretKey());
    var verifier = JWT.require(algorithm)
        .withIssuer(appProperties.getJwtIssuer())
        .build();
    try {
      var user = verifier.verify(token);

      return Optional.of(
          User.builder()
              .id(user.getClaim(USER_ID_CLAIM).asLong())
              .username(user.getClaim(USERNAME_CLAIM).asString())
              .taskQuote(user.getClaim(TASK_QUOTE_CLAIM).asLong())
              .build()
      );
    } catch (Exception exception) {
      return Optional.empty();
    }
  }
}
