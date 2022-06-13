package com.manabie.todo.entity;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.google.gson.annotations.SerializedName;
import io.swagger.annotations.ApiModelProperty;
import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.io.Serializable;
import java.util.UUID;

@Entity
@Table(name = "todo")
@Setter
@Getter
@JsonInclude(JsonInclude.Include.NON_NULL)
public class Todo implements Serializable {

    private static final long serialVersionUID = -4516672169828022295L;

    @Id
    @Column(name = "id", updatable = false, nullable = false)
    @SerializedName("id")
    @JsonProperty("id")
    @ApiModelProperty(example = "3d7a2dc5-e8b3-48c2-8e3d-b1ed7882a082")
    private String id = UUID.randomUUID().toString();

    @Column(name = "name", nullable = false)
    @SerializedName("name")
    @ApiModelProperty(name = "name", example = "name", required = true)
    @JsonProperty("name")
    private String name;

    @Column(name = "title", nullable = false)
    @SerializedName("title")
    @ApiModelProperty(name = "title", example = "title", required = true)
    @JsonProperty("title")
    private String title;

    @Column(name = "note", nullable = false)
    @SerializedName("note")
    @ApiModelProperty(name = "notenote", example = "content", required = true)
    @JsonProperty("note")
    private String note;

    @ApiModelProperty(example = "f667cfa2-5b51-4266-afde-8a5963ec7d2a")
    @JsonProperty("create_by_user_id")
    @SerializedName("create_by_user_id")
    private String createByUserId;

    @Column(name = "date_created")
    @ApiModelProperty()
    @JsonProperty("date_created")
    @SerializedName("date_created")
    private Long dateCreated = System.currentTimeMillis();


//    @OneToOne(fetch = FetchType.LAZY,
//            cascade =  CascadeType.ALL,
//            mappedBy = "todo")
//    private UserTodo userTodo;

}
