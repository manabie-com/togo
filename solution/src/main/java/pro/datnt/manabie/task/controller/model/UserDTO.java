package pro.datnt.manabie.task.controller.model;

import lombok.Getter;
import lombok.NonNull;
import lombok.Setter;

@Getter
@Setter
public class UserDTO {
    @NonNull
    private String username;
    @NonNull
    private String password;
}
