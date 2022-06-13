package com.manabie.todo.model.payload.request;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.google.gson.annotations.SerializedName;
import io.swagger.annotations.ApiModelProperty;
import lombok.Getter;
import lombok.Setter;

import java.io.Serializable;
import java.util.UUID;

/**
 * @author j2ee
 */
@Getter
@Setter
public class TodoModelUpdateRequest implements Serializable {

    @SerializedName("id")
    @JsonProperty("id")
    @ApiModelProperty(example = "3d7a2dc5-e8b3-48c2-8e3d-b1ed7882a082")
    private String id = UUID.randomUUID().toString();

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

    @ApiModelProperty(example = "1")
    @JsonProperty("create_user_id")
    @SerializedName("create_user_id")
    private String createUserId;
}
