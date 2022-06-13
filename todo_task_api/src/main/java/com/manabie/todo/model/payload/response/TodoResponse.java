package com.manabie.todo.model.payload.response;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.google.gson.annotations.SerializedName;
import com.manabie.todo.entity.UserTodo;
import io.swagger.annotations.ApiModelProperty;
import lombok.Getter;
import lombok.Setter;

import javax.persistence.Column;
import java.io.Serializable;

/**
 *
 * @author long
 * @version 1.0
 * @created_by long
 */
@Getter
@Setter
public class TodoResponse implements Serializable {

    @SerializedName("id")
    @ApiModelProperty(name = "id", example = "id", required = true)
    @JsonProperty("id")
    private String id;

    @SerializedName("name")
    @ApiModelProperty(name = "name", example = "name", required = true)
    @JsonProperty("name")
    private String name;

    @SerializedName("title")
    @ApiModelProperty(name = "title", example = "title", required = true)
    @JsonProperty("title")
    private String title;

    @SerializedName("note")
    @ApiModelProperty(name = "notenote", example = "content", required = true)
    @JsonProperty("note")
    private String note;

    @ApiModelProperty(example = "f667cfa2-5b51-4266-afde-8a5963ec7d2a")
    @JsonProperty("create_user_id")
    @SerializedName("create_user_id")
    private String createUserId;

    @ApiModelProperty()
    @JsonProperty("date_created")
    @SerializedName("date_created")
    private Long dateCreated = System.currentTimeMillis();
}
