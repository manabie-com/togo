package com.manabie.todotaskapplication.common.utils;

import org.apache.commons.lang3.math.NumberUtils;
import org.apache.logging.log4j.util.Strings;
import org.springframework.util.StringUtils;

import javax.servlet.http.HttpServletRequest;
import javax.swing.text.html.Option;
import java.util.Objects;
import java.util.Optional;

/**
 * @author @quoctrung.phan
 * @created 04/05/2022
 * @project todo-task-application
 */
public class UserUtils {
    private UserUtils() {
    }

    public static Optional<String> getUserIdByHttpServletRequest(HttpServletRequest request) {
        if (Objects.isNull(request)) {
            return Optional.empty();
        }

        String authorizationValue = request.getHeader("Authorization");

        return getUserIdByToken(authorizationValue);
    }

    public static Optional<String> getUserIdByToken(String token) {

        if (Strings.isEmpty(token) || !NumberUtils.isParsable(token)) {
            return Optional.empty();
        }
        return Optional.of(token);
    }
}
