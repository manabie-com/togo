package com.manabie.todo.model.payload.request;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.google.gson.annotations.SerializedName;
import io.swagger.annotations.ApiModelProperty;
import lombok.Getter;
import lombok.Setter;

import java.io.Serializable;

/**
 *
 * @author long
 * @version 1.0
 * @created_by long
 */
@Getter
@Setter
public class TodoRequest implements Serializable {

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
    @JsonProperty("create_by_user_id")
    @SerializedName("create_by_user_id")
    private String createByUserId;
}
