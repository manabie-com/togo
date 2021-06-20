package pro.datnt.manabie.task.properties;

import lombok.Getter;
import lombok.Setter;
import org.springframework.boot.context.properties.ConfigurationProperties;

@Getter @Setter
@ConfigurationProperties(prefix = "com.manabie.authentication")
public class AuthenticationProperties {
    private String key = "UESwa4AaHvIuRAgmW0g6vQzoyCEg0vcczNpjSYLsCypewi9N3GJfFZWO6Iws5CkX";
    private String signUpUrl = "/users/register";
    private String authorizationHeader = "Authorization";
    private Long expirationTime = 1000L * 60 * 30;
}
